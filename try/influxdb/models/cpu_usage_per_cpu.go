package models

import (
	"github.com/influxdata/influxdb/client/v2"
	"github.com/tianxinbaiyun/practice/try/influxdb/core/db"
	"log"
	"time"
)

// CPUUsagePerCPU CPUUsagePerCPU
type CPUUsagePerCPU struct {
	ID            int64     `json:"id"`
	Time          time.Time `json:"time"`
	ContainerName string    `json:"container_name"`
	Instance      string    `json:"instance"`
	Machine       string    `json:"machine"`
	Value         string    `json:"value"`
}

// Create Create
func Create(par *CPUUsagePerCPU) (err error) {
	mainDb := db.Get()
	tags := map[string]string{"container": "mysql"}
	fields := map[string]interface{}{
		"id": 1,
	}

	pt, err := client.NewPoint("cpu_usage_per_cpu", tags, fields, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	mainDb.Bp.AddPoint(pt)

	if err := mainDb.Conn.Write(mainDb.Bp); err != nil {
		log.Fatal(err)
	}
	return
}
