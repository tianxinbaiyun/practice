@[TOC](Golang的并发安全)

## 1、通道channel（CAP模型）
channel是Go中代替共享内存的通信方式，
channel从底层实现上是一种队列，
在使用的时候需要通道的发送方和接收方需要知道数据类型和具体通道。
如果有一端没有准备好或消息没有被处理会阻塞当前端。

### Actor模型和CAP模型
Actor模型：在Actor模型中，主角是Actor，
类似一种worker，Actor彼此之间直接发送消息，
不需要经过什么中介，消息是异步发送和处理的

Actor模型描述了一组为了避免并发编程的常见问题的公理:

1.所有Actor状态是Actor本地的，外部无法访问。

2.Actor必须只有通过消息传递进行通信。　
　
3.一个Actor可以响应消息:推出新Actor,改变其内部状态,或将消息发送到一个或多个其他参与者。
　　
4.Actor可能会堵塞自己,但Actor不应该堵塞它运行的线程。

### Cap模型：

Cap模型中，worker之间不直接彼此联系，
而是通过不同channel进行消息发布和侦听。
消息的发送者和接收者之间通过Channel松耦合，
发送者不知道自己消息被哪个接收者消费了，
接收者也不知道是哪个发送者发送的消息。


Go语言的CSP模型是由协程Goroutine与通道Channel实现：

Go协程goroutine: 是一种轻量线程，它不是操作系统的线程，
而是将一个操作系统线程分段使用，通过调度器实现协作式调度。
是一种绿色线程，微线程，它与Coroutine协程也有区别，
能够在发现堵塞后启动新的微线程。

通道channel: 类似Unix的Pipe，用于协程之间通讯和同步。
协程之间虽然解耦，但是它们和Channel有着耦合。

### Cap模型和Actor模型的区别：
Cap的worker彼此之间不直接连接，通信是通过channel实现的

Actor之间是直接通信

channel不但可以传递消息（数据），也可以用作事件通知。

```
package main

func main(){
	done := make(chan struct{})		//发送空结构体（通知）
	c := make(chan string)			//数据传输通道
	go func() {
		s := <-c					//接收消息
		println(s)				
		close(done)					//关闭通道，为结束通知
	}()
	c <- "hi!"						//发送消息
	<-done							//阻塞
}
```

同步模式下需要有goroutine配合，否则会一直阻塞。

异步模式时当缓冲区没有满或者数据未读完的时候，不会出现阻塞：
```text
package main

func main(){
	c := make(chan int, 3)		//创建带有3个缓冲区的异步通道
	c <- 1
	c <- 2						//缓冲区没满
	println(<-c)				//缓冲区有数据不会阻塞
	println(<-c)
}

//在程序中异步通道可以提高程序的性能减少排队阻塞
//channel变量本身为指针
```
channel的收发，channel中还可以使用ok-idom和range模式处理数据
```text
package main

func main() {
	done := make(chan struct{})
	c := make(chan int)
	go func() {
		defer close(done)
		for {
			x, ok := <-c
			if !ok { //判断通道是否关闭
				return
			}
			println(x)
		}
	}()
	c <- 1
	c <- 2
	c <- 3
	close(c)
	<-done

	donr := make(chan struct{})
	cr := make(chan int)
	go func() {
		defer close(donr)
		for x := range cr { //循环获取消息
			println(x)
		}

	}()
	cr <- 1
	cr <- 2
	cr <- 3
	close(cr)
	<-donr
}

//及时使用close函数关闭通道引发结束通知，避免出现可能的死锁
```

### channel的关闭中：close和sync.Cond
一次性事件使用close效率更高，没有多余的开销。
使用sync.Cond实现单播或广播事件

使用close或nil通道时的原则

向已关闭通道发送数据，引发panic

从已关闭接收数据，返回以缓冲数据或零值（在使用channel发送结束时最好使用空struct）

无论收发，nil通道都会阻塞

在使用goroutine和channel时一般使用工厂方法绑定
```text


package main

import "sync"

type receive struct{
	sync.WaitGroup
	date chan int
}
func newR() *receive{
	r := &receive{
		date :make(chan int),
	}
	r.Add(1)
	go func() {
		defer r.Done()
		for x := range r.date{
			println("recv : ",x)
		}
	}()
	return r
}

func main(){
	r := newR()
	r.date <- 1
	r.date <- 2
	close(r.date)		//关闭通道
	r.Wait()
}
//recv :  1
//recv :  2
```

通道可能会引发goroutine leak，指goroutine处于发送或接收阻塞状态，但没有未被唤醒。GC并不收集此类资源，导致他们在队列里长久休眠，形成资源泄露

```text
    package main
    
    import (
    	"runtime"
    	"time"
    )
    
    func testv(){
    	c := make(chan int)
    	for i := 0; i < 10; i++{
    		go func() {
    			<-c
    		}()
    	}
    }
    func main(){
    	testv()
    	for {
    		time.Sleep(time.Second)
    		runtime.GC()				//强制垃圾回收
    	}
    }
//GODEBUG="gctrace=1,schedtrace=1000,scheddetail=1" ./channel5
```
监控程序goroutine状态，查看监控结果可以看出有大量的goroutine处于chan receive状态，不能结束

————————————————

版权声明：本文为CSDN博主「alvin_666」的原创文章，遵循 CC 4.0 BY-SA 版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/alvin_666/article/details/84933164