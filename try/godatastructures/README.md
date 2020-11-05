
go-datastructures
=================

Go-datastructures是一个有用的、高性能的、线程安全的Go数据结构的集合。

### 注意:仅在Go 1.3+时测试。

#### Augmented Tree

Augmented Tree（区间树）的冲突在n维范围。通过一个红黑扩充树实现。
额外的维是在同时插入/查询中处理的，以节省空间，尽管这可能会导致次优的时间复杂度。
使用位数组确定交集。在单一维度中，插入、删除和查询的时间应该是O(log n)。

#### Bitarray

Bitarray（点阵列）用于检测存在性，而不必诉诸于哈希映射的哈希。
要求实体具有uint64唯一标识符。有两种实现，常规的和稀疏的。
稀疏节省了大量的空间，但是插入是O(log n).
在位数组接口上有一些有用的函数来检测两个位数组之间的交集。
这个包还包括长度为32和64的位图，通过将位图存储在无符号整数而不是数组中，为所有操作提供了更快的速度和O(1)。

#### Futures

Futures 一个有用的工具，以发送“广播”消息给听众。
通道存在这样的问题:一旦一个侦听器从通道获取消息，
其他侦听器就不会得到通知。
在很多情况下，我希望将单个事件通知许多侦听器，而这个包可以提供帮助。

#### Queue

Queue（队列）包含一个普通队列和优先队列。这两种实现都不会在发送时阻塞，也不会根据需要不断增长。
如果您试图推入一个已释放的队列，这两种方法也只会返回错误，
并且不会像在一个关闭的通道上发送消息那样出现恐慌。
优先队列还允许您在队列中按优先级顺序放置项目。
如果您给常规队列一个有用的提示，它实际上比通道要快。
优先队列目前有点慢，目标是更新斐波那契堆。

在Queue包中还包括一个MPMC threadsafe环形缓冲区。
这是一个阻塞满/空队列，但将返回一个阻塞线程，如果队列被释放而线程被阻塞。
这可以用来同步goroutines并确保goroutines退出，这样对象就可以被GC。
线程安全是通过只使用CAS操作来实现的，这使得这个队列非常快。
在这个包中可以找到基准测试。

#### Fibonacci Heap

Fibonacci Heap（斐波那契堆）一个标准的斐波那契堆提供通常的操作。可以在理论上最短的时间内执行Dijkstra或Prim算法。

作为通用优先队列也很有用。
与其他堆变体相比，斐波那契堆的特殊之处在于它具有廉价的降伏键操作。
这个堆的查找最小值、插入和合并两个堆的复杂度是恒定的，
降低键的平摊常数复杂度和最小删除或离队列复杂度O(log(n))。
实际上，常数因子很大，所以Fibonacci堆可能比配对堆慢，这取决于使用情况。
基准测试-在项目子文件夹中。堆不是为线程安全而设计的。

#### Range Tree

Range Tree（范围树）用于确定n维点是否落在n维范围内。
但这不是一个典型的范围树，因为我们实际上使用的是n维树
排序的点列表，因为这被证明比尝试传统的范围树更简单和更快，同时节省空间在任何维大于1。
插入是典型的在O(log n^d)处的BBST次数，其中d是维数。

#### Set
Set（集合）我们的Set实现非常简单，接受' interface{} '类型的项，并且只包含几个方法。
如果您的应用程序需要一个更丰富的集合实现而不是类型为“sort.Interface”的列表, see
[xtgo/set](https://github.com/xtgo/set) and
[goware/set](https://github.com/goware/set).

#### Threadsafe
Thread safe（线程安全）以线程安全的方式包含一些常用项的包。
示例:这里有一个线程安全错误，因为我经常发现自己想要同时在许多线程中设置一个错误(是的，我知道，但是通道很慢)。

#### AVL Tree

AVL Tree（平衡二叉树）这是一个不可变AVL BBST分支复制的例子。
对节点的任何操作都会复制该节点的分支。
因此，该树本质上是线程安全的，尽管写操作可能仍然需要序列化。
如果你的用例是大量的读和不频繁的写，这种结构是很好的，因为读是高可用的，但是由于复制，写有点慢。
这个结构作为大量功能性数据结构的基础。

#### X-Fast Trie

X-Fast Trie（X-Fast 字典树）这是一种有趣的设计，将整数视为单词，并使用trie结构通过匹配前缀来减少时间复杂性。
这种结构在查找值或执行前任/继承者类型的查询时速度非常快，但也会导致比线性空间消耗更大的结果。
确切的时间复杂性可以在这个包中找到。

#### Y-Fast Trie

Y-Fast Trie（Y-Fast 字典树）对X-Fast trie的扩展，其中X-Fast trie与其他一些有序数据结构相结合，以减少空间消耗并改进CRUD操作类型。
这些二级结构通常是bst，但我们的实现使用了一个简单的有序列表，因为我相信这将改进缓存的局部性。
我们还使用固定大小的桶来帮助并行化操作。确切的时间复杂性就在这个包里。

#### Fast integer hashmap

Fast integer hashmap（快速整数哈希表）用于检查是否存在但不知道数据边界的数据结构。
如果你有一个有限的小边界，位数组包可能是一个更好的选择。
这个实现使用了一个相当简单的散列算法，结合了线性探测和平面数据结构来提供最多几百万个整数的最佳性能(比本机Golang实现快)。
除此之外，本机实现更快(我相信他们使用的是一个大的-ary b树)。在将来，这将通过b树来实现。

#### Skiplist

Skiplist（跳跃列表）一种有序结构，提供了平摊对数操作，但没有BSTs所要求的复杂的旋转。
然而，在测试中，跳跃表的性能往往比BBST保证的log n时间差得多。
高的节点倾向于“投射阴影”，特别是当需要较大的位码时，因为节点的最佳最大高度通常是基于此。
该包中提供了更详细的性能特征。

#### Sort

Sort（排序）排序包实现了多线程的桶排序，比本地Golang排序包快3倍。
然后使用对称合并来合并这些存储段，类似于Golang包中的稳定排序。然而，我们的算法被修改，使两个排序列表可以合并使用对称分解。

#### Numerics

Numerics（数字）对一些非线性优化问题的早期研究。
最初的实现允许一个带有线性或非线性约束的简单用例。
您可以找到最小/最大或目标的最佳值。
该包目前采用概率全局重新启动系统，试图避免局部临界点。
更多细节可以在该包中找到。

#### B+ Tree

B+ Tree（B+树）初步实现了B+树。Delete方法仍然需要添加，还需要进行一些性能优化。
可以在该包中找到特定的性能特征。尽管BSTs在理论上具有优势，但由于缓存的局部性，b -树通常具有更好的整体性能。

当前的实现是可变的，但是不可变的AVL树可以用来构建不可变版本。
不幸的是，为了使b -树具有通用性，我们需要一个接口，而在CPU分析中开销最大的操作是接口方法，
该方法反过来调用runtime.assertI2T。我们需要泛型。

#### Immutable B Tree
Immutable B Tree（不可变的B树）基于两个原则的btree:不变性和并发性。
对于单值查找和put，它有些慢，但对于批量操作非常快。
可以注入一个persister使这个索引持久。

#### Ctrie

Ctrie(C字典树)并发、无锁哈希数组将trie映射为高效的非阻塞快照。
对于查找，ctry具有与并发跳跃表和并发hashmap相当的性能。
ctry的一个关键优点是它们是动态分配的。
内存消耗总是与Ctrie中的键数成比例，而hashmap通常必须增长和收缩。查找、插入和删除是O(logn)。

与传统并发数据结构相比，ctry的一个有趣优势是支持无锁、可线性化、恒定时间的快照。
大多数并发数据结构不支持快照，而是选择锁或需要静默状态。这允许ctry有O(1)迭代器创建和清除操作和O(logn)大小的检索。

#### Dtrie

Dtrie(D字典树)动态扩展或收缩以提供有效内存分配的持久散列trie。
由于Dtrie是持久的，所以它是不可变的，任何修改都会生成Dtrie的新版本，而不是改变原来的版本。
位图节点允许O(log32(n))的获取、删除和更新操作。插入是O(n)，迭代为O(1)。

#### Persistent List

Persistent List(持续列表)一个持久的、不可变的链表。
所有的写操作都产生一个新的、更新过的结构，它保留和重用以前的版本。
这使用了一种非常实用的、con风格的列表操作。
如您所料，插入、获取、删除和大小操作是O(n)。

#### Simple Graph

Simple Graph(简单图)一种可变的、非持久的无向图，其中不允许并行边和自循环。
添加一条边和检索顶点/边总数的操作是O(1)，而检索与目标相邻的顶点的操作是O(n)。
For more details see [wikipedia](https://en.wikipedia.org/wiki/Graph_(discrete_mathematics)#Simple_graph)

### Installation

 1. Install Go 1.3 or higher.
 2. Run `go get github.com/Workiva/go-datastructures/...`

### Updating

When new code is merged to master, you can use

	go get -u github.com/Workiva/go-datastructures/...

To retrieve the latest version of go-datastructures.

### Testing

To run all the unit tests use these commands:

	cd $GOPATH/src/github.com/Workiva/go-datastructures
	go get -t -u ./...
	go test ./...

Once you've done this once, you can simply use this command to run all unit tests:

	go test ./...


### Contributing

Requirements to commit here:

 - Branch off master, PR back to master.
 - `gofmt`'d code.
 - Compliance with [these guidelines](https://code.google.com/p/go-wiki/wiki/CodeReviewComments)
 - Unit test coverage
 - [Good commit messages](http://tbaggery.com/2008/04/19/a-note-about-git-commit-messages.html)


### Maintainers

 - Dustin Hiatt <[dustin.hiatt@workiva.com](mailto:dustin.hiatt@workiva.com)>
 - Alexander Campbell <[alexander.campbell@workiva.com](mailto:alexander.campbell@workiva.com)>
