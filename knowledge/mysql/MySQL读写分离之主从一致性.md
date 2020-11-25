@[TOC](数据库读写分离一致性问题)

MySQL一主多从时如何保证从库读到的数据是最新的？

## 数据库读写分离一致性问题

说说思路

### 1 半同步复制

  简单的说就是: 主库发生增删改操作的时候,
  会等从库及时复制了并且通知了主库, 才会把这个操作叫做成功.

  优点：保证数据一致性

  缺点:就是会慢

专业的讲：

  半同步复制，是等待其中一个从库也接收到Binlog事务并成功写入Relay Log之后，
  才返回Commit操作成功给客户端；
  如此半同步就保证了事务成功提交后至少有两份日志记录，
  一份在主库Binlog上，另一份在从库的Relay Log上，
  从而进一步保证数据完整性；
  
  半同步复制很大程度取决于主从网络RTT（往返时延），
  以插件 semisync_master/semisync_slave 形式存在。 

mysql具体如何配置可参考

https://www.cnblogs.com/zero-gg/p/9057092.html

mqspl半同步存在的一些问题

https://www.cnblogs.com/jichunhu/p/5825801.html

 

### 2 全同步复制

介绍了半同步，全同步概念也就不言而喻啦。


### 3 数据库中间件

如果有了数据库中间件，所有的数据库请求都走中间件，这个主从不一致的问题可以这么解决：

所有的读写请求都走中间件，然后写的请求路由到主库，读的请求路由到从库

但是我们中间件会记录写库的一个key,在设置一个允许同步时间,假设是1s

当有一个写请求过来时候，生成一个key A ，
马上路由写到主库，然后立马有一个读请求过来。 
从库可能是旧数据，或者没有来得及同步。 如果时间是在1s内的，
就对应的key继续路由到主库。如果在1s以后的，就路由到从库。

说白了，中间件就是给个同步时间，给你同步，在同步时间内，所有的请求都落在主库

不知有没有可用的中间件，不过实现相对简单，可以用本地缓存来实现，
通过查找本地缓存来判断数据是否刚写入。


虽然MySQL一直在优化数据的一致性问题，但问题依然存在，
使得各大企业纷纷各自设计一套MySQL补丁来保证数据一致。
腾讯数平的TDSQL，腾讯微信的PhxSQL，阿里的AliSQL，
网易的InnoSQL等设计都是为了保证数据一致性。
MySQL5.7发布的lossless半同步，虽然宣称zero loss，
解决了5.6版本中有可能出现的data lost问题，但其数据一致性仍未完全解决。

cluster 方案

MySQL异步复制

在MySQL发展的早期，就提供了异步复制的技术，只要写的压力不是特别大，在网络条件较好的情况下，发生主备切换基本上能将影响控制到秒级别，因此吸引了很多开发者的关注和使用。但这套方案提供的一致性保证，对于计费或者金融行业是不够的。

图4是异步复制的大致流程，很显然主机提交了binlog就会返回给业务成功，没有保证binlog同步到了备机，这样在切换的瞬间很有可能丢失这部分事务。

 



 

图4 异步复制


https://blog.csdn.net/weixin_40735752/article/details/88074704
https://blog.csdn.net/nawenqiang/article/details/85161847?depth_1-utm_source=distribute.pc_relevant.none-task-blog-BlogCommendFromBaidu-1&utm_source=distribute.pc_relevant.none-task-blog-BlogCommendFromBaidu-1