package comsumer

import (
	"github.com/Shopify/sarama"
	"log"
	"strings"
	"sync"
)

const (
	KafkaConfig = "127.0.0.1:9092"
	KafkaTopic  = "nginx_log"
)

func Consumer() {
	consumer, err := sarama.NewConsumer(strings.Split(KafkaConfig, ","), nil)
	if err != nil {
		log.Printf("Failed to start,err:%v", err)
	}
	partitionList, err := consumer.Partitions(KafkaTopic)
	if err != nil {
		log.Printf("Failed to get the list of partitions:%v ", err)
		return
	}
	log.Printf("partitionList:%v", partitionList)
	var wg sync.WaitGroup
	for partition := range partitionList {
		pc, err := consumer.ConsumePartition("nginx_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			log.Printf("Failed to start consumer for partition %d: %s\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		go func(pc sarama.PartitionConsumer) {
			wg.Add(1)
			for msg := range pc.Messages() {
				log.Printf("Partition:%d, Offset:%d, Key:%s, Value:%s",
					msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
				log.Println()
			}
			wg.Done()
		}(pc)
	}
	wg.Wait()
	select {}
	consumer.Close()
	log.Println("Consumer stop")
}
