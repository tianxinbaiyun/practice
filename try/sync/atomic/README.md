@[TOC](sync 原子变量)

## 介绍

go语言提供的原子操作都是非侵入式的，它们由标准库代码包sync/atomic中的众多函数代表。

　　 我们调用sync/atomic中的几个函数可以对几种简单的类型进行原子操作。这些类型包括int32,int64,uint32,uint64,uintptr,unsafe.Pointer,共6个。这些函数的原子操作共有5种：增或减，比较并交换、载入、存储和交换它们提供了不同的功能，切使用的场景也有区别。

## 操作


### 增或减
　　 顾名思义，原子增或减即可实现对被操作值的增大或减少。因此该操作只能操作数值类型。

　　 被用于进行增或减的原子操作都是以“Add”为前缀，并后面跟针对具体类型的名称。
```
//方法源码
func AddUint32(addr *uint32, delta uint32) (new uint32)
```

增
例子：（在原来的基础上加n）
```
atomic.AddUint32(&addr,n)
```

减
例子：(在原来的基础上加n（n为负数))
```
atomic.AddUint32(*addr,uint32(int32(n)))

//或
atomic.AddUint32(&addr,^uint32(-n-1))
```
### 比较并交换
　　 比较并交换----Compare And Swap 简称CAS

　　 他是假设被操作的值未曾被改变（即与旧值相等），并一旦确定这个假设的真实性就立即进行值替换

　　 如果想安全的并发一些类型的值，我们总是应该优先使用CAS
```
//方法源码
func CompareAndSwapInt32(addr *int32, old, new int32) (swapped bool)
例子：（如果addr和old相同,就用new代替addr）

ok:=atomic.CompareAndSwapInt32(&addr,old,new)
```

### 载入
　　 如果一个写操作未完成，有一个读操作就已经发生了，这样读操作使很糟糕的。

　　 为了原子的读取某个值sync/atomic代码包同样为我们提供了一系列的函数。这些函数都以"Load"为前缀，意为载入。
```
//方法源码
func LoadInt32(addr *int32) (val int32)
```
例子
```
fun addValue(delta int32){
    for{
        v:=atomic.LoadInt32(&addr)
        if atomic.CompareAndSwapInt32(&v,addr,(delta+v)){
            break;
        }
    }
}
```

### 存储
　　 与读操作对应的是写入操作，sync/atomic也提供了与原子的值载入函数相对应的原子的值存储函数。这些函数的名称均以“Store”为前缀

　　 在原子的存储某个值的过程中，任何cpu都不会进行针对进行同一个值的读或写操作。如果我们把所有针对此值的写操作都改为原子操作，那么就不会出现针对此值的读操作读操作因被并发的进行而读到修改了一半的情况。

　　 原子操作总会成功，因为他不必关心被操作值的旧值是什么。
```
//方法源码
func StoreInt32(addr *int32, val int32)
```

例子
```
atomic.StoreInt32(被操作值的指针,新值)
atomic.StoreInt32(&value,newaddr)
```

### 交换
　　 原子交换操作，这类函数的名称都以“Swap”为前缀。

　　 与CAS不同，交换操作直接赋予新值，不管旧值。

　　 会返回旧值
```
//方法源码
func SwapInt32(addr *int32, new int32) (old int32)
```

例子
```
atomic.SwapInt32(被操作值的指针,新值)（返回旧值）
oldval：=atomic.StoreInt32(&value,newaddr)
```



---------------------

作者：吃猫的鱼0
链接：https://www.jianshu.com/p/228c119a7d0e
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。