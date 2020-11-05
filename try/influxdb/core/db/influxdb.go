package db

import (
	"fmt"
	"github.com/influxdata/influxdb/client/v2"
	"github.com/tianxinbaiyun/practice/try/influxdb/core/config"
	"log"
)

var (
	err   error
	defDb *mainDB
)

type mainDB struct {
	Conn client.Client
	Bp   client.BatchPoints
}

// Init Init
func Init() {
	defDb = new(mainDB)
	defDb.Conn, err = client.NewHTTPClient(client.HTTPConfig{
		Addr:     fmt.Sprintf("%s:%s", config.Cfg.Influxdb.Host, config.Cfg.Influxdb.Port),
		Username: config.Cfg.Influxdb.User,
		Password: config.Cfg.Influxdb.Password,
	})
	return
}

// NewPersist NewPersist
func NewPersist(conf config.Config) (conn client.Client, err error) {
	conn, err = client.NewHTTPClient(client.HTTPConfig{
		Addr:     fmt.Sprintf("%s:%s", config.Cfg.Influxdb.Host, config.Cfg.Influxdb.Port),
		Username: config.Cfg.Influxdb.User,
		Password: config.Cfg.Influxdb.Password,
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	return
}

//Get 获取数据库示例
func Get() *mainDB {
	defDb.Bp, err = client.NewBatchPoints(client.BatchPointsConfig{
		Database:  config.Cfg.Influxdb.Database,
		Precision: "s",
	})

	if err != nil {
		log.Fatal(err)
	}
	return defDb
}
