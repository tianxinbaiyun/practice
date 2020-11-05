@[TOC](golang tcp 编程)

## 打开链接

TCP Socket的连接的建立需要经历客户端和服务端的三次握手的过程。
连接建立过程中，服务端是一个标准的Listen + Accept的结构(可参考上面的代码)，而在客户端Go语言使用net.Dial或DialTimeout进行连接建立：

阻塞Dial：
```text
conn, err := net.Dial("tcp", "google.com:80")
```

超时机制的Dial：
```text
conn, err := net.DialTimeout("tcp", ":8080", 2 * time.Second)

```
## 客户端连接的建立会遇到如下几种情形

1 网络不可达或对方服务未启动

如果传给Dial的Addr是可以立即判断出网络不可达，或者Addr中端口对应的服务没有启动，端口未被监听，Dial会几乎立即返回错误，比如：

2 对方服务的listen backlog满

还有一种场景就是对方服务器很忙，瞬间有大量client端连接尝试向server建立，server端的listen backlog队列满，server accept不及时((即便不accept，那么在backlog数量范畴里面，connect都会是成功的，因为new conn已经加入到server side的listen queue中了，accept只是从queue中取出一个conn而已)，这将导致client端Dial阻塞。我们还是通过例子感受Dial的行为特点：

3 网络延迟较大，Dial阻塞并超时

如果网络延迟较大，TCP握手过程将更加艰难坎坷（各种丢包），
时间消耗的自然也会更长。
Dial这时会阻塞，如果长时间依旧无法建立连接，
则Dial也会返回“ getsockopt: operation timed out”错误。

在连接建立阶段，多数情况下，Dial是可以满足需求的，即便阻塞一小会儿。

但对于某些程序而言，需要有严格的连接时间限定，
如果一定时间内没能成功建立连接，程序可能会需要执行一段“异常”处理逻辑，
为此我们就需要DialTimeout了。
下面的例子将Dial的最长阻塞时间限制在2s内，超出这个时长，
Dial将返回timeout error.

## 服务端读的行为特点
1 Socket中无数据

连接建立后，如果对方未发送数据到socket，
接收方(Server)会阻塞在Read操作上。
执行该Read操作的goroutine也会被挂起。
runtime会监视该socket，直到其有数据才会重新
调度该socket对应的Goroutine完成read。

2 Socket中有部分数据

如果socket中有部分数据，
且长度小于一次Read操作所期望读出的数据长度，
那么Read将会成功读出这部分数据并返回，
而不是等待所有期望数据全部读取后再返回。

3 Socket中有足够数据

如果socket中有数据，
且长度大于等于一次Read操作所期望读出的数据长度，
那么Read将会成功读出这部分数据并返回。

4 Socket关闭

如果client端主动关闭了socket，
那么Server的Read分为“有数据关闭”和“无数据关闭”。

“有数据关闭”是指在client关闭时，
socket中还有server端未读取的数据，Read返回“EOF error“。
“无数据关闭”情形下的结果，那就是Read直接返回EOF error。

5 读取操作超时

有些场合对Read的阻塞时间有严格限制，
在这种情况下，会反复执行了多次，
没能出现“读出部分数据且返回超时错误”的情况。

## 服务端写的行为特点
1 成功写

client端在Write时并未判断Write的返回值。
所谓“成功写”指的就是Write调用返回的n与预期要写入的数据长度相等，
且error = nil。

2 写阻塞

TCP连接通信两端的OS都会为该连接保留数据缓冲，
一端调用Write后，实际上数据是写入到OS的协议栈的数据缓冲的。
TCP是全双工通信，因此每个方向都有独立的数据缓冲。
当发送方将对方的接收缓冲区以及自身的发送缓冲区写满后，
Write就会阻塞。当接收方读取的时候，缓冲区腾出了空间，
客户端就又可以写入了。

3 写入部分数据

Write操作存在写入部分数据的情况，此时服务端关闭，
但是写入的缓冲区不会阻塞。而是后续又写入部分数据后发生了阻塞，
程序需要对这部分写入的部分字节做特定处理。

4 写入超时

如果非要给Write增加一个期限，那我们可以调用SetWriteDeadline方法。
可以看到在写入超时时，依旧存在部分数据写入的情况。

Socket属性

原生Socket API提供了丰富的sockopt设置接口，
但Golang有自己的网络架构模型，golang提供的socket必要的属性设置。

SetKeepAlive 是否开启长连接
SetKeepAlivePeriod 设置长连接的周期，超出会断开
SetLinger 设定当连接中仍有数据等待发送或接受时的Close方法的行为。
SetNoDelay （默认no delay） 设定操作系统是否应该延迟数据包传递，以便发送更少的数据包（Nagle's算法）。默认为真，即数据应该在Write方法后立刻发送。
SetWriteBuffer 连接的系统发送缓冲
SetReadBuffer 连接的系统接收缓冲

关闭连接

由于socket是全双工的，
client和server端在己方已关闭的socket和对方关闭的socket上操作的结果有不同。

从client的结果来看，在己方已经关闭的socket上再进行read和write操作，
会得到”use of closed network connection” error；

从server1的结果来看，
在对方关闭的socket上执行read操作会得到EOF error，
但write操作会成功，因为数据会成功写入己方的内核socket缓冲区中，
即便最终发不到对方socket缓冲区了，
因为己方socket并未关闭。因此当发现对方socket关闭后，
己方应该正确合理处理自己的socket，再继续write已经无任何意义了。

## 参考 Go语言TCP Socket编程

Tcp编程常见问题及解决方法总结

问题1、粘包问题

解决方法一：TCP提供了强制数据立即传送的操作指令push，TCP软件收到该操作指令后，就立即将本段数据发送出去，而不必等待发送缓冲区满；

解决方法二：发送固定长度的消息

解决方法三：把消息的尺寸与消息一块发送

解决方法四：双方约定每次传送的大小

解决方法五：双方约定使用特殊标记来区分消息间隔

解决方法六：标准协议按协议规则处理，如Sip协议

问题2、字符串编码问题

将中文字符串用utf8编码格式转换为字节数组发送时，
一个中文字符可能会占用2～4个字节（假设为3个字节），
这3个字节可能分3次接收，接收端每次接收完后用utf8编码格式转换为字符串，
就会出现乱码，并导致接收长度计算错误的情况。

解决方法一：以字节数做为消息长度的计算单位，而不是字符个数。

解决方法二：发送方和接收方都采用unicode编码格式。

问题3、长连接的保活问题

标准TCP层协议里把对方超时设为2小时，若服务器端超过了2小时还没收到客户的信息，它就发送探测报文段，若发送了10个探测报文段（每一个相隔75S）还没有收到响应，就假定客户出了故障，并终止这个连接。因此应对tcp长连接进行保活。

以下是异步通信时会遇到的问题：

问题4、缓冲区脏数据问题

同步发送的拷贝，是直接拷贝数据到基础系统缓冲区，拷贝完成后返回；

异步发送消息的拷贝，是将Socket自带的Buffer空间内的所有数据，
拷贝到基础系统发送缓冲区，并立即返回；

因此异步发送时缓冲区设置不好会导致接收到脏数据的问题，如下所示：

第一次发送数据：1234567890

第一次接受数据：1234567890

第二次发送数据：abc

第二次接受数据：abc4567890

请参考：http://www.cnblogs.com/tianzhiliang/archive/2010/09/08/1821623.html

解决方法一：将缓冲区的大小设置为实际发送数据的大小。

问题5、内存碎片问题

频繁的申请缓冲区会导致内存碎片的问题。

解决方法一：使用对象池和内存池。

请参考MSDN：http://msdn.microsoft.com/zh-cn/library/bb517542(v=vs.100).aspx

http://msdn.microsoft.com/zh-cn/library/system.net.sockets.socketasynceventargs.socketasynceventargs(v=vs.100).aspx

问题6、乱序问题

多个线程使用异步通信方式向同一个接收端(socket)同时发送数据，会导致接收端接收的数据混乱。如下所示：

线程1第一次发送：123456789，假设未发送完，只发送了123

线程2第一次发送：abcdefgh，假设未发送完，只发送了abc

线程1第二次发送：456789，发送完成

线程2第二次发送：defgh，发送完成

接收端最终接收的数据为：123abc456789defgh。

解决方法一：一个连接的发送端线程排队发送数据。

## 代码示例
服务端：
```text
package main


import (
    "fmt"
    "net"
    "strings"
)

// 读取数据
func handleConnection(conn net.Conn) {

    for {
        buf := make([]byte, 1024)
        if _,err := conn.Read(buf);err == nil {
            result := strings.Replace(string(buf),"\n","",1)
            fmt.Println(result)
        }else{
            fmt.Println(err)
        }

    }
}

func main() {

    /*
    Listen: 返回在一个本地网络地址laddr上监听的Listener。网络类型参数net必须是面向流的网络： "tcp"、"tcp4"、"tcp6"、"unix"或"unixpacket"。
    */
    listener, err := net.Listen("tcp", "localhost:9999")
    if err != nil {
        fmt.Println("listen error:", err)
        return
    }

    //todo 限速算法
    fmt.Println("server listen success")
    for {
        //等待客户端接入
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("accept error:", err)
            break
        }
        // 使用协程
        go handleConnection(conn)
    }
}
```

客户端:
```text
package main

import (
    "bufio"
    "fmt"
    "log"
    "net"
    "os"
    "strings"
    "time"
)

func main() {

    //阻塞Dial
    /*
    Dial:
        在网络network上连接地址address，并返回一个Conn接口。
        可用的网络类型有："tcp"、"tcp4"、"tcp6"、"udp"、"udp4"、"udp6"、"ip"、"ip4"、"ip6"、"unix"、"unixgram"、"unixpacket"
        对TCP和UDP网络，地址格式是host:port或[host]:port
    */
    //conn, err := net.Dial("tcp", "localhost:7777")
    //超时
    conn, err := net.DialTimeout("tcp", "localhost:9999",time.Second*2)
    if err != nil {
        log.Println("dial error:", err)
        return
    }
    fmt.Println("client dial success")

    inputReader := bufio.NewReader(os.Stdin)
    for {

        fmt.Println("Please enter a message? 'quit' exit")
        //读取消息
        input, _ := inputReader.ReadString('\n')
        msg := strings.Trim(input, "\r\n")
        //quit 退出
        if msg == "quit" {
            fmt.Println("quit")
            conn.Write([]byte("client quit "))
            return
        }
        _, err = conn.Write([]byte( msg))
    }
}
```


————————————————

参考 Tcp编程常见问题及解决方法总结

作者：谁不曾年少轻狂过
链接：https://www.jianshu.com/p/9e713f63879d
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。