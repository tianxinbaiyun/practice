package main

import (
	"encoding/json"
	"fmt"
	"github.com/influxdata/influxdb/client/v2"
	"log"
	"time"
)

// 定义
const (
	MyDB          = "github.com/tianxinbaiyun/practice"
	username      = "admin"
	password      = "admin"
	MyMeasurement = "cpu_usage"
)

func main() {
	conn := connInflux()
	fmt.Println(conn)

	//insert
	WritesPoints(conn)

	//获取10条数据并展示
	qs := fmt.Sprintf("SELECT * FROM %s LIMIT %d", MyMeasurement, 10)
	res, err := QueryDB(conn, qs)
	if err != nil {
		log.Fatal(err)
	}

	for i, row := range res[0].Series[0].Values {
		t, err := time.Parse(time.RFC3339, row[0].(string))
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(reflect.TypeOf(row[1]))
		valu := row[2].(json.Number)
		log.Printf("[%2d] %s: %s\n", i, t.Format(time.Stamp), valu)
	}
}

func connInflux() client.Client {
	cli, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://192.168.15.131:18086",
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Fatal(err)
	}
	return cli
}

//QueryDB QueryDB
func QueryDB(cli client.Client, cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: MyDB,
	}
	if response, err := cli.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	} else {
		return res, err
	}
	return res, nil
}

//WritesPoints WritesPoints
func WritesPoints(cli client.Client) {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  MyDB,
		Precision: "s",
	})
	if err != nil {
		log.Fatal(err)
	}

	tags := map[string]string{"cpu": "ih-cpu"}
	fields := map[string]interface{}{
		"idle":   20.1,
		"system": 43.3,
		"user":   86.6,
	}

	pt, err := client.NewPoint(
		"cpu_usage",
		tags,
		fields,
		time.Now(),
	)
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)

	if err := cli.Write(bp); err != nil {
		log.Fatal(err)
	}
}
