package config

import (
	"flag"
	"fmt"
	"github.com/tianxinbaiyun/practice/try/mysql/core/util"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

// Cfg Cfg
var Cfg Config

// Config Config
type Config struct {
	Mysql Mysql `yaml:"mysql"`
	Debug bool  `yaml:"debug"`
}

// Mysql Mysql
type Mysql struct {
	DefMaster    MysqlBase   `yaml:"def_master"`
	DefSlaves    []MysqlBase `yaml:"def_slaves"`
	TsMall       MysqlBase   `yaml:"ts_mall"`
	CoreMaster   MysqlBase   `yaml:"core_master"`
	CoreSlaves   []MysqlBase `yaml:"core_slaves"`
	BuddhaMaster MysqlBase   `yaml:"buddha_master"`
	BuddhaSlaves []MysqlBase `yaml:"buddha_slaves"`
	MaxIDleConns int         `yaml:"max_idle_conns"`
	MaxOpenConns int         `yaml:"max_open_conns"`
}

// MysqlBase MysqlBase
type MysqlBase struct {
	Host     string `yaml:"host"`
	Database string `yaml:"database"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
	Charset  string `yaml:"charset"`
}

func init() {
	fmt.Println(util.GetExecpath())
	fmt.Println(util.GetCurrentPath())
	versiontype := flag.String("c", "local", "配置文件，可选项:local(本地环境),dev(测试环境),release(正式环境)")
	flag.Parse()
	//配置文件名字
	confFile := *versiontype + ".yaml"
	//当前程序运行的目录，获取配置文件
	filePath := util.GetExecpath() + "/../conf/" + confFile

	//配置文件不存在，从配置文件指定的目录找
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		filePath = util.GetCurrentPath() + "/../conf/" + confFile
	}
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("读取配置文件错误")
	}

	//读取的数据为yaml格式，需要进行解码
	err = yaml.Unmarshal(data, &Cfg)
	if err != nil {
		log.Printf("%v\n", err)
		log.Fatal("解析配置文件错误")
	}
}
