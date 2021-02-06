package rate

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"testing"
	"time"
)

var configDB *gorm.DB

func init() {
	var (
		err                error
		configDBConfig     string
		defaultMysqlConfig = "root:123456@tcp(localhost:3306)/%s?&parseTime=True&loc=Local"
	)
	mysqlConfig := os.Getenv("MYSQL_CONFIG")

	if mysqlConfig == "" {
		mysqlConfig = defaultMysqlConfig
	}

	configDBConfig = fmt.Sprintf(mysqlConfig, "spa_config")

	if configDB, err = gorm.Open(mysql.Open(configDBConfig), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)}); err != nil {
		panic(err)
	}
	//configDB.AutoMigrate(&RateLimit{})
}

func TestCreateRateLimit(t *testing.T) {
	data := &RateLimit{
		Key:        "123",
		Value:      123,
		ExpireTime: time.Now(),
		Rate:       1,
		Burst:      1,
		Period:     1,
	}
	rows := CreateRateLimit(configDB, data)
	t.Log(rows)
}

func TestSetRateLimit(t *testing.T) {
	data := &RateLimit{
		Key:        "123",
		Value:      123,
		ExpireTime: time.Now(),
		Rate:       1,
		Burst:      2,
		Period:     1,
	}
	rows := SetRateLimit(configDB, data, 123)
	t.Log(rows)
}

func TestGetRateLimit(t *testing.T) {
	data := GetRateLimit(configDB, "123")
	t.Log(data)
}

func TestDeleteRateLimit(t *testing.T) {
	err := DeleteRateLimit(configDB, "123")
	t.Log(err)
}

func TestAllow(t *testing.T) {
	limiter := NewLimiter(configDB)
	for i := 0; i < 10; i++ {
		go func() {
			res, err := limiter.Allow("project:123", PerMinute(2))
			if err != nil {
				panic(err)
			}
			for res.Allowed == 0 {
				time.Sleep(time.Millisecond * 500)
				//t.Logf("res:%+v", *res)
				res, err = limiter.Allow("project:123", PerMinute(2))
				if err != nil {
					panic(err)
				}
			}
			log.Printf("i:%d,res:%+v", i, *res)
		}()

	}
	for i := 0; i < 10; i++ {

		res, err := limiter.Allow("project:123", PerMinute(2))
		if err != nil {
			panic(err)
		}
		for res.Allowed == 0 {
			time.Sleep(time.Millisecond * 500)
			//t.Logf("res:%+v", *res)
			res, err = limiter.Allow("project:123", PerMinute(2))
			if err != nil {
				panic(err)
			}
		}
		log.Printf("i:%d,res:%+v", i, *res)

	}
	select {}
}
