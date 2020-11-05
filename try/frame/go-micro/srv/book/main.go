package main

import (
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	//"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/server"
	"github.com/tianxinbaiyun/practice/try/frame/go-micro/core/config"
	"github.com/tianxinbaiyun/practice/try/frame/go-micro/core/db"
	pb_book "github.com/tianxinbaiyun/practice/try/frame/go-micro/core/pb/book"
	"github.com/tianxinbaiyun/practice/try/frame/go-micro/core/slog"
	"github.com/tianxinbaiyun/practice/try/frame/go-micro/srv/book/handler"
	"log"
	"time"
)

// ServiceName
var (
	ServiceName = "book"
)

func main() {
	options := []micro.Option{
		micro.Name("shanxing.srv." + ServiceName),
		micro.Version("latest"),
		micro.Flags(cli.StringFlag{Name: "c"}),
		micro.RegisterTTL(config.Cfg.RegisterTTL * time.Second),
		micro.RegisterInterval(config.Cfg.RegisterInterval * time.Second),
	}
	//if config.Cfg.RegConsul {
	//	reg := consul.NewRegistry(func(op *registry.Options) {
	//		op.Addrs = config.Cfg.Consuls
	//	})
	//	options = append(options, micro.Registry(reg))
	//}

	service := micro.NewService(options...)

	_, err := slog.NewLog("Info", true, 10)
	if err != nil {
		log.Fatal(err)
	}
	service.Init(
		micro.Action(func(c *cli.Context) {
			slog.Info(ServiceName + " srv server is start ...")
			db.Init()
		}),
		micro.AfterStop(func() error {
			slog.Info(ServiceName + " srv server is stop ...")
			return nil
		}),
	)

	pb_book.RegisterBookServiceHandler(service.Server(), handler.NewBookService(), server.InternalHandler(true))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
