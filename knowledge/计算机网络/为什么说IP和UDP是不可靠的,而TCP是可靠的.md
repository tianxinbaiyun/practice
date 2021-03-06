@[TOC](为什么说IP和UDP是不可靠的,而TCP是可靠的)

## 名词定义

TCP/IP（Transmission Control Protocol/Internet Protocol，传输控制协议/网际协议）是指能够在多个不同网络间实现信息传输的协议簇。
TCP/IP协议不仅仅指的是TCP 和IP两个协议，而是指一个由FTP、SMTP、TCP、UDP、IP等协议构成的协议簇， 只是因为在TCP/IP协议中TCP协议和IP协议最具代表性，所以被称为TCP/IP协议。

TCP:传输控制协议（TCP，Transmission Control Protocol）是一种面向连接的、可靠的、基于字节流的传输层通信协议，由IETF的RFC 793 [1]  定义。
TCP旨在适应支持多网络应用的分层协议层次结构。 连接到不同但互连的计算机通信网络的主计算机中的成对进程之间依靠TCP提供可靠的通信服务。

IP是Internet Protocol（网际互连协议）的缩写，是TCP/IP体系中的网络层协议。
设计IP的目的是提高网络的可扩展性：一是解决互联网问题，实现大规模、异构网络的互联互通；二是分割顶层网络应用和底层网络技术之间的耦合关系，以利于两者的独立发展。根据端到端的设计原则，IP只为主机提供一种无连接、不可靠的、尽力而为的数据报传输服务。

UDP:该协议称为用户数据报协议（UDP，User Datagram Protocol）。UDP 为应用程序提供了一种无需建立连接就可以发送封装的 IP 数据包的方法。

## 区别总结

### 通信方式
TCP是面向连接的。UDP是面向无连接的。

在网络中，有些服务，如HTTP、FTP等，对数据的可靠性要求较高，在使用这些服务时，必须保证数据包能够完整无误的送达;
而另外一些服务，如DNS、即时聊天工具等，并不需要这么高的可靠性，高效率和实时性才是它们所关心的。
根据这两种服务不同的需求，也就诞生了面向连接的TCP协议，以及面向无连接的UDP协议。

面向连接的协议比面向无连接的协议在可靠性上有着显著的优势，
但建立连接前必须等待接收方响应，传输信息过程中必须确认信息是否传到，断开连接时需要发出响应信号等，无形中加大了面向连接协议的资源开销。
具体到TCP和UDP协议来说，除了源端口和目的端口，TCP还包括序号、确认信号、数据偏移、控制标志(通常说的URG、ACK、PSH、RST、SYN、FIN)、窗口、校验和、紧急指针、选项等信息.
UDP则只包含长度和校验和信息。
UDP数据报比TCP小许多，这意味着更小的负载和更有效的使用带宽。许多即时聊天软件采用UDP协议，与此有莫大的关系。

### 传输数据模式
TCP流模式:
你通过TCP连接给另一端发送数据，你只调用了一次write，发送了100个字节，但是对方可以分10次收完，每次10个字节；你也可以调用10次write，每次10个字节，但是对方可以一次就收完。
（假设数据都能到达）但是，你发送的数据量不能大于对方的接收缓存（流量控制），如果你硬是要发送过量数据，则对方的缓存满了就会把多出的数据丢弃。

UDP包模式:
发送端调用了几次write，接收端必须用相同次数的read读完。UPD是基于报文的，在接收的时候，每次最多只能读取一个报文，报文和报文是不会合并的，如果缓冲区小于报文长度，则多出的部分会被丢弃。
也就说，如果不指定MSG_PEEK标志，每次读取操作将消耗一个报文。

### 包头格式
TCP包头需要24个字节,共192bit;UDP包头只需要8个字节,共64bit;
TCP和UDP都包含源地址,目标地址,头部长度,校验和四个字段,UDP只有这四个字段;
TCP的头部长度占用4bit,而UDP的头部长度占用16bit;
TCP比UDP多的字段:顺序号,确认号,标志位,窗口大小,紧急指针,选填字段


#### TCP包头结构:

<table>
    <tr>
        <td colspan="16">源端口</td>
        <td colspan="16">目标端口</td>
    </tr>
    <tr><td colspan="32">顺序号</td></tr>
    <tr><td  colspan="32">确认号</td></tr>
    <tr>
        <td  colspan="4">头部长度</td>
        <td  colspan="6">保留位</td>
        <td  colspan="6">标志位</td>
        <td  colspan="16">窗口大小</td>
    </tr>
    <tr>
        <td colspan="16">TCP校验和</td>
        <td colspan="16">紧急指针</td>
    </tr>
    <tr><td  colspan="32">选项字段</td></tr>
</table>

**源、目标端口号字段**：占16比特。TCP协议通过使用"端口"来标识源端和目标端的应用进程。端口号可以使用0到65535之间的任何数字。在收到服务请求时，操作系统动态地为客户端的应用程序分配端口号。在服务器端，每种服务在"众所周知的端口"（Well-Know Port）为用户提供服务。

**顺序号字段**：占32比特。用来标识从TCP源端向TCP目标端发送的数据字节流，它表示在这个报文段中的第一个数据字节。

**确认号字段**：占32比特。只有ACK标志为1时，确认号字段才有效。它包含目标端所期望收到源端的下一个数据字节。

**头部长度字段**：占4比特。给出头部占32比特的数目。没有任何选项字段的TCP头部长度为20字节；最多可以有60字节的TCP头部。

标志位字段（U、A、P、R、S、F）：占6比特。各比特的含义如下：
<table>
    <tr><td>URG</td><td>紧急指针（urgent pointer）有效。</td></tr>
    <tr><td>ACK</td><td>确认序号有效。</td></tr>
    <tr><td>PSH</td><td>接收方应该尽快将这个报文段交给应用层。</td></tr>
    <tr><td>RST</td><td>重建连接。</td></tr>
    <tr><td>SYN</td><td>发起一个连接。</td></tr>
    <tr><td>FIN</td><td>释放一个连接。</td></tr>
</table>

**窗口大小字段**：占16比特。此字段用来进行流量控制。单位为字节数，这个值是本机期望一次接收的字节数。

**TCP校验和字段**：占16比特。对整个TCP报文段，即TCP头部和TCP数据进行校验和计算，并由目标端进行验证。

**紧急指针字段**：占16比特。它是一个偏移量，和序号字段中的值相加表示紧急数据最后一个字节的序号。

**选项字段**：占32比特。可能包括"窗口扩大因子"、"时间戳"等选项。

#### UDP包头结构:

<table>
    <tr>
        <td colspan="16">源端口</td>
        <td colspan="16">目标端口</td>
    </tr>
    <tr>
        <td colspan="16">长度</td>
        <td colspan="16">校验和</td>
    </tr>
</table>

**源、目标端口号字段**：占16比特。TCP协议通过使用"端口"来标识源端和目标端的应用进程。端口号可以使用0到65535之间的任何数字。在收到服务请求时，操作系统动态地为客户端的应用程序分配端口号。在服务器端，每种服务在"众所周知的端口"（Well-Know Port）为用户提供服务。

**长度字段**：占16比特。标明UDP头部和UDP数据的总长度字节。

**校验和字段**：占16比特。用来对UDP头部和UDP数据进行校验。和TCP不同的是，对UDP来说，此字段是可选项，而TCP数据段中的校验和字段是必须有的

### 连接的可靠性

#### 顺序处理
TCP保证数据顺序，UDP不保证:

TCP为了保证顺序性，每个包都有一个 ID。在建立连接的时候会商定起始 ID 是什么，然后按照 ID 一个个发送，为了保证不丢包，需要对发送的包都要进行应答，当然，这个应答不是一个一个来的，而是会应答某个之前的 ID，表示都收到了，这种模式成为累计应答或累计确认。

为了记录所有发送的包和接收的包，TCP 需要发送端和接收端分别来缓存这些记录，发送端的缓存里是按照包的 ID 一个个排列，根据处理的情况分成四个部分

发送并且确认的 / 发送尚未确认的 / 没有发送等待发送的 / 没有发送并且暂时不会发送的

**那么，TCP具体是通过怎样的方式来保证数据的顺序化传输呢？**

1. 主机每次发送数据时，TCP就给每个数据包分配一个序列号并且在一个特定的时间内等待接收主机对分配的这个序列号进行确认，

2. 如果发送主机在一个特定时间内没有收到接收主机的确认，则发送主机会重传此数据包。

3. 接收主机利用序列号对接收的数据进行确认，以便检测对方发送的数据是否有丢失或者乱序等，

4. 接收主机一旦收到已经顺序化的数据，它就将这些数据按正确的顺序重组成数据流并传递到高层进行处理。

**具体步骤如下：**

（1）为了保证数据包的可靠传递，发送方必须把已发送的数据包保留在缓冲区； 

（2）并为每个已发送的数据包启动一个超时定时器； 

（3）如在定时器超时之前收到了对方发来的应答信息（可能是对本包的应答，也可以是对本包后续包的应答），则释放该数据包占用的缓冲区; 

（4）否则，重传该数据包，直到收到应答或重传次数超过规定的最大次数为止。

（5）接收方收到数据包后，先进行CRC校验，如果正确则把数据交给上层协议，然后给发送方发送一个累计应答包，表明该数据已收到，如果接收方正好也有数据要发给发送方，应答包也可方在数据包中捎带过去。

#### 掉包处理
TCP保证数据正确性，UDP可能丢包:

TCP是一个“流”协议，一个完整的包可能被TCP拆分成多个包发送，也有可能把小的封装成大的发送。通过滑动窗口保证数据的顺序可靠,防止掉包.

假设客户端分别发送数据包D1和D2给服务端，由于服务端一次性读取到的字节数是不确定的，所以可能存在以下4种情况。

1.服务端分2次读取到了两个独立的包，分别是D1,D2,没有粘包和拆包；

2.服务端一次性接收了两个包，D1和D2粘在一起了，被成为TCP粘包;

3.服务端分2次读取到了两个数据包，第一次读取到了完整的D1和D2包的部分内容,第二次读取到了D2包的剩余内容，这被称为拆包；

4.如果此时服务端TCP接收滑动窗口非常小,而数据包D1和D2都很大，很有可能发送第五种可能，即服务端多次才能把D1和D2接收完全，期间多次发生拆包情况。

接收滑动窗：所谓滑动窗口协议，自己理解有两点：

1. “窗口”对应的是一段可以被发送者发送的字节序列，其连续的范围称之为“窗口”；

2. “滑动”则是指这段“允许发送的范围”是可以随着发送的过程而变化的，方式就是按顺序“滑动”

由于底层的TCP无法理解上层的业务逻辑，所以在底层是无法确保数据包不被拆分和重组的，这个问题只能通过上层的应用协议栈设计来解决，根据业界的主流协议的解决方案，归纳如下：

消息定长，例如每个报文的大小为固定长度200字节,如果不够，空位补空格

在包尾增加回车换行符进行分割，例如FTP协议

将消息分为消息头和消息体，消息头中包含表示消息总长度（或者消息体长度）的字段，通常设计思路是消息头的第一个字段用int来表示消息的总长度
更复杂的应用层协议

#### 流量控制

在流量控制的机制里面，在对于包的确认中，会携带一个窗口的大小

简单的说一下就是接收端在发送 ACK 的时候会带上缓冲区的窗口大小，但是一般在窗口达到一定大小才会更新窗口，
因为每次都更新的话，刚空下来就又被填满了

#### 拥塞控制

TCP 拥塞控制主要来避免两种现象，包丢失和超时重传，一旦出现了这些现象说明发送的太快了，要慢一点。

具体的方法就是发送端慢启动，比如倒水，刚开始倒的很慢，渐渐变快。然后设置一个阈值，当超过这个值的时候就要慢下来

慢下来还是在增长，这时候就可能水满则溢，出现拥塞，需要降低倒水的速度，等水慢慢渗下去。

拥塞的一种表现是丢包，需要超时重传，这个时候，采用快速重传算法，将当前速度变为一半。所以速度还是在比较高的值，也没有一夜回到解放前。





————————————————

版权声明：本文为CSDN博主「Object object」的原创文章，遵循 CC 4.0 BY-SA 版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/zhang6223284/article/details/81414149#24__124


