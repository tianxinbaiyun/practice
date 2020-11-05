package main

import (
	"github.com/micro/cli"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-micro/web"
	"github.com/tianxinbaiyun/practice/try/go-micro/api/book/router"
	"github.com/tianxinbaiyun/practice/try/go-micro/core/config"
	"github.com/tianxinbaiyun/practice/try/go-micro/core/slog"
	"log"
	"time"
)

// 服务名称
var (
	ServiceName = "book"
)

//micro api --handler http
func main() {
	options := []web.Option{
		web.Name("shanxing.api." + ServiceName),
		web.Version("latest"),
		web.Flags(cli.StringFlag{Name: "c"}),
		web.RegisterTTL(config.Cfg.RegisterTTL * time.Second),
		web.RegisterInterval(config.Cfg.RegisterInterval * time.Second),
	}
	if config.Cfg.RegEtcd {
		//reg := consul.NewRegistry(func(op *registry.Options) {
		//	op.Addrs = config.Cfg.Consuls
		//})
		reg := etcd.NewRegistry(func(op *registry.Options) {
			op.Addrs = config.Cfg.Etcd
		})
		options = append(options, web.Registry(reg))
	}

	service := web.NewService(options...)

	_, err := slog.NewLog("Info", true, 10)
	if err != nil {
		log.Fatal(err)
	}
	err = service.Init(
		web.Action(func(c *cli.Context) {
			//db.Init()
			slog.Info(ServiceName + " api is start ...")
		}),
		web.AfterStop(func() error {
			slog.Info(ServiceName + " api is stop ...")
			slog.Close()
			return nil
		}),
	)

	if err != nil {
		log.Fatal(err)
	}

	service.Handle("/", router.Init(ServiceName))
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
