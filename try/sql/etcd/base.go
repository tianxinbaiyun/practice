package etcd

import (
	"context"
	"log"
	"sync"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

var (
	config clientv3.Config
	cli    *clientv3.Client
	err    error
	once   sync.Once
)

type clientOption struct {
	Endpoints   []string
	DialTimeout time.Duration
}

// ClientOption ClientOption
type ClientOption interface {
	apply(*clientOption)
}

type endpointOption struct {
	Endpoints []string
}

func (eo *endpointOption) apply(co *clientOption) {
	co.Endpoints = eo.Endpoints
}

// SetEndpoint set ip:port of etcd
func SetEndpoint(endpoint string) ClientOption {
	return &endpointOption{
		Endpoints: []string{endpoint},
	}
}

// Init init etcd client
func Init(opts ...ClientOption) (*clientv3.Client, error) {
	once.Do(func() {
		co := &clientOption{
			Endpoints:   []string{"0.0.0.0:2379"},
			DialTimeout: 5 * time.Second,
		}

		for _, opt := range opts {
			opt.apply(co)
		}

		config = clientv3.Config{
			Endpoints:   co.Endpoints,
			DialTimeout: co.DialTimeout,
		}

		log.Printf("endpoints: %v", config.Endpoints)
		cli, err = clientv3.New(config)
		if err != nil {
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_, err = clientv3.NewKV(cli).Get(ctx, "a")
	})
	return cli, err
}
