## 1. k8s的pause容器有什么用？

 1.提供Pod在Linux中共享命名空间的基础

 2.提供Pid Namespace并使用init进程

## 2. 简述namespace机制

Namespace是对全局系统资源的一种封装隔离，
使得处于不同namespace的进程拥有独立的全局系统资源，
改变一个namespace中的系统资源只会影响当前namespace里的进程，
对其他namespace中的进程没有影响。
隔离的条目很多IPC、Network、Mount、PID、User、UTC，通过Cgroup还可以限制CPU、内存等。面试者能够说明其中的所代表的含义即可。

Mount - isolate filesystem mount points

UTS - isolate hostname and domainname

IPC - isolate interprocess communication (IPC) resources

PID - isolate the PID number space

Network - isolate network interfaces

## 3.Kubernetes有哪些核心组件这些组件负责什么工作？

etcd：提供数据库服务保存了整个集群的状态

kube-apiserver：提供了资源操作的唯一入口，并提供认证、授权、访问控制、API注册和发现等机制

kube-controller-manager：负责维护集群的状态，比如故障检测、自动扩展、滚动更新等

cloud-controller-manager：是与底层云计算服务商交互的控制器

kube-scheduler：负责资源的调度，按照预定的调度策略将Pod调度到相应的机器上

kubelet：负责维护容器的生命周期，同时也负责Volume和网络的管理

kube-proxy：负责为Service提供内部的服务发现和负载均衡，并维护网络规则

container-runtime：是负责管理运行容器的软件，比如docker

## 4.kubenetes针对pod资源对象的健康监测机制

livenessProbe探针

可以根据用户自定义规则来判定pod是否健康，如果livenessProbe探针探测到容器不健康，则kubelet会根据其重启策略来决定是否重启，如果一个容器不包含livenessProbe探针，则kubelet会认为容器的livenessProbe探针的返回值永远成功。

ReadinessProbe探针

同样是可以根据用户自定义规则来判断pod是否健康，如果探测失败，控制器会将此pod从对应service的endpoint列表中移除，从此不再将任何请求调度到此Pod上，直到下次探测成功。

## 5.探针支持探测方法有哪些

Exec：通过执行命令的方式来检查服务是否正常，比如使用cat命令查看pod中的某个重要配置文件是否存在，若存在，则表示pod健康。

Httpget：通过发送http/htps请求检查服务是否正常，返回的状态码为200-399则表示容器健康（注http get类似于命令curl -I）。

tcpSocket：通过容器的IP和Port执行TCP检查，如果能够建立TCP连接，则表明容器健康，这种方式与HTTPget的探测机制有些类似，tcpsocket健康检查适用于TCP业务。

————————————————

版权声明：本文为CSDN博主「土地南瓜」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/hanjiangxue1006/article/details/105111510