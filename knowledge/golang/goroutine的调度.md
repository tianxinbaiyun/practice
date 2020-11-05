@[TOC](goroutine的调度)

## 一、goroutine简介

goroutine是go语言中最为NB的设计，也是其魅力所在，goroutine的本质是协程，是实现并行计算的核心。
goroutine使用方式非常的简单，只需使用go关键字即可启动一个协程，并且它是处于异步方式运行，你不需要等它运行完成以后在执行以后的代码。

## 二、goroutine内部原理

###　调度模型简介
groutine能拥有强大的并发实现是通过GPM调度模型实现，下面就来解释下goroutine的调度模型。

![Alt](https://images2018.cnblogs.com/blog/1075473/201807/1075473-20180704144900055-654632620.jpg)

Go的调度器内部有四个重要的结构：M，P，S，Sched，如上图所示（Sched未给出）

M:M代表内核级线程，一个M就是一个线程，goroutine就是跑在M之上的；M是一个很大的结构，里面维护小对象内存cache（mcache）、当前执行的goroutine、随机数发生器等等非常多的信息

G:代表一个goroutine，它有自己的栈，instruction pointer和其他信息（正在等待的channel等等），用于调度。

P:P全称是Processor，处理器，它的主要用途就是用来执行goroutine的，所以它也维护了一个goroutine队列，里面存储了所有需要它来执行的goroutine

Sched：代表调度器，它维护有存储M和G的队列以及调度器的一些状态信息等。

### 调度实现

![Alt](https://images2018.cnblogs.com/blog/1075473/201807/1075473-20180704160300058-287296807.jpg)


从上图中看，有2个物理线程M，每一个M都拥有一个处理器P，每一个也都有一个正在运行的goroutine。

P的数量可以通过GOMAXPROCS()来设置，它其实也就代表了真正的并发度，即有多少个goroutine可以同时运行。

图中灰色的那些goroutine并没有运行，而是出于ready的就绪态，正在等待被调度。
P维护着这个队列（称之为runqueue），
Go语言里，启动一个goroutine很容易：go function 就行，所以每有一个go语句被执行，runqueue队列就在其末尾加入一个
goroutine，在下一个调度点，就从runqueue中取出（如何决定取哪个goroutine？）一个goroutine执行。