package db

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/tianxinbaiyun/practice/try/tidb/core/config"
)

var (
	err   error
	defDB *xorm.EngineGroup
)

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
	defDB.SetMaxIDleConns(config.Cfg.Mysql.MaxIDleConns)
	defDB.SetMaxOpenConns(config.Cfg.Mysql.MaxOpenConns)
	if config.Cfg.Debug {
		defDB.ShowSQL(true)
	}

}

//获取rrf_plus数据库示例
func Get() *xorm.EngineGroup {
	return defDB
}

// 开始事物
func SessionBegin(se *xorm.Session) (err error) {
	err = se.Begin()
	return
}

// 回滚事物
func SessionRollback(se *xorm.Session) (err error) {
	err = se.Rollback()
	return
}

// 提交事物
func SessionCommit(se *xorm.Session) (err error) {
	err = se.Commit()
	return
}
