package db

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/go-xorm/xorm"
	"github.com/tianxinbaiyun/practice/try/frame/go-micro/core/config"
	"time"
)

var (
	err       error
	defDB     *xorm.EngineGroup
	redisPool *redis.Pool
)

// Init Init
func Init() {

	connGroup := make([]string, 0)

	master := config.Cfg.Mysql.DefMaster
	maDB := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&loc=Local",
		master.User, master.Password, master.Host, master.Port, master.Database, master.Charset,
	)

	connGroup = append(connGroup, maDB)

	for _, val := range config.Cfg.Mysql.DefSlaves {
		connGroup = append(connGroup, fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=%s&loc=Local",
			val.User, val.Password, val.Host, val.Port, val.Database, val.Charset,
		))
	}

	defDB, err = xorm.NewEngineGroup("mysql", connGroup)
	if err != nil {
		panic(err)
	}
	defDB.SetMaxIdleConns(config.Cfg.Mysql.MaxIDleConns)
	defDB.SetMaxOpenConns(config.Cfg.Mysql.MaxOpenConns)
	if config.Cfg.Debug {
		defDB.ShowSQL(true)
	}

}

// Get 获取数据库示例
func Get() *xorm.EngineGroup {
	return defDB
}

// SessionBegin 开始事物
func SessionBegin(se *xorm.Session) (err error) {
	err = se.Begin()
	return
}

// SessionRollback 回滚事物
func SessionRollback(se *xorm.Session) (err error) {
	err = se.Rollback()
	return
}

// SessionCommit 提交事物
func SessionCommit(se *xorm.Session) (err error) {
	err = se.Commit()
	return
}

//InitRedisPool 初始化redis
func InitRedisPool() {
	redisConf := config.Cfg.Redis
	redisPool = &redis.Pool{
		MaxIdle:     redisConf.RedisMaxIDle,
		MaxActive:   redisConf.RedisMaxActive,
		IdleTimeout: time.Duration(redisConf.RedisIDleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisConf.RedisAddr)
			if err != nil {
				return nil, err
			}
			if redisConf.RedisPassword != "" {
				if _, err := c.Do("AUTH", redisConf.RedisPassword); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
	}
}

// GetRedis 获取redis
func GetRedis() *redis.Pool {
	return redisPool
}
