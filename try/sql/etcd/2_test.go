package etcd

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"log"
	"testing"
	"time"
)

type EtcdMutex struct {
	s *concurrency.Session
	m *concurrency.Mutex
}

func NewMutex(key string, client *clientv3.Client) (mutex *EtcdMutex, err error) {
	mutex = &EtcdMutex{}
	mutex.s, err = concurrency.NewSession(client)
	if err != nil {
		return
	}
	mutex.m = concurrency.NewMutex(mutex.s, key)
	return
}

func (mutex *EtcdMutex) Lock() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second) //设置20s超时
	defer cancel()
	if err = mutex.m.Lock(ctx); err != nil {
		return err
	}
	return
}

func (mutex *EtcdMutex) Unlock() (err error) {
	err = mutex.m.Unlock(context.TODO())
	if err != nil {
		return
	}
	err = mutex.s.Close()
	if err != nil {
		return
	}
	return
}

func Test2(t *testing.T) {
	client, err := Init(SetEndpoint("127.0.0.1:12379"))
	if err != nil {
		return
	}
	locker1, err := NewMutex("123", client)
	if err != nil {
		return
	}
	locker2, err := NewMutex("123", client)
	if err != nil {
		return
	}
	//groutine1
	go func() {
		err := locker1.Lock()
		if err != nil {
			log.Println("groutine1抢锁失败")
			log.Println(err)
			return
		}
		log.Println("groutine1抢锁成功")
		time.Sleep(10 * time.Second)
		defer locker1.Unlock()

	}()

	//groutine2
	go func() {
		err := locker2.Lock()
		if err != nil {
			log.Println("groutine2抢锁失败")
			log.Println(err)
			return
		}
		log.Println("groutine2抢锁成功")
		defer locker2.Unlock()
	}()
	time.Sleep(30 * time.Second)
}
