## docker简单安装kafka

### 安装文档

docker-compose.yaml

```
version: '3'
services:
  zookeeper:
    image: wurstmeister/zookeeper
    ports:
      - "2181:2181"
  kafka:
    image: wurstmeister/kafka
    depends_on: [ zookeeper ]
    ports:
      - "9092:9092"
    environment:
      KAFKA_ADVERTISED_HOST_NAME: localhost
      KAFKA_CREATE_TOPICS: "test:1:1"
      KAFKA_ZOOKEEPER_CONNECT: zookeeper:2181
    volumes:
      - ./data/kafka/logs:/kafka
      - ./data/kafka/docker.sock:/var/run/docker.sock

  kafka-manager:
    image: sheepkiller/kafka-manager:latest
    container_name: kafa-manager
    hostname: kafka-manager
    ports:
      - "9000:9000"
    links:            # 连接本compose文件创建的container
      - kafka
    external_links:   # 连接本compose文件以外的container
      - zookeeper
    environment:
      ZK_HOSTS: zookeeper:2181
      KAFKA_BROKERS: kafka:9092
      APPLICATION_SECRET: letmein
      KM_ARGS: -Djava.net.preferIPv4Stack=true
```

### 验证安装
进入容器
```
$ docker exec -it kafka_kafka_1 bash
```


进入 /opt/kafka_2.13-2.6.0/bin 目录下
```
$ cd /opt/kafka_2.13-2.6.0/bin
```

运行kafka消费者监听消息
```
$ ./kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic sun --from-beginning
``` 
 
运行kafka生产者发送消息
```
$ ./kafka-console-producer.sh --broker-list localhost:9092 --topic sun
 ```
发送消息
```
> hello word!
```

### 进入kafka-manage管理

kafka-manage管理参考：https://www.jianshu.com/p/6a592d558812
