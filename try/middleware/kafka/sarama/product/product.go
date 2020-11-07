package product

import (
	"github.com/Shopify/sarama"
	"log"
	"time"
)

const (
	KafkaConfig = "127.0.0.1:9092"
	KafkaTopic  = "nginx_log"
)

func Product() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	client, err := sarama.NewSyncProducer([]string{KafkaConfig}, config)
	if err != nil {
		log.Println("producer close, err:", err)
		return
	}
	defer client.Close()
	for {
		msg := &sarama.ProducerMessage{}
		msg.Topic = KafkaTopic
		msg.Value = sarama.StringEncoder("this is a good test, my message is good")

		pid, offset, err := client.SendMessage(msg)
		if err != nil {
			log.Println("send message failed,", err)
			return
		}

		log.Printf("pid:%v offset:%v\n", pid, offset)
		time.Sleep(1 * time.Second)
	}
}
