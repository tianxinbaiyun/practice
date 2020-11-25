package product

import (
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/assert"
	"sync/atomic"
	"testing"
	"time"
)

func TestSimpleProduct(t *testing.T) {
	var number int32
	for {
		//for i := 0; i < 10; i++ {
		//	go func() {
		message := fmt.Sprintf("my love %d", number)
		time.Sleep(time.Second)
		err := SimpleProduct(message)
		assert.NoError(t, err)
		atomic.AddInt32(&number, 1)
		//	}()
		//}
	}
}

func TestCreateTopic(t *testing.T) {
	err := CreateTopic(kafka.TopicConfig{
		Topic:             "clx_12",
		NumPartitions:     3,
		ReplicationFactor: 1,
	})
	assert.NoError(t, err)
}

func TestGetTopicList(t *testing.T) {
	list, err := GetTopicList()
	assert.NoError(t, err)
	t.Log(list)
}
