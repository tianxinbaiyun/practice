
@[TOC](并发map --- sync map分析)

## 介绍

本文基于1.10源码分析
如之前的文章可以看到，golang中的map是不支持并发操作的，golang推荐用户直接用读写锁对map进行保护，也有第三方类库使用分段锁。
在1.19版本中，golang基于原本的map，新增了一个支持并发操作的map，叫sync map。

下面我们先介绍一下它的用法，然后在介绍原理，最后详细看看代码。

## 用法

基本api有这几个

Store 写入
Load 读取，返回值有两个，第一个是value，第二个是bool变量表示key是否存在
Delete 删除
LoadOrStore 存在就读，不存在就写
Range 遍历，注意遍历的快照
sync map底层使用map[interface{}]* entry来做存储，所以无论key还是value都是支持多种数据类型。
一个简单的例子：
```
package main

import (
    "fmt"
     "sync"
)


type MySyncMap struct {
    sync.Map
}

func (m MySyncMap) Print(k interface{}) {
    value, ok := m.Load(k)
    fmt.Println(value, ok)
} 

func main() {
        var syncMap MySyncMap 

        syncMap.Print("Key1")

        syncMap.Store("Key1", "Value1")
        syncMap.Print("Key1")

        syncMap.Store("Key2", "Value2")

        syncMap.Store("Key3", 2)
        syncMap.Print("Key3")

        syncMap.Store(4, 4)
        syncMap.Print(4)

        syncMap.Delete("Key1")
        syncMap.Print("Key1")
}'
```
输出：
```
<nil> false
Value1 true
2 true
4 true
<nil> false
```
## 设计原理

常用方案比较

并发hashmap的方案有很多，下面简单提一下几种，然后再讨论golang实现时的考虑。

第一种是最简单的，直接在不支持并发的hashmap上，使用一个读写锁的保护，这也是golang sync map还没出来前，大家常用的方法。这种方法的缺点是写会堵塞读。

第二种是数据库常用的方法，分段锁，每一个读写锁保护一段区间，golang的第三方库也有人是这么实现的。java的ConcurrentHashMap也是这么实现的。
平均情况下这样的性能还挺好的，但是极端情况下，如果某个区间有热点写，那么那个区间的读请求也会受到影响。

第三种方法是我们C++自己造轮子时经常用的，使用使用链表法解决冲突，然后链表使用CAS去解决并发下冲突，这样读写都是无锁，我觉得这种挺好的，性能非常高，
不知为啥其他语言不这么实现。

然后在《An overview of sync.Map》中有提到，在cpu核数很多的情况下，因为cache contention，reflect.New、sync.RWMutex、atomic.AddUint32都会很慢，
golang团队为了适应cpu核很多的情况，没有采用上面的几种常见的方案。

golang sync map的目标是实现适合读多写少的场景、并且要求稳定性很好，不能出现像分段锁那样读经常被阻塞的情况。
golang sync map基于map做了一层封装，在大部分情况下，不过写入性能比较差。下面来详细说说实现。

## 实现思路

要读受到的影响尽量小，那么最容易想到的想法，就是读写分离。
golang sync map也是受到这个想法的启发（我自认为）设计出来的。
使用了两个map，一个叫read，一个叫dirty，两个map存储的都是指针，指向value数据本身，所以两个map是共享value数据的，更新value对两个map同时可见。

dirty可以进行增删查，当时都要进行加互斥锁。

read中存在的key，可以无锁的读，借助CAS进行无锁的更新、删除操作，但是不能新增key，相当于dirty的一个cache，
由于value共享，所以能通过read对已存在的value进行更新。

read不能新增key，那么数据怎么来的呢？
sync map中会记录miss cache的次数，当miss次数大于等于dirty元素个数时，就会把dirty变成read，原来的dirty清空。

为了方便dirty直接变成read，那么得保证read中存在的数据dirty必须有，
所以在dirty是空的时候，如果要新增一个key，那么会把read中的元素复制到dirty中，然后写入新key。

然后删除操作也很有意思，使用的是延迟删除，优先看read中没有，read中有，就把read中的对应entry指针中的p置为nil，作为一个标记。
在read中标记为nil的，只有在dirty提升为read时才会被实际删除。

## 源码

结构
```
// The zero Map is empty and ready for use. A Map must not be copied after first use.
type Map struct {
    mu Mutex

    // read contains the portion of the map's contents that are safe for
    // concurrent access (with or without mu held).
    //
    // The read field itself is always safe to load, but must only be stored with
    // mu held.
    //
    // Entries stored in read may be updated concurrently without mu, but updating
    // a previously-expunged entry requires that the entry be copied to the dirty
    // map and unexpunged with mu held.
    read atomic.Value // readOnly

    // dirty contains the portion of the map's contents that require mu to be
    // held. To ensure that the dirty map can be promoted to the read map quickly,
    // it also includes all of the non-expunged entries in the read map.
    //
    // Expunged entries are not stored in the dirty map. An expunged entry in the
    // clean map must be unexpunged and added to the dirty map before a new value
    // can be stored to it.
    //
    // If the dirty map is nil, the next write to the map will initialize it by
    // making a shallow copy of the clean map, omitting stale entries.
    dirty map[interface{}]*entry

    // misses counts the number of loads since the read map was last updated that
    // needed to lock mu to determine whether the key was present.
    //
    // Once enough misses have occurred to cover the cost of copying the dirty
    // map, the dirty map will be promoted to the read map (in the unamended
    // state) and the next store to the map will make a new dirty copy.
    misses int
}

//read的实际结构体
// readOnly is an immutable struct stored atomically in the Map.read field.
type readOnly struct {
    m       map[interface{}]*entry
    amended bool // true if the dirty map contains some key not in m.
}

// expunged is an arbitrary pointer that marks entries which have been deleted
// from the dirty map.
var expunged = unsafe.Pointer(new(interface{}))

// An entry is a slot in the map corresponding to a particular key.
type entry struct {
    // p points to the interface{} value stored for the entry.
    //
    // If p == nil, the entry has been deleted and m.dirty == nil.
    //
    // If p == expunged, the entry has been deleted, m.dirty != nil, and the entry
    // is missing from m.dirty.
    //
    // Otherwise, the entry is valid and recorded in m.read.m[key] and, if m.dirty
    // != nil, in m.dirty[key].
    //
    // An entry can be deleted by atomic replacement with nil: when m.dirty is
    // next created, it will atomically replace nil with expunged and leave
    // m.dirty[key] unset.
    //
    // An entry's associated value can be updated by atomic replacement, provided
    // p != expunged. If p == expunged, an entry's associated value can be updated
    // only after first setting m.dirty[key] = e so that lookups using the dirty
    // map find the entry.
    p unsafe.Pointer // *interface{}
}
```
## sync map结构
mu是用来保护dirty的互斥锁
missed是记录没命中read的次数

注意对于entry.p，有两个特殊值，一个是nil，另一个是expunged。
nil代表的意思是，在read中被删除了，但是dirty中还在，所以能直接更新值(如果dirty==nill的特殊情况,下次写入新值时会复制)；
expunged代表数据在ditry中已经被删除了，更新值的时候要先把这个entry复制到dirty。

### Load 读取

```
// Load returns the value stored in the map for a key, or nil if no
// value is present.
// The ok result indicates whether value was found in the map.
func (m *Map) Load(key interface{}) (value interface{}, ok bool) {
    read, _ := m.read.Load().(readOnly)
    e, ok := read.m[key]
    if !ok && read.amended {
        m.mu.Lock()
        // Avoid reporting a spurious miss if m.dirty got promoted while we were
        // blocked on m.mu. (If further loads of the same key will not miss, it's
        // not worth copying the dirty map for this key.)
        read, _ = m.read.Load().(readOnly)
        e, ok = read.m[key]
        if !ok && read.amended {
            e, ok = m.dirty[key]
            // Regardless of whether the entry was present, record a miss: this key
            // will take the slow path until the dirty map is promoted to the read
            // map.
            m.missLocked()
        }
        m.mu.Unlock()
    }
    if !ok {
        return nil, false
    }
    return e.load()
}

func (e *entry) load() (value interface{}, ok bool) {
    p := atomic.LoadPointer(&e.p)
    if p == nil || p == expunged {
        return nil, false
    }
    return *(*interface{})(p), true
}

func (m *Map) missLocked() {
    m.misses++
    if m.misses < len(m.dirty) {
        return
    }
    m.read.Store(readOnly{m: m.dirty})
    m.dirty = nil
    m.misses = 0
}
```
读取时，先去read读取；如果没有，就加锁，然后去dirty读取，同时调用missLocked()，再解锁。在missLocked中，
会递增misses变量，如果misses>len(dirty)，那么把dirty提升为read，清空原来的dirty。

在代码中，我们可以看到一个double check，检查read没有，上锁，再检查read中有没有，
是因为有可能在第一次检查之后，上锁之前的间隙，dirty提升为read了，这时如果不double check，
可能会导致一个存在的key却返回给调用方说不存在。 在下面的其他操作中，我们经常会看到这个double check。

### Store 写入
```
// Store sets the value for a key.
func (m *Map) Store(key, value interface{}) {
    read, _ := m.read.Load().(readOnly)
    if e, ok := read.m[key]; ok && e.tryStore(&value) {
        return
    }

    m.mu.Lock()
    read, _ = m.read.Load().(readOnly)
    if e, ok := read.m[key]; ok {
        if e.unexpungeLocked() {
            // The entry was previously expunged, which implies that there is a
            // non-nil dirty map and this entry is not in it.
            m.dirty[key] = e
        }
        e.storeLocked(&value)
    } else if e, ok := m.dirty[key]; ok {
        e.storeLocked(&value)
    } else {
        if !read.amended {
            // We're adding the first new key to the dirty map.
            // Make sure it is allocated and mark the read-only map as incomplete.
            m.dirtyLocked()
            m.read.Store(readOnly{m: read.m, amended: true})
        }
        m.dirty[key] = newEntry(value)
    }
    m.mu.Unlock()
}

// tryStore stores a value if the entry has not been expunged.
//
// If the entry is expunged, tryStore returns false and leaves the entry
// unchanged.
func (e *entry) tryStore(i *interface{}) bool {
    p := atomic.LoadPointer(&e.p)
    if p == expunged {
        return false
    }
    for {
        if atomic.CompareAndSwapPointer(&e.p, p, unsafe.Pointer(i)) {
            return true
        }
        p = atomic.LoadPointer(&e.p)
        if p == expunged {
            return false
        }
    }
}

func (m *Map) dirtyLocked() {
    if m.dirty != nil {
        return
    }

    read, _ := m.read.Load().(readOnly)
    m.dirty = make(map[interface{}]*entry, len(read.m))
    for k, e := range read.m {
        if !e.tryExpungeLocked() {
            m.dirty[k] = e
        }
    }
}

func (e *entry) tryExpungeLocked() (isExpunged bool) {
    p := atomic.LoadPointer(&e.p)
    for p == nil {
        if atomic.CompareAndSwapPointer(&e.p, nil, expunged) {
            return true
        }
        p = atomic.LoadPointer(&e.p)
    }
    return p == expunged
}

// unexpungeLocked ensures that the entry is not marked as expunged.
//
// If the entry was previously expunged, it must be added to the dirty map
// before m.mu is unlocked.
func (e *entry) unexpungeLocked() (wasExpunged bool) {
    return atomic.CompareAndSwapPointer(&e.p, expunged, nil)
}
```
写入的时候，先看read中能否查到key，在read中存在的话，直接通过read中的entry来更新值；在read中不存在，那么就上锁，然后double check。
这里需要留意，分几种情况：

double check发现read中存在，如果是expunged，那么就先尝试把expunged替换成nil，
最后如果entry.p==expunged就复制到dirty中，再写入值；否则不用替换直接写入值。
dirty中存在，直接更新
dirty中不存在，如果此时dirty为空，那么需要将read复制到dirty中，最后再把新值写入到dirty中。复制的时候调用的是dirtyLocked()，在复制到dirty的时候，read中为nil的元素，会更新为expunged，并且不复制到dirty中。
我们可以看到，在更新read中的数据时，使用的是tryStore，通过CAS来解决冲突，在CAS出现冲突后，如果发现数据被置为expung，tryStore那么就不会写入数据，而是会返回false，在Store流程中，就是接着往下走，在dirty中写入。

再看下情况1的时候，为啥要那么做。double check的时候，在read中存在，那么就是说在加锁之前，有并发线程先写入了key，然后由Load触发了dirty提升为read，这时dirty可能为空，也可能不为空，但无论dirty状态如何，都是可以直接更新entry.p。如果是expunged的话，那么要先替换成nil，再复制entry到dirty中。

疑问：这里不太懂，为啥在read中直接更新就用cas去更新，跑到下面的流程，就用原子更新，可是尽管上了锁，key在read中存在，那么就会并发写，为啥可以不用cas更新？？

### Delete 删除
```
// Delete deletes the value for a key.
func (m *Map) Delete(key interface{}) {
    read, _ := m.read.Load().(readOnly)
    e, ok := read.m[key]
    if !ok && read.amended {
        m.mu.Lock()
        read, _ = m.read.Load().(readOnly)
        e, ok = read.m[key]
        if !ok && read.amended {
            delete(m.dirty, key)
        }
        m.mu.Unlock()
    }
    if ok {
        e.delete()
    }
}

func (e *entry) delete() (hadValue bool) {
    for {
        p := atomic.LoadPointer(&e.p)
        if p == nil || p == expunged {
            return false
        }
        if atomic.CompareAndSwapPointer(&e.p, p, nil) {
            return true
        }
    }
}
```
删除很简单，read中存在，就把read中的entry.p置为nil，如果只在ditry中存在，那么就直接从dirty中删掉对应的entry。



————————————————

版权声明：本文为简书博主「Gray123archu」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://www.jianshu.com/p/ec51dac3c65b