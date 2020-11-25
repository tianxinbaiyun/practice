@[TOC](Socket技术详解)

## １、什么是Socket
在计算机通信领域，socket 被翻译为“套接字”，它是计算机之间进行通信的一种约定或一种方式。

通过 socket 这种约定，一台计算机可以接收其他计算机的数据，也可以向其他计算机发送数据

　　socket起源于Unix，而Unix/Linux基本哲学之一就是“一切皆文件”，都可以用“打开open –> 读写write/read –> 关闭close”模式来操作。
　　我的理解就是Socket就是该模式的一个实现：即socket是一种特殊的文件，一些socket函数就是对其进行的操作（读/写IO、打开、关闭）。
　　Socket()函数返回一个整型的Socket描述符，随后的连接建立、数据传输等操作都是通过该Socket实现的。


## 2.网络中进程如何通信

既然Socket主要是用来解决网络通信的，那么我们就来理解网络中进程是如何通信的。

### 2.1、本地进程间通信

a、消息传递（管道、消息队列、FIFO）

b、同步（互斥量、条件变量、读写锁、文件和写记录锁、信号量）？【不是很明白】

c、共享内存（匿名的和具名的，eg:channel）

d、远程过程调用(RPC)

### 2.2、网络中进程如何通信

我们要理解网络中进程如何通信，得解决两个问题：

ａ、我们要如何标识一台主机，即怎样确定我们将要通信的进程是在那一台主机上运行。

ｂ、我们要如何标识唯一进程，本地通过pid标识，网络中应该怎样标识？

解决办法：

ａ、TCP/IP协议族已经帮我们解决了这个问题，网络层的“ip地址”可以唯一标识网络中的主机

ｂ、传输层的“协议+端口”可以唯一标识主机中的应用程序（进程），因此，我们利用三元组（ip地址，协议，端口）就可以标识网络的进程了，网络中的进程通信就可以利用这个标志与其它进程进行交互


## ３、Socket怎么通信

现在，我们知道了网络中进程间如何通信，即利用三元组【ip地址，协议，端口】可以进行网络间通信了，那我们应该怎么实现了，
因此，我们socket应运而生，它就是利用三元组解决网络通信的一个中间件工具，
就目前而言，几乎所有的应用程序都是采用socket，如UNIX BSD的套接字（socket）和UNIX System V的TLI（已经被淘汰）。

Socket通信的数据传输方式，常用的有两种：

ａ、SOCK_STREAM：

表示面向连接的数据传输方式。数据可以准确无误地到达另一台计算机，如果损坏或丢失，可以重新发送，但效率相对较慢。
常见的 http 协议就使用 SOCK_STREAM 传输数据，因为要确保数据的正确性，否则网页不能正常解析。

ｂ、SOCK_DGRAM：

表示无连接的数据传输方式。计算机只管传输数据，不作数据校验，如果数据在传输中损坏，或者没有到达另一台计算机，是没有办法补救的。也就是说，数据错了就错了，无法重传。
因为 SOCK_DGRAM 所做的校验工作少，所以效率比 SOCK_STREAM 高。

例如：QQ 视频聊天和语音聊天就使用 SOCK_DGRAM 传输数据，因为首先要保证通信的效率，尽量减小延迟，而数据的正确性是次要的，即使丢失很小的一部分数据，视频和音频也可以正常解析，最多出现噪点或杂音，不会对通信质量有实质的影响

## ４、TCP/IP协议

### 4.1、概念
TCP/IP【TCP（传输控制协议）和IP（网际协议）】提供点对点的链接机制，将数据应该如何封装、定址、传输、路由以及在目的地如何接收，都加以标准化。它将软件通信过程抽象化为四个抽象层，采取协议堆栈的方式，分别实现出不同通信协议。协议族下的各种协议，依其功能不同，被分别归属到这四个层次结构之中，常被视为是简化的七层OSI模型。

它们之间好比送信的线路和驿站的作用，比如要建议送信驿站，必须得了解送信的各个细节。

TCP（Transmission Control Protocol，传输控制协议）是一种面向连接的、可靠的、基于字节流的通信协议，数据在传输前要建立连接，传输完毕后还要断开连接，客户端在收发数据前要使用 connect() 函数和服务器建立连接。建立连接的目的是保证IP地址、端口、物理链路等正确无误，为数据的传输开辟通道。
TCP建立连接时要传输三个数据包，俗称三次握手（Three-way Handshaking）。可以形象的比喻为下面的对话：
```
[Shake 1] 套接字A：“你好，套接字B，我这里有数据要传送给你，建立连接吧。”
[Shake 2] 套接字B：“好的，我这边已准备就绪。”
[Shake 3] 套接字A：“谢谢你受理我的请求。
```
### 4.2、TCP的粘包问题以及数据的无边界性：　

https://blog.csdn.net/m0_37947204/article/details/80490512

### 4.3、TCP数据报结构：

20180529001000428.jpeg: ![Alt](https://imgconvert.csdnimg.cn/aHR0cHM6Ly91cGxvYWQtaW1hZ2VzLmppYW5zaHUuaW8vdXBsb2FkX2ltYWdlcy8xMTM2MjU4NC1iZjFiZmZjY2Q5Y2NlYWZmLmpwZWc?x-oss-process=image/format,png)

带阴影的几个字段需要重点说明一下：

(1) 序号：Seq（Sequence Number）序号占32位，用来标识从计算机A发送到计算机B的数据包的序号，计算机发送数据时对此进行标记。

(2) 确认号：Ack（Acknowledge Number）确认号占32位，客户端和服务器端都可以发送，Ack = Seq + 1。

(3) 标志位：每个标志位占用1Bit，共有6个，分别为 URG、ACK、PSH、RST、SYN、FIN，具体含义如下：

（1）URG：紧急指针（urgent pointer）有效。

（2）ACK：确认序号有效。

（3）PSH：接收方应该尽快将这个报文交给应用层。

（4）RST：重置连接。

（5）SYN：建立一个新连接。

（6）FIN：断开一个连接。

### 4.4、连接的建立（三次握手）：
使用 connect() 建立连接时，客户端和服务器端会相互发送三个数据包，请看下图：
　　
![Alt](https://imgconvert.csdnimg.cn/aHR0cHM6Ly91cGxvYWQtaW1hZ2VzLmppYW5zaHUuaW8vdXBsb2FkX2ltYWdlcy8xMTM2MjU4NC03NWMyMDhlZGNmYjk4NmZjLmpwZWc?x-oss-process=image/format,png)

客户端调用 socket() 函数创建套接字后，因为没有建立连接，所以套接字处于CLOSED状态；服务器端调用 listen() 函数后，套接字进入LISTEN状态，开始监听客户端请求
这时客户端发起请求：

1) 当客户端调用 connect() 函数后，TCP协议会组建一个数据包，并设置 SYN 标志位，表示该数据包是用来建立同步连接的。同时生成一个随机数字 1000，填充“序号（Seq）”字段，表示该数据包的序号。完成这些工作，开始向服务器端发送数据包，客户端就进入了SYN-SEND状态。
　　
2) 服务器端收到数据包，检测到已经设置了 SYN 标志位，就知道这是客户端发来的建立连接的“请求包”。服务器端也会组建一个数据包，并设置 SYN 和 ACK 标志位，SYN 表示该数据包用来建立连接，ACK 用来确认收到了刚才客户端发送的数据包
　　服务器生成一个随机数 2000，填充“序号（Seq）”字段。2000 和客户端数据包没有关系。
　　服务器将客户端数据包序号（1000）加1，得到1001，并用这个数字填充“确认号（Ack）”字段。
　　服务器将数据包发出，进入SYN-RECV状态

3) 客户端收到数据包，检测到已经设置了 SYN 和 ACK 标志位，就知道这是服务器发来的“确认包”。客户端会检测“确认号（Ack）”字段，看它的值是否为 1000+1，如果是就说明连接建立成功。
　　接下来，客户端会继续组建数据包，并设置 ACK 标志位，表示客户端正确接收了服务器发来的“确认包”。同时，将刚才服务器发来的数据包序号（2000）加1，得到 2001，并用这个数字来填充“确认号（Ack）”字段。
　　客户端将数据包发出，进入ESTABLISED状态，表示连接已经成功建立。

4) 服务器端收到数据包，检测到已经设置了 ACK 标志位，就知道这是客户端发来的“确认包”。服务器会检测“确认号（Ack）”字段，看它的值是否为 2000+1，如果是就说明连接建立成功，服务器进入ESTABLISED状态。
　　至此，客户端和服务器都进入了ESTABLISED状态，连接建立成功，接下来就可以收发数据了。

###  4.5、TCP四次握手断开连接
建立连接非常重要，它是数据正确传输的前提；断开连接同样重要，它让计算机释放不再使用的资源。如果连接不能正常断开，不仅会造成数据传输错误，还会导致套接字不能关闭，持续占用资源，如果并发量高，服务器压力堪忧。
断开连接需要四次握手，可以形象的比喻为下面的对话：
```shell
[Shake 1] 套接字A：“任务处理完毕，我希望断开连接。”
[Shake 2] 套接字B：“哦，是吗？请稍等，我准备一下。”
等待片刻后……
[Shake 3] 套接字B：“我准备好了，可以断开连接了。”
[Shake 4] 套接字A：“好的，谢谢合作。”
```
下图演示了客户端主动断开连接的场景：


![Alt](https://imgconvert.csdnimg.cn/aHR0cHM6Ly91cGxvYWQtaW1hZ2VzLmppYW5zaHUuaW8vdXBsb2FkX2ltYWdlcy8xMTM2MjU4NC03NWMyMDhlZGNmYjk4NmZjLmpwZWc?x-oss-process=image/format,png)

建立连接后，客户端和服务器都处于ESTABLISED状态。这时，客户端发起断开连接的请求：

客户端调用 close() 函数后，向服务器发送 FIN 数据包，进入FIN_WAIT_1状态。FIN 是 Finish 的缩写，表示完成任务需要断开连接。

服务器收到数据包后，检测到设置了 FIN 标志位，知道要断开连接，于是向客户端发送“确认包”，进入CLOSE_WAIT状态。

注意：服务器收到请求后并不是立即断开连接，而是先向客户端发送“确认包”，告诉它我知道了，我需要准备一下才能断开连接。

客户端收到“确认包”后进入FIN_WAIT_2状态，等待服务器准备完毕后再次发送数据包。

等待片刻后，服务器准备完毕，可以断开连接，于是再主动向客户端发送 FIN 包，告诉它我准备好了，断开连接吧。然后进入LAST_ACK状态。

客户端收到服务器的 FIN 包后，再向服务器发送 ACK 包，告诉它你断开连接吧。然后进入TIME_WAIT状态。

服务器收到客户端的 ACK 包后，就断开连接，关闭套接字，进入CLOSED状态。

### 4.6、关于 TIME_WAIT 状态的说明
客户端最后一次发送 ACK包后进入 TIME_WAIT 状态，而不是直接进入 CLOSED 状态关闭连接，这是为什么呢？

TCP 是面向连接的传输方式，必须保证数据能够正确到达目标机器，不能丢失或出错，而网络是不稳定的，随时可能会毁坏数据，所以机器A每次向机器B发送数据包后，都要求机器B”确认“，回传ACK包，告诉机器A我收到了，这样机器A才能知道数据传送成功了。如果机器B没有回传ACK包，机器A会重新发送，直到机器B回传ACK包。

客户端最后一次向服务器回传ACK包时，有可能会因为网络问题导致服务器收不到，服务器会再次发送 FIN 包，如果这时客户端完全关闭了连接，那么服务器无论如何也收不到ACK包了，所以客户端需要等待片刻、确认对方收到ACK包后才能进入CLOSED状态。那么，要等待多久呢？

数据包在网络中是有生存时间的，超过这个时间还未到达目标主机就会被丢弃，并通知源主机。这称为报文最大生存时间（MSL，Maximum Segment Lifetime）。TIME_WAIT 要等待 2MSL 才会进入 CLOSED 状态。ACK 包到达服务器需要 MSL 时间，服务器重传 FIN 包也需要 MSL 时间，2MSL 是数据包往返的最大时间，如果 2MSL 后还未收到服务器重传的 FIN 包，就说明服务器已经收到了 ACK 包

### 4.7.优雅的断开连接–shutdown()

close()/closesocket()和shutdown()的区别

确切地说，close() / closesocket() 用来关闭套接字，将套接字描述符（或句柄）从内存清除，之后再也不能使用该套接字，与C语言中的 fclose() 类似。应用程序关闭套接字后，与该套接字相关的连接和缓存也失去了意义，TCP协议会自动触发关闭连接的操作。

shutdown() 用来关闭连接，而不是套接字，不管调用多少次 shutdown()，套接字依然存在，直到调用 close() / closesocket() 将套接字从内存清除。
调用 close()/closesocket() 关闭套接字时，或调用 shutdown() 关闭输出流时，都会向对方发送 FIN 包。FIN 包表示数据传输完毕，计算机收到 FIN 包就知道不会再有数据传送过来了。

默认情况下，close()/closesocket() 会立即向网络中发送FIN包，不管输出缓冲区中是否还有数据，而shutdown() 会等输出缓冲区中的数据传输完毕再发送FIN包。也就意味着，调用 close()/closesocket() 将丢失输出缓冲区中的数据，而调用 shutdown() 不会




## ５、OSI模型
TCP/IP对OSI的网络模型层进行了划分如下：


![Alt](https://imgconvert.csdnimg.cn/aHR0cHM6Ly91cGxvYWQtaW1hZ2VzLmppYW5zaHUuaW8vdXBsb2FkX2ltYWdlcy8xMTM2MjU4NC1kNjI3NWFjMjVhYmFjNWNjLmpwZWc?x-oss-process=image/format,png)
TCP/IP协议参考模型把所有的TCP/IP系列协议归类到四个抽象层中
　　应用层：TFTP，HTTP，SNMP，FTP，SMTP，DNS，Telnet 等等
　　传输层：TCP，UDP
　　网络层：IP，ICMP，OSPF，EIGRP，IGMP
　　数据链路层：SLIP，CSLIP，PPP，MTU
　　每一抽象层建立在低一层提供的服务上，并且为高一层提供服务，看起来大概是这样子的
　　
![Alt](https://imgconvert.csdnimg.cn/aHR0cHM6Ly91cGxvYWQtaW1hZ2VzLmppYW5zaHUuaW8vdXBsb2FkX2ltYWdlcy8xMTM2MjU4NC0yZDI2MDEzYzc1ZWU0NWUxLnBuZw?x-oss-process=image/format,png)

![Alt](https://imgconvert.csdnimg.cn/aHR0cHM6Ly91cGxvYWQtaW1hZ2VzLmppYW5zaHUuaW8vdXBsb2FkX2ltYWdlcy8xMTM2MjU4NC0xNTdhMzk4ODJmMjYzMzYxLnBuZw?x-oss-process=image/format,png)

## ６、Socket常用函数接口及其原理
图解socket函数：



![Alt](https://imgconvert.csdnimg.cn/aHR0cHM6Ly91cGxvYWQtaW1hZ2VzLmppYW5zaHUuaW8vdXBsb2FkX2ltYWdlcy8xMTM2MjU4NC0zMWQ1OTQzNjNmOGNhZmEwLnBuZw?x-oss-process=image/format,png)

![alt](https://imgconvert.csdnimg.cn/aHR0cHM6Ly91cGxvYWQtaW1hZ2VzLmppYW5zaHUuaW8vdXBsb2FkX2ltYWdlcy8xMTM2MjU4NC0zMWQ1OTQzNjNmOGNhZmEwLnBuZw?x-oss-process=image/format,png)
### 6.1、使用socket()函数创建套接字
```
int socket(int af, int type, int protocol);
```
1.af 为地址族（Address Family），也就是 IP 地址类型，常用的有 AF_INET 和 AF_INET6。AF 是“Address Family”的简写，INET是“Inetnet”的简写。AF_INET 表示 IPv4 地址，例如 127.0.0.1；AF_INET6 表示 IPv6 地址，例如 1030::C9B4:FF12:48AA:1A2B。
大家需要记住127.0.0.1，它是一个特殊IP地址，表示本机地址，后面的教程会经常用到。

2.type 为数据传输方式，常用的有 SOCK_STREAM 和 SOCK_DGRAM

3.protocol 表示传输协议，常用的有 IPPROTO_TCP 和 IPPTOTO_UDP，分别表示 TCP 传输协议和 UDP 传输协议

### 6.2、使用bind()和connect()函数

socket() 函数用来创建套接字，确定套接字的各种属性，然后服务器端要用 bind() 函数将套接字与特定的IP地址和端口绑定起来，只有这样，流经该IP地址和端口的数据才能交给套接字处理；而客户端要用 connect() 函数建立连接
```
int bind(int sock, struct sockaddr *addr, socklen_t addrlen);  
```
sock 为 socket 文件描述符，addr 为 sockaddr 结构体变量的指针，addrlen 为 addr 变量的大小，可由 sizeof() 计算得出
下面的代码，将创建的套接字与IP地址 127.0.0.1、端口 1234 绑定：
```
//创建套接字
int serv_sock = socket(AF_INET, SOCK_STREAM, IPPROTO_TCP);
//创建sockaddr_in结构体变量
struct sockaddr_in serv_addr;
memset(&serv_addr, 0, sizeof(serv_addr));  //每个字节都用0填充
serv_addr.sin_family = AF_INET;  //使用IPv4地址
serv_addr.sin_addr.s_addr = inet_addr("127.0.0.1");  //具体的IP地址
serv_addr.sin_port = htons(1234);  //端口
//将套接字和IP、端口绑定
bind(serv_sock, (struct sockaddr*)&serv_addr, sizeof(serv_addr));
connect() 函数用来建立连接，它的原型为：

int connect(int sock, struct sockaddr *serv_addr, socklen_t addrlen); 
```
### 6.3、使用listen()和accept()函数
于服务器端程序，使用 bind() 绑定套接字后，还需要使用 listen() 函数让套接字进入被动监听状态，再调用 accept() 函数，就可以随时响应客户端的请求了。
通过** listen() 函数**可以让套接字进入被动监听状态，它的原型为：
```
int listen(int sock, int backlog); 
```
sock 为需要进入监听状态的套接字，backlog 为请求队列的最大长度。
所谓被动监听，是指当没有客户端请求时，套接字处于“睡眠”状态，只有当接收到客户端请求时，套接字才会被“唤醒”来响应请求。

请求队列
当套接字正在处理客户端请求时，如果有新的请求进来，套接字是没法处理的，只能把它放进缓冲区，待当前请求处理完毕后，再从缓冲区中读取出来处理。如果不断有新的请求进来，它们就按照先后顺序在缓冲区中排队，直到缓冲区满。这个缓冲区，就称为请求队列（Request Queue）。

缓冲区的长度（能存放多少个客户端请求）可以通过 listen() 函数的 backlog 参数指定，但究竟为多少并没有什么标准，可以根据你的需求来定，并发量小的话可以是10或者20。

如果将 backlog 的值设置为 SOMAXCONN，就由系统来决定请求队列长度，这个值一般比较大，可能是几百，或者更多。

当请求队列满时，就不再接收新的请求，对于 Linux，客户端会收到 ECONNREFUSED 错误

注意：listen() 只是让套接字处于监听状态，并没有接收请求。接收请求需要使用 accept() 函数。

当套接字处于监听状态时，可以通过 accept() 函数来接收客户端请求。它的原型为：
```
int accept(int sock, struct sockaddr *addr, socklen_t *addrlen); 
```
它的参数与 listen() 和 connect() 是相同的：sock 为服务器端套接字，addr 为 sockaddr_in 结构体变量，addrlen 为参数 addr 的长度，可由 sizeof() 求得。

accept() 返回一个新的套接字来和客户端通信，addr 保存了客户端的IP地址和端口号，而 sock 是服务器端的套接字，大家注意区分。后面和客户端通信时，要使用这个新生成的套接字，而不是原来服务器端的套接字。

最后需要说明的是：listen() 只是让套接字进入监听状态，并没有真正接收客户端请求，listen() 后面的代码会继续执行，直到遇到 accept()。accept() 会阻塞程序执行（后面代码不能被执行），直到有新的请求到来。

### 6.4、socket数据的接收和发送
Linux下数据的接收和发送
Linux 不区分套接字文件和普通文件，使用 write() 可以向套接字中写入数据，使用 read() 可以从套接字中读取数据。

前面我们说过，两台计算机之间的通信相当于两个套接字之间的通信，在服务器端用 write() 向套接字写入数据，客户端就能收到，然后再使用 read() 从套接字中读取出来，就完成了一次通信。
write() 的原型为：
```
ssize_t write(int fd, const void *buf, size_t nbytes);
```
fd 为要写入的文件的描述符，buf 为要写入的数据的缓冲区地址，nbytes 为要写入的数据的字节数。
write() 函数会将缓冲区 buf 中的 nbytes 个字节写入文件 fd，成功则返回写入的字节数，失败则返回 -1。
read() 的原型为：
```
ssize_t read(int fd, void *buf, size_t nbytes);
```
fd 为要读取的文件的描述符，buf 为要接收数据的缓冲区地址，nbytes 为要读取的数据的字节数。

read() 函数会从 fd 文件中读取 nbytes 个字节并保存到缓冲区 buf，成功则返回读取到的字节数（但遇到文件结尾则返回0），失败则返回 -1。

### 6.5、socket缓冲区以及阻塞模式
socket缓冲区
每个 socket 被创建后，都会分配两个缓冲区，输入缓冲区和输出缓冲区。

write()/send() 并不立即向网络中传输数据，而是先将数据写入缓冲区中，再由TCP协议将数据从缓冲区发送到目标机器。一旦将数据写入到缓冲区，函数就可以成功返回，不管它们有没有到达目标机器，也不管它们何时被发送到网络，这些都是TCP协议负责的事情。

TCP协议独立于 write()/send() 函数，数据有可能刚被写入缓冲区就发送到网络，也可能在缓冲区中不断积压，多次写入的数据被一次性发送到网络，这取决于当时的网络情况、当前线程是否空闲等诸多因素，不由程序员控制。

read()/recv() 函数也是如此，也从输入缓冲区中读取数据，而不是直接从网络中读取

20180528234331238.jpeg
这些I/O缓冲区特性可整理如下：


（1）I/O缓冲区在每个TCP套接字中单独存在；
（2）I/O缓冲区在创建套接字时自动生成；
（3）即使关闭套接字也会继续传送输出缓冲区中遗留的数据；
（4）关闭套接字将丢失输入缓冲区中的数据。
输入输出缓冲区的默认大小一般都是 8K，可以通过 getsockopt() 函数获取：
```

unsigned optVal;
int optLen = sizeof(int);
getsockopt(servSock, SOL_SOCKET, SO_SNDBUF, (char*)&optVal, &optLen);
printf("Buffer length: %d\n", optVal);
```
### 6.5、 阻塞模式
对于TCP套接字（默认情况下），当使用 write()/send() 发送数据时：

1) 首先会检查缓冲区，如果缓冲区的可用空间长度小于要发送的数据，那么 write()/send() 会被阻塞（暂停执行），直到缓冲区中的数据被发送到目标机器，腾出足够的空间，才唤醒 write()/send() 函数继续写入数据。

2) 如果TCP协议正在向网络发送数据，那么输出缓冲区会被锁定，不允许写入，write()/send() 也会被阻塞，直到数据发送完毕缓冲区解锁，write()/send() 才会被唤醒。

3) 如果要写入的数据大于缓冲区的最大长度，那么将分批写入。

4) 直到所有数据被写入缓冲区 write()/send() 才能返回。

当使用 read()/recv() 读取数据时：

1) 首先会检查缓冲区，如果缓冲区中有数据，那么就读取，否则函数会被阻塞，直到网络上有数据到来。

2) 如果要读取的数据长度小于缓冲区中的数据长度，那么就不能一次性将缓冲区中的所有数据读出，剩余数据将不断积压，直到有 read()/recv() 函数再次读取。

3) 直到读取到数据后 read()/recv() 函数才会返回，否则就一直被阻塞。

这就是TCP套接字的阻塞模式。所谓阻塞，就是上一步动作没有完成，下一步动作将暂停，直到上一步动作完成后才能继续，以保持同步性。
TCP套接字默认情况下是阻塞模式



