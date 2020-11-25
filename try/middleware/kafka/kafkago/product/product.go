package product

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"net"
	"strconv"
)

// 定义
const (
	KafkaConfig    = "localhost:9092"
	KafkaTopic     = "clx_12"
	KafkaPartition = 0
)

var conn *kafka.Conn

//func init() {
//	//var err error
//	//// to produce messages
//	//conn, err = kafka.DialLeader(context.Background(), "tcp", KafkaConfig, KafkaTopic, KafkaPartition)
//	//if err != nil {
//	//	log.Fatal("failed to dial leader:", err)
//	//}
//}

//SimpleProduct SimpleProduct
func SimpleProduct(msg string) (err error) {
	conn, err = kafka.DialLeader(context.Background(), "tcp", KafkaConfig, KafkaTopic, KafkaPartition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}
	//err = conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	//if err != nil {
	//	return
	//}
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte(msg)},
	//kafka.Message{Value: []byte("one!")},
	//kafka.Message{Value: []byte("two!")},
	//kafka.Message{Value: []byte("three!")},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
	return
}

// CreateTopic CreateTopic
func CreateTopic(config kafka.TopicConfig) (err error) {
	// to create topics when auto.create.topics.enable='false'
	//topic := "my-topic"
	//partition := 0

	conn, err := kafka.Dial("tcp", KafkaConfig)
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		panic(err.Error())
	}
	var controllerConn *kafka.Conn
	controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		panic(err.Error())
	}
	defer controllerConn.Close()

	topicConfigs := []kafka.TopicConfig{
		config,
	}

	err = controllerConn.CreateTopics(topicConfigs...)
	if err != nil {
		panic(err.Error())
	}
	return
}

// GetTopicList GetTopicList
func GetTopicList() (list []string, err error) {
	conn, err := kafka.Dial("tcp", KafkaConfig)
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	partitions, err := conn.ReadPartitions()
	if err != nil {
		panic(err.Error())
	}

	m := map[string]struct{}{}

	for _, p := range partitions {
		m[p.Topic] = struct{}{}
	}
	for k := range m {
		list = append(list, k)
		//log.Println(k)
	}
	return
}
