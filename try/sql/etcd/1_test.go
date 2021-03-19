package etcd

import (
	"context"
	"log"
	"testing"
	"time"

	"go.etcd.io/etcd/client/v3/concurrency"
)

func Test1(t *testing.T) {

	key := "asdfg"

	cli, err := Init(SetEndpoint("127.0.0.1:12379"))
	if err != nil {
		log.Fatal(err)
	}

	s1, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatal(err)
	}

	s2, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatal(err)
	}

	mutex := concurrency.NewMutex(s1, key)
	mutex2 := concurrency.NewMutex(s2, key)

	go func() {
		if err := mutex.Lock(context.TODO()); err != nil {
			log.Fatal(err)
		}
		log.Println("aaaa")
		time.Sleep(10 * time.Second)
		mutex.Unlock(context.TODO())
	}()

	go func() {
		if err := mutex2.Lock(context.TODO()); err != nil {
			log.Fatal(err)
		}

		log.Println("bbbbb")
		mutex2.Unlock(context.TODO())
	}()

	select {}

}
