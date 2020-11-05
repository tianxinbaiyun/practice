@[TOC](ES和MYSQL数据同步)

## elasticsearch 安装

### docker 安装 elasticsearch-kibana

为简洁软件部署，我们使用集成的docker安装

如果自己单个插件部署，注意elasticsearch，kibana,等服务的版本一致

```text
docker pull eaglechen/go-mysql-elasticsearch
docker pull nshou/elasticsearch-kibana
```

启动docker

```text
docker run -d -p 9200:9200 -p 5601:5601 --name es-kibana nshou/elasticsearch-kibana
docker run -d -v ~/mysql/go_mysql_river.toml:/go_mysql_river.toml:ro --name go-mysql-es eaglechen/go-mysql-elasticsearch
```


## mysqldump工具

mysqldump是一个对mysql数据库中的数据进行全量导出的一个工具.

mysqldump的使用方式如下：

```
mysqldump -u root -p '123456' --host=192.168.15.131 -e --max_allowed_packet=4194304 --net_buffer_length=16384 -F practice > dump.sql
```

上述命令表示从远程数据库192.168.15.131:3306中导出database:practice 的所有数据，
写入到dump.sql文件中，指定-F参数表示在导出数据后重新生成一个新的binlog日志文件以记录后续的所有数据操作。


mysqldump导出的sql文件包含create table, drop table以及插入数据的sql语句，但是不包含create database建库语句。

## 使用go-mysql-elasticsearch开源工具同步数据到ES

go-mysql-elasticsearch是用于同步mysql数据到ES集群的一个开源工具，
项目github地址：https://github.com/siddontang/go-mysql-elasticsearch

go-mysql-elasticsearch的基本原理是：
如果是第一次启动该程序，首先使用mysqldump工具对源mysql数据库进行一次全量同步，
通过elasticsearch client执行操作写入数据到ES；然后实现了一个mysql client,
作为slave连接到源mysql,源mysql作为master会将所有数据的更新操作通过binlog event同步给slave， 
通过解析binlog event就可以获取到数据的更新内容，之后写入到ES.

另外，该工具还提供了操作统计的功能，每当有数据增删改操作时，会将对应操作的计数加1，程序启动时会开启一个http服务，通过调用http接口可以查看增删改操作的次数。

### 使用限制：

```text
1. mysql binlog必须是ROW模式
    
2. 要同步的mysql数据表必须包含主键，否则直接忽略，这是因为如果数据表没有主键，UPDATE和DELETE操作就会因为在ES中找不到对应的document而无法进行同步

3. 不支持程序运行过程中修改表结构

4. 要赋予用于连接mysql的账户RELOAD权限以及REPLICATION权限, SUPER权限：

       GRANT REPLICATION SLAVE ON *.* TO 'root'@'192.168.15.131';
       GRANT RELOAD ON *.* TO 'root'@'192.168.15.131';
       UPDATE mysql.user SET Super_Priv='Y' WHERE user='root' AND host='192.168.15.131';
```


### 使用方式：

#### 克隆工具
```shell script
git clone https://github.com/siddontang/go-mysql-elasticsearch
```

#### 切到工具的目录下
```shell script
cd go-mysql-elasticsearch/src/github.com/siddontang/go-mysql-elasticsearch
```

#### 3.修改配置
vi etc/river.toml, 修改配置文件，同步172.16.0.101:3306数据库中的webservice.building表到ES集群172.16.32.64:9200的building index(更详细的配置文件说明可以参考项目文档)
```shell script
    # MySQL address, user and password
    # user must have replication privilege in MySQL.
    my_addr = "172.16.0.101:3306"
    my_user = "bellen"
    my_pass = "Elastic_123"
    my_charset = "utf8"
    
    # Set true when elasticsearch use https
    #es_https = false
    # Elasticsearch address
    es_addr = "172.16.32.64:9200"
    # Elasticsearch user and password, maybe set by shield, nginx, or x-pack
    es_user = ""
    es_pass = ""
    
    # Path to store data, like master.info, if not set or empty,
    # we must use this to support breakpoint resume syncing.
    # TODO: support other storage, like etcd.
    data_dir = "./var"
    
    # Inner Http status address
    stat_addr = "127.0.0.1:12800"
    
    # pseudo server id like a slave
    server_id = 1001
    
    # mysql or mariadb
    flavor = "mariadb"
    
    # mysqldump execution path
    # if not set or empty, ignore mysqldump.
    mysqldump = "mysqldump"
    
    # if we have no privilege to use mysqldump with --master-data,
    # we must skip it.
    #skip_master_data = false
    
    # minimal items to be inserted in one bulk
    bulk_size = 128
    
    # force flush the pending requests if we don't have enough items >= bulk_size
    flush_bulk_time = "200ms"
    
    # Ignore table without primary key
    skip_no_pk_table = false
    
    # MySQL data source
    [[source]]
    schema = "webservice"
    tables = ["building"]
    [[rule]]
    schema = "webservice"
    table = "building"
    index = "building"
    type = "buildingtype"
```
#### 4.在ES集群中创建building index
在ES集群中创建building index, 因为该工具并没有使用ES的auto create index功能，如果index不存在会报错

#### 5.启动go-mysql-elasticsearch
执行命令：./bin/go-mysql-elasticsearch -config=./etc/river.toml

控制台输出结果：
```shell script

2018/06/02 16:13:21 INFO  create BinlogSyncer with config {1001 mariadb 172.16.0.101 3306 bellen   utf8 false false <nil> false false 0 0s 0s 0}
2018/06/02 16:13:21 INFO  run status http server 127.0.0.1:12800
2018/06/02 16:13:21 INFO  skip dump, use last binlog replication pos (mysql-bin.000001, 120) or GTID %!s(<nil>)
2018/06/02 16:13:21 INFO  begin to sync binlog from position (mysql-bin.000001, 120)
2018/06/02 16:13:21 INFO  register slave for master server 172.16.0.101:3306
2018/06/02 16:13:21 INFO  start sync binlog at binlog file (mysql-bin.000001, 120)
2018/06/02 16:13:21 INFO  rotate to (mysql-bin.000001, 120)
2018/06/02 16:13:21 INFO  rotate binlog to (mysql-bin.000001, 120)
2018/06/02 16:13:21 INFO  save position (mysql-bin.000001, 120)
测试：向mysql中插入、修改、删除数据，都可以反映到ES中
```

#### 6.使用体验
go-mysql-elasticsearch完成了最基本的mysql实时同步数据到ES的功能，业务如果需要更深层次的功能如允许运行中修改mysql表结构，可以进行自行定制化开发。
异常处理不足，解析binlog event失败直接抛出异常
据作者描述，该项目并没有被其应用于生产环境中，所以使用过程中建议通读源码，知其利弊。

作者：bellengao
链接：https://www.jianshu.com/p/c3faa26bc221
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

=========

作者：bellengao
链接：https://www.jianshu.com/p/c3faa26bc221
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。



