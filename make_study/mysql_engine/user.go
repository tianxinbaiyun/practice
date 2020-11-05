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

//User 定义结构体(xorm支持双向映射)
type User struct {
	UserID    int64  `xorm:"pk autoincr"` //指定主键并自增
	Name      string `xorm:"unique"`      //唯一的
	Balance   float64
	Time      int64 `xorm:"updated"` //修改后自动更新时间
	CreatTime int64 `xorm:"created"` //创建时间
}

//  News News
type News struct {
	ID      uint32 `xorm:"pk autoincr"` //指定主键并自增
	Content string `xorm:"content"`     //内容
}

//创建orm引擎
func init() {
	var err error
	x, err = xorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3306)/study?charset=utf8")
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	if err := x.Sync(new(User)); err != nil {
		log.Fatal("数据表同步失败:", err)
	}
}

func main() {
	now := time.Now().UnixNano()

	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			InsertUserData(10)
			wg.Done()
			return
		}()
	}
	wg.Wait()
	fmt.Println("all:", time.Now().UnixNano()-now)
	fmt.Println("--------------------")
	fmt.Println("end")
}

// InsertUserData InsertUserData
func InsertUserData(n int) {
	for i := 0; i < n; i++ {
		u1 := uuid.Must(uuid.NewV4(), nil)
		name := fmt.Sprintf("%s", u1)
		_, _ = InsertUser(name, float64(i))
	}
	return
}

// InsertUser 增
func InsertUser(name string, balance float64) (int64, bool) {
	user := new(User)
	user.Name = name
	user.Balance = balance
	affected, err := x.Insert(user)
	if err != nil {
		return affected, false
	}
	return affected, true
}

//DelUser 删
func DelUser(id int64) {
	user := new(User)
	x.ID(id).Delete(user)
}

// UpdateUser 改
func UpdateUser(id int64, user *User) bool {
	affected, err := x.ID(id).Update(user)
	if err != nil {
		log.Fatal("错误:", err)
	}
	if affected == 0 {
		return false
	}
	return true
}

//GetUserinfo 查
func GetUserinfo(id int64) *User {
	user := &User{UserID: id}
	is, _ := x.Get(user)
	if !is {
		log.Fatal("搜索结果不存在!")
	}
	return user
}
