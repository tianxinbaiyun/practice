package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/satori/go.uuid"
	"log"
	"sync"
	"time"
)

//LogsInnodb 定义结构体(xorm支持双向映射)
type LogsInnodb struct {
	Name      string `json:"name" xorm:"VARCHAR(255)"`
	Time      int64  `json:"time" xorm:"updated BIGINT(20)"`
	CreatTime int64  `json:"creat_time" xorm:"created BIGINT(20)"`
}

// LogsMyisam LogsMyisam
type LogsMyisam struct {
	Name      string `json:"name" xorm:"VARCHAR(255)"`
	Time      int64  `json:"time" xorm:"updated BIGINT(20)"`
	CreatTime int64  `json:"creat_time" xorm:"created BIGINT(20)"`
}

// LogsArchive LogsArchive
type LogsArchive struct {
	Name      string `json:"name" xorm:"VARCHAR(255)"`
	Time      int64  `json:"time" xorm:"updated BIGINT(20)"`
	CreatTime int64  `json:"creat_time" xorm:"created BIGINT(20)"`
}

var x *xorm.Engine

//创建orm引擎
func init() {
	var err error
	x, err = xorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3306)/study?charset=utf8")
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	if err := x.Sync(new(LogsMyisam)); err != nil {
		log.Fatal("数据表同步失败:", err)
	}
}

func main2() {
	now := time.Now().UnixNano()
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			InsertLogsArchiveData(100000)
			wg.Done()
			return
		}()
	}
	wg.Wait()
	fmt.Println("time:", float64(time.Now().UnixNano()-now)/1000000000)
	fmt.Println("--------------------")
}

// InsertLogsInnodbData InsertLogsInnodbData
func InsertLogsInnodbData(n int) {
	for i := 0; i < n; i++ {
		u1 := uuid.Must(uuid.NewV4(), nil)
		name := fmt.Sprintf("%s", u1)
		_, _ = InsertLogsInnodb(name, float64(i))
	}
	return
}

// InsertLogsMyisamData InsertLogsMyisamData
func InsertLogsMyisamData(n int) {
	for i := 0; i < n; i++ {
		u1 := uuid.Must(uuid.NewV4(), nil)
		name := fmt.Sprintf("%s", u1)
		_, _ = InsertLogsMyisam(name, float64(i))
	}
	return
}

// InsertLogsArchiveData InsertLogsArchiveData
func InsertLogsArchiveData(n int) {
	for i := 0; i < n; i++ {
		u1 := uuid.Must(uuid.NewV4(), nil)
		name := fmt.Sprintf("%s", u1)
		_, _ = InsertLogsArchive(name, float64(i))
	}
	return
}

//InsertLogsInnodb 增
func InsertLogsInnodb(name string, balance float64) (int64, bool) {
	data := new(LogsInnodb)
	data.Name = name
	affected, err := x.Insert(data)
	if err != nil {
		return affected, false
	}
	return affected, true
}

// InsertLogsMyisam InsertLogsMyisam
func InsertLogsMyisam(name string, balance float64) (int64, bool) {
	data := new(LogsMyisam)
	data.Name = name
	affected, err := x.Insert(data)
	if err != nil {
		return affected, false
	}
	return affected, true
}

// InsertLogsArchive InsertLogsArchive
func InsertLogsArchive(name string, balance float64) (int64, bool) {
	data := new(LogsArchive)
	data.Name = name
	affected, err := x.Insert(data)
	if err != nil {
		return affected, false
	}
	return affected, true
}
