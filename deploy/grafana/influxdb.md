@[TOC](cAdvisor+InfluxDB+Grafana 监控Docker)

## 1 介绍
### 1.1. InfluxDB是什么
influxDB是用GO语言编写的一个开源分布式时序、事件和指标数据库，无需外部的依赖，类似的数据库有Elasticsearch、Graphite等等
 
InfluxDB主要的功能：
:基于时间序列：支持与时间有关的相关函数(如最大、最小、求和等)
:可度量性：可以实时对大量数据进行计算
:基于事件：它支持任意的事件数据
 
InfluxDB的主要特点：
:无结构(无模式):可以是任意数量的列
:可拓展的
:支持min, max, sum, count, mean, median 等一系列函数，方便统计
:原生的HTTP支持，内置HTTP API
:强大的类SQL语法
:自带管理界面，方便使用
 
### 1.2. cAdvisor是什么
它是Google用来监测单节点的资源信息的监控工具。
Cadvisor提供了一目了然的单节点多容器的资源监控功能。
Google的Kubernetes中也缺省地将其作为单节点的资源监控工具，各个节点缺省会被安装上Cadvisor
:cAvisor是利用docker status的数据信息，了解运行时容器资源使用和性能特征的一种工具
:cAdvisor的容器抽象基于Google的lmctfy容器栈，因此原生支持Docker容器并能够“开箱即用”地支持其他的容器类型。
:cAdvisor部署为一个运行中的daemon，它会收集、聚集、处理并导出运行中容器的信息。
:这些信息能够包含容器级别的资源隔离参数、资源的历史使用状况、反映资源使用和网络统计数据完整历史状况的柱状图。
 
cAdvisor功能：
:展示Host和容器两个层次的监控数据
:展示历史变化数据
 
温馨提示：
由于 cAdvisor 提供的操作界面略显简陋，而且需要在不同页面之间跳转，并且只能监控一个 host，这不免会让人质疑它的实用性。
但 cAdvisor 的一个亮点是它可以将监控到的数据导出给第三方工具，由这些工具进一步加工处理。
我们可以把 cAdvisor 定位为一个监控数据收集器，收集和导出数据是它的强项，而非展示数据
 
### 1.3. Grafana是什么
Grafana是一个可视化面板（Dashboard），有着非常漂亮的图表和布局展示，功能齐全的度量仪表盘和图形编辑器，支持Graphite、zabbix、InfluxDB、Prometheus和OpenTSDB作为数据源
 
Grafana主要特性：
:灵活丰富的图形化选项；
:可以混合多种风格；
:支持白天和夜间模式；
:支持多个数据源；
 
温馨提示：在这套监控方案中：InfluxDB用于数据存储，cAdvisor用户数据采集，Grafana用于数据展示


## 2.单节点部署

### 2.1.下载镜像
1. 下载镜像(可做可不做，在创建容器的时候会如果本地没有会自动下载)
```
下载镜像
[root@master1 ~]# docker pull tutum/influxdb
[root@master1 ~]# docker pull google/cadvisor
[root@master1 ~]# docker pull grafana/grafana
 
查看镜像
[root@master1 ~]# docker images
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
grafana/grafana     latest              7038dbc9a50c        7 days ago          223MB
google/cadvisor     latest              75f88e3ec333        10 months ago       62.2MB
influxdb            latest              c061e5808198        2 years ago   
```

### 2.2.创建influxDB
```shell script
docker run -itd -p 18083:8083 -p 18086:8086 --name influxdb tutum/influxdb
```

登录URL：http://192.168.15.131:18083

设置管理员用户名密码，并添加数据库

### 2.3 创建cadvisor容器

```shell script
docker run -itd --name cadvisor -p 18080:8080 \
--mount type=bind,src=/,dst=/rootfs,ro \
--mount type=bind,src=/var/run,dst=/var/run \
--mount type=bind,src=/sys,dst=/sys,ro \
--mount type=bind,src=/var/lib/docker/,dst=/var/lib/docker,ro google/cadvisor \
-storage_driver=influxdb \
-storage_driver_db=cadvisor \
-storage_driver_user=admin \
-storage_driver_password=admin \
-storage_driver_host=172.16.4.215:18086
 
参数详解：
-itd：已交互模式运行容器，并分配伪终端，并在后台启动容器
-p: 端口映射 18080为cadvisor的管理平台端口
--name：给容器起个名字
--mout：把宿主机的相文目录绑定到容器中，这些目录都是cadvisor需要采集的目录文件和监控内容
google/cadvisor：以这个镜像运行容器(本地有使用本地，没有先去下载然后启动容器)
-storage_driver：需要指定cadvisor的存储驱动这里是influxdb
-storage_driver_db：需要指定存储的数据库
-storage_driver_user：influxdb数据库的用户名(测试可以加可以不加)
-storage_driver_password：influxdb数据库的密码(测试可以加可以不加)
-storage_driver_host：influxdb数据库的地址和端口
 
# 查看容器
[root@master1 ~]# docker ps
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                                            NAMES
7c2005bb79d1        google/cadvisor     "/usr/bin/cadvisor -…"   3 seconds ago       Up 2 seconds        0.0.0.0:18080->8080/tcp                           cadvisor
2fa150d3c52b        tutum/influxdb      "/run.sh"                10 minutes ago      Up 10 minutes       0.0.0.0:18083->8083/tcp, 0.0.0.0:8086->8086/tcp   influxdb

```

查看cadvisor管理平台
登录URL：http://192.168.15.129:18080

登录数据库查看有没有把采集的数据写入(执行这个命令)
```shell script
SHOW MEASUREMENTS
```

### 2.4. 创建grafana容器

```
# 创建grafana容器
[root@master1 ~]# docker run -itd --name grafana  -p 13000:3000 grafana/grafana
 
参数详解：
-itd：已交互模式运行容器，并分配伪终端，并在后台启动容器
-p: 端口映射 3000为grafana的管理平台端口
--name：给容器起个名字
grafana/grafana：以这个镜像运行容器(本地有使用本地，没有先去下载然后启动容器)
 
# 查看容器
[root@master1 ~]# docker ps
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                                            NAMES
57f335665902        grafana/grafana     "/run.sh"                2 seconds ago       Up 1 second         0.0.0.0:13000->3000/tcp                           grafana
7c2005bb79d1        google/cadvisor     "/usr/bin/cadvisor -…"   15 minutes ago      Up 15 minutes       0.0.0.0:18080->8080/tcp                           cadvisor
2fa150d3c52b        tutum/influxdb      "/run.sh"                25 minutes ago      Up 25 minutes       0.0.0.0:18083->8083/tcp, 0.0.0.0:8086->8086/tcp   influxdb

```


## 3.Swarm多节点部署
刚刚上面的例子是在一台主机上监控一台主机的容器信息，这里我们要使用Swarm的集群部署多台主机容器之间的监控
温馨提示：
:主机IP：192.168.15.129 主机名：master1 角色：Swarm的主 granfana容器 influxdb容器 cadvisor容器
:主机IP：192.168.15.130 主机名：node1 角色：Swarm的node节点 cadvisor容器
:主机IP：192.168.15.131 主机名：node2 角色：Swarm的node节点 cadvisor容器

### 3.1. 准备工作

#### 3.1.1 创建InfluxDB的宿主机目录挂载到容器
```shell script
mkdir -p /opt/influxdb
```

 
#### 3.1.2下载镜像(可做可不做，在创建容器的时候会如果本地没有会自动下载)
[root@master1 ~]# docker pull tutum/influxdb
[root@master1 ~]# docker pull google/cadvisor
[root@master1 ~]# docker pull grafana/grafana
 
#### 3.1.3 查看镜像
[root@master1 ~]# docker images
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
grafana/grafana     latest              7038dbc9a50c        7 days ago          223MB
google/cadvisor     latest              75f88e3ec333        10 months ago       62.2MB
tutum/influxdb      latest              c061e5808198        2 years ago         290MB


### 3.2 编写docker-compose.yml文件
```shell script


[root@master1 ~]# mkdir test
[root@master1 test]# cat docker-compose.yml
version: '3.7'
 
services:
  influx:
    image: tutum/influxdb
    ports:
      - "8083:8083"
      - "8086:8086"
    volumes:
      - "/opt/influxdb:/var/lib/influxdb"
    deploy:
      replicas: 1
      placement:
        constraints: [node.role==manager]
 
  grafana:
    image: grafana/grafana
    ports:
      - "3000:3000"
    depends_on:
      - "influx"
    deploy:
      replicas: 1
      placement:
        constraints: [node.role==manager]
 
  cadvisor:
    image: google/cadvisor
    ports:
      - "8080:8080"
    hostname: '{{.Node.Hostname}}'
    command: -logtostderr -docker_only -storage_driver=influxdb -storage_driver_db=cadvisor -storage_driver_host=influx:8086
    volumes:
      - /:/rootfs:ro
      - /var/run:/var/run:rw
      - /sys:/sys:ro
      - /var/lib/docker/:/var/lib/docker:ro
    depends_on:
      - influx
    deploy:
      mode: global
 
volumes:
  influx:
    driver: local
  grafana:
    driver: local
```
### 3.3. 创建Swarm集群

#### 3.3.1 在master1上执行
```shell script
[root@master1 test]# docker swarm init --advertise-addr 192.168.15.129
```

Swarm initialized: current node (xtooqr30af6fdcu51jzdv79wh) is now a manager.
 
To add a worker to this swarm, run the following command:
    # 这里已经提示使用下面的命令在node节点上执行就可以加入集群(前提docker服务一定是启动的)
    docker swarm join --token SWMTKN-1-3yyjydabd8v340kptius215s29rbsq8tviy00s08g6md1y25k2-81tp7lpv114a393g4wlgx4a30 192.168.15.129:2377
 
To add a manager to this swarm, run 'docker swarm join-token manager' and follow the instructions.
 
 
#### 3.3.2 在node1和node2上执行
```shell script
[root@node1 ~]# docker swarm join --token SWMTKN-1-3yyjydabd8v340kptius215s29rbsq8tviy00s08g6md1y25k2-81tp7lpv114a393g4wlgx4a30 192.168.15.129:2377
This node joined a swarm as a worker
 
[root@node2 ~]# docker swarm join --token SWMTKN-1-3yyjydabd8v340kptius215s29rbsq8tviy00s08g6md1y25k2-81tp7lpv114a393g4wlgx4a30 192.168.15.129:2377
This node joined a swarm as a worker.
```

 
#### 3.3.3 在master1上查看集群主机
```shell script
[root@master1 test]# docker node ls
ID                            HOSTNAME            STATUS              AVAILABILITY        MANAGER STATUS      ENGINE VERSION
xtooqr30af6fdcu51jzdv79wh *   master1             Ready               Active              Leader              18.06.1-ce
y24c6sfs3smv5sd5h7k66x8zv     node1               Ready               Active                                  18.06.1-ce
k554xe59lcaeu1suaguvxdnel     node2               Ready               Active                                  18.06.1-ce
```

### 3.4. 创建集群容器

创建集群容器
```shell script
[root@master1 test]# docker stack deploy -c docker-compose.yml swarm-monitor
Creating network swarm-monitor_default
Creating service swarm-monitor_cadvisor
Creating service swarm-monitor_influx
Creating service swarm-monitor_grafana
```

查看创建的容器
```shell script
[root@master1 test]# docker service  ls
ID                  NAME                     MODE                REPLICAS            IMAGE                    PORTS
wn36f7be6i5a        swarm-monitor_cadvisor   global              3/3                 google/cadvisor:latest   *:8080->8080/tcp
ufn3lqbhbww3        swarm-monitor_grafana    replicated          1/1                 grafana/grafana:latest   *:3000->3000/tcp
lf0z6dp1u8sn        swarm-monitor_influx     replicated          1/1                 tutum/influxdb:latest    *:8083->8083/tcp, *:8086->8086/tcp
```

查看容器的服务
```shell script
[root@master1 test]# docker service ps swarm-monitor_cadvisor
ID                  NAME                                               IMAGE                    NODE                DESIRED STATE       CURRENT STATE                ERROR               PORTS
vy1kqg5u8x3f        swarm-monitor_cadvisor.k554xe59lcaeu1suaguvxdnel   google/cadvisor:latest   node2               Running             Running about a minute ago                      
a08b5bysra3d        swarm-monitor_cadvisor.y24c6sfs3smv5sd5h7k66x8zv   google/cadvisor:latest   node1               Running             Running about a minute ago                      
kkca4kyojgr2        swarm-monitor_cadvisor.xtooqr30af6fdcu51jzdv79wh   google/cadvisor:latest   master1             Running             Running 59 seconds ago  
 
[root@master1 test]# docker service ps swarm-monitor_grafana
ID                  NAME                      IMAGE                    NODE                DESIRED STATE       CURRENT STATE                ERROR               PORTS
klyjl7rxzmoz        swarm-monitor_grafana.1   grafana/grafana:latest   master1             Running             Running about a minute ago       
 
[root@master1 test]# docker service ps swarm-monitor_influx
ID                  NAME                     IMAGE                   NODE                DESIRED STATE       CURRENT STATE                ERROR               PORTS
pan5yvwq7b79        swarm-monitor_influx.1   tutum/influxdb:latest   master1             Running             Running about a minute ago    
```

### 3.5. 访问web测试
#### 3.5.1 访问influxdb并创建数据库
登录InfluxDB的8083端口，并添加数据库
登录URL：http://192.168.15.129:8083



#### 3.5.2 访问cadvisor
登录URL：http://192.168.15.129:8080
登录数据库查看有没有把采集的数据写入



#### 3.5.3 访问grafana并配置
登录URL：http://192.168.15.129:3000
默认用户名:admin
默认密码：admin
温馨提示：
首次登录会提示修改密码才可以登录，我这里修改密码为admin

这个动图比较长 主要是对grafana的配置操作，注意里面的alpine_test容器不是和集群一块创建的是我单独创建的　　


## 4. 创建报警邮件提醒



===========
文章参考：一本正经的搞事情 
https://www.cnblogs.com/zhujingzhi/p/9844558.html