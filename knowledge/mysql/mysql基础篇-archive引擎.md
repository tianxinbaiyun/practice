@[TOC](mysql基础篇-archive引擎)

## 1.archive 引擎特点

Archive引擎作用：为大量很少引用的历史、归档、或安全审计信息的存储和检索提供了完美的解决方案。

#### 优点:

1.**可以压缩**:以zlib对表数据进行压缩，磁盘I/O更少,数据存储在ARZ为后缀的文件中。

2.**存储空间小**:Archive表比MyISAM表要小大约75%，比支持事务处理的InnoDB表小大约83%。

3.**插入数据性能好**:当表内的数据达到1.5GB这个量级，CPU又比较快的时候，Archive表的执行性能就会超越MyISAM表。因为这个时候，CPU会取代I/O子系统成为性能瓶颈。别忘了Archive表比其他任何类型的表执行的物理I/O操作都要少。

4.**数据迁移简单**:Archive表可以方便地移植到新的MySQL环境，你只需将保存Archive表的底层文件复制过去就可以了。

5.**支持行锁**:Archivec存储引擎使用行锁来实现高并发插入操作

6.**支持分区**:Archive存储引擎支持分区

#### 缺点

1.**不支持索引**:archive引擎不支持索引。

2.**不支持更新**:archive引擎不支持更新,删除。

3.**不支持事务**:archive引擎不支持事务。

## 2. 存储

每个archive表在磁盘上存在两个文件:.frm(存储表定义).arz(存储数据)

往archive表插入的数据会经过压缩，archive使用zlib进行数据压缩，archive支持optimize table、 check table操作。

一个insert语句仅仅往压缩缓存中插入数据，插入的数据在压缩缓存中被锁定，当select操作时会触发压缩缓存中的数据进行刷新。insert delay除外。

对于一个bulk insert操作只有当它完全执行完才能看到记录，除非在同一时刻还有其它的inserts操作，在这种情况下可以看到部分记录，select从不刷新bulk insert除非在它加载时存在一般的Insert操作。

## 3.索引

1.archive存储引擎支持insert、replace和select操作，但是不支持update和delete。

2.archive存储引擎支持blob、text等大字段类型。支持auto_increment自增列同时自增列可以不是唯一索引。

3.archive支持auto_increment列，但是不支持往auto_increment列插入一个小于当前最大的值的值。

4.archive不支持索引所以无法在archive表上创建主键、唯一索引、和一般的索引。

## 4.innodb , myisam 和 archive 的测试

测试机器配置为:6核 cpu ,8G 内存 ,Ubuntu 64位系统

数据库结构如下
```mysql
CREATE TABLE `logs_archive` (
  `name` varchar(255) DEFAULT NULL,
  `time` bigint(20) DEFAULT NULL,
  `creat_time` bigint(20) DEFAULT NULL
) ENGINE=ARCHIVE DEFAULT CHARSET=utf8;

CREATE TABLE `logs_innodb` (
  `name` varchar(255) DEFAULT NULL,
  `time` bigint(20) DEFAULT NULL,
  `creat_time` bigint(20) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `logs_myisam` (
  `name` varchar(255) DEFAULT NULL,
  `time` bigint(20) DEFAULT NULL,
  `creat_time` bigint(20) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

```


#### 批量添加测试
以一千万条数据为例,golang语言环境运行,使用orm方式,开启100个协程插入:

**InnoDB**
 
使用时间为:1116.403146241秒

存储数据长度:960.00 MB


**MyISAM**

使用时间为: 888.011918396秒

存储数据长度:572.20 MB


**ARCHIVE**

使用时间为: 785.65768241秒

存储数据长度:207.87 MB 


**总结**:
InnoDB,MyISAM,ARCHIVE三个存储引擎,批量插入速度,ARCHIVE最快,MyISAM次之,InnoDB最慢;
存储数据长度ARCHIVE最小,MyISAM次之,InnoDB最大;

#### 删除测试

DELETE FROM logs_innodb WHERE `name` ='new_test'
> Affected rows: 1
> 时间: 4.128s

DELETE FROM logs_myisam WHERE `name` ='new_test'
> Affected rows: 1
> 时间: 1.1s

DELETE FROM logs_archive WHERE `name` ='new_test'
> 1031 - Table storage engine for 'logs_archive' doesn't have this option
> 时间: 5.232s

**总结**:
InnoDB,MyISAM,ARCHIVE三个存储引擎,删除性能,MyISAM最快,InnoDB次之,ARCHIVE最慢且报错;

#### 更改测试

UPDATE logs_innodb SET `name`='new_test2' WHERE `name` =  'new_test'
> Affected rows: 1
> 时间: 4.096s

UPDATE logs_myisam SET `name`='new_test2' WHERE `name` =  'new_test'
> Affected rows: 1
> 时间: 19.741s

UPDATE logs_archive SET `name`='new_test2' WHERE `name` =  'new_test'
> 1031 - Table storage engine for 'logs_archive' doesn't have this option
> 时间: 5.243s

**总结**:
InnoDB,MyISAM,ARCHIVE三个存储引擎,更改操作性能,InnoDB 最快,MyISAM 最慢,ARCHIVE报错;


#### 查询测试

**条件查询**
SELECT * FROM logs_innodb WHERE creat_time = '1579414712'  LIMIT 1
>时间: 0.003s

SELECT * FROM logs_myisam  WHERE creat_time = '1579416525'  LIMIT 1

> 时间: 0.001s

SELECT * FROM logs_archive WHERE creat_time = '1579418376' LIMIT 1
> 时间: 0.01s


SELECT * FROM logs_innodb WHERE creat_time = '1579414712'  LIMIT 100000
> 时间: 3.019s

SELECT * FROM logs_myisam  WHERE creat_time = '1579416525'  LIMIT 100000
> 时间: 0.79s

SELECT * FROM logs_archive WHERE creat_time = '1579418376' LIMIT 100000
> 时间: 4.911s


**使用count查询条数**

SELECT COUNT(*) FROM logs_innodb;
> 时间: 2.663s

SELECT COUNT(*) FROM logs_myisam;
> 时间: 0s
 
SELECT COUNT(*) FROM logs_archive;
> 时间: 0s

SELECT COUNT(*) FROM logs_innodb WHERE `creat_time` = '1579414712'
> 时间: 2.758s

SELECT COUNT(*) FROM logs_myisam WHERE `creat_time` = '1579416525'
> 时间: 0.774s


SELECT COUNT(*) FROM logs_archive WHERE `creat_time` = '1579418376'
> 时间: 4.927s


**总结**:
InnoDB,MyISAM,ARCHIVE三个存储引擎,单行数据查询操作性能, MyISAM 最快,InnoDB 次之,ARCHIVE最慢;

文章参考网上资料,如有侵权请联系删除