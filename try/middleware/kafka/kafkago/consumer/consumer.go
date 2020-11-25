package consumer

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

// 定义
const (
	KafkaConfig    = "localhost:9092"
	KafkaTopic     = "clx_12"
	KafkaPartition = 0
)

//SimpleConsumer SimpleConsumer
func SimpleConsumer() (err error) {
	// to consume messages

	conn, err := kafka.DialLeader(context.Background(), "tcp", KafkaConfig, KafkaTopic, KafkaPartition)
	if err != nil {
		log.Printf("failed to dial leader:%v", err)
	}

	err = conn.SetReadDeadline(time.Now().Add(3 * time.Second))
	if err != nil {
		log.Printf("err:%v", err)
		return
	}
	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	b := make([]byte, 10e3) // 10KB max per message
	for {
		_, err := batch.Read(b)
		if err != nil {
			break
		}
		log.Println(string(b))
	}

	if err := batch.Close(); err != nil {
		log.Printf("failed to close batch:%v", err)
	}

	if err := conn.Close(); err != nil {
		log.Printf("failed to close connection:%v", err)
	}
	return
}

// GetReader GetReader
func GetReader() (err error) {
	// make a new reader that consumes from topic-A, partition 0, at offset 42
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{KafkaConfig},
		Topic:     KafkaTopic,
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})
	//err = r.SetOffset(42)
	//if err != nil {
	//	log.Printf("err:%v", err)
	//	return
	//}
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		log.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}

	if err := r.Close(); err != nil {
		log.Printf("failed to close reader:%v", err)
	}
	return
}

// GroupConsumer GroupConsumer
func GroupConsumer() (err error) {
	// make a new reader that consumes from topic-A
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{KafkaConfig},
		GroupID:  "consumer-group-id",
		Topic:    KafkaTopic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		log.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
	return
}
