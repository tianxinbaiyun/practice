package config

import (
	"flag"
	"fmt"
	"github.com/tianxinbaiyun/practice/try/go-micro/core/util"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"time"
)

// Cfg Cfg
var Cfg Config

// Config Config
type Config struct {
	Mysql               Mysql            `yaml:"mysql"`
	Debug               bool             `yaml:"debug"`
	Jwt                 Jwt              `yaml:"jwt"`
	RegConsul           bool             `yaml:"reg_consul"`
	Consuls             []string         `yaml:"consuls"`
	RegEtcd             bool             `yaml:"reg_etcd"`
	Etcd                []string         `yaml:"etcd"`
	RrfURL              string           `yaml:"rrf_url"`
	H5URL               string           `yaml:"h5_url"`
	RegisterTTL         time.Duration    `yaml:"register_ttl"`
	RegisterInterval    time.Duration    `yaml:"register_interval"`
	TsURL               string           `yaml:"ts_url"`
	Redis               Redis            `yaml:"redis"`
	FileServerURL       string           `yaml:"file_server_url"`
	FileServerUploadURL string           `yaml:"file_server_upload_url"`
	FileHasMerge        bool             `yaml:"file_has_merge"`
	Sms                 Sms              `yaml:"sms"`
	Wechat              Wechat           `yaml:"wechat"`
	Im                  Im               `yaml:"im"`
	MediaServer         MediaServer      `yaml:"media_server"`
	JPush               JPush            `yaml:"j_push"`
	Chinaums            Chinaums         `yaml:"chinaums"`
	Tencentcloudlive    Tencentcloudlive `yaml:"tencentcloudlive"`
	Rabbitmq            Rabbitmq         `yaml:"rabbitmq"`
	Environment         string           `yaml:"environment"`
	PadDeviceSn         string           `yaml:"pad_device_sn"`
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

// Jwt Jwt
type Jwt struct {
	Signkey    string        `yaml:"signkey"`
	Duration   time.Duration `yaml:"duration"`
	RefreshTTL time.Duration `yaml:"refresh_ttl"`
}

// Redis Redis
type Redis struct {
	RedisAddr        string `yaml:"redis_addr"`
	RedisMaxIDle     int    `yaml:"redis_max_idle"`     //最大等待连接中的数量
	RedisMaxActive   int    `yaml:"redis_max_active"`   //最大连接数据库连接数
	RedisIDleTimeout int    `yaml:"redis_idle_timeout"` //最大等待毫秒数
	RedisPassword    string `yaml:"redis_password"`     //密码
}

//Sms Sms
type Sms struct {
	Default string    `yaml:"default"`
	IsOn    bool      `yaml:"is_on"`
	Aliyun  SmsAliyun `yaml:"aliyun"`
	Juhe    SmsJuhe   `yaml:"juhe"`
}

// SmsAliyun SmsAliyun
type SmsAliyun struct {
	RegionID        string `yaml:"region_id"`
	AccessKeyID     string `yaml:"access_key_id"`
	AccessKeySecret string `yaml:"access_key_secret"`
	SignName        string `yaml:"sign_name"`
	TemplateCode    string `yaml:"template_code"`
}

// SmsJuhe SmsJuhe
type SmsJuhe struct {
	TplID int    `yaml:"tpl_id"`
	Key   string `yaml:"key"`
}

// Wechat Wechat
type Wechat struct {
	Open WechatOpen `yaml:"open"`
}

// WechatOpen WechatOpen
type WechatOpen struct {
	Appid          string `yaml:"appid"`
	AppSecret      string `yaml:"app_secret"`
	Token          string `yaml:"token"`
	EncodingAESKey string `yaml:"encoding_AES_key"`
}

// Im Im
type Im struct {
	URL     string        `yaml:"url"`
	Timeout time.Duration `yaml:timeout`
}

// MediaServer MediaServer
type MediaServer struct {
	HTTPAddr string `yaml:"http_addr"`
	RtspAddr string `yaml:"rtsp_addr"`
	RtspPort string `yaml:"rtsp_port"`
	Ffmpeg   string `yaml:"ffmpeg"`
}

// JPush JPush
type JPush struct {
	ApnsProduction bool   `yaml:"apns_production"`
	Secret         string `yaml:"secret"`
	AppKey         string `yaml:"app_key"`
}

// Chinaums Chinaums
type Chinaums struct {
	PayURL   string `yaml:"pay_url"`
	QueryURL string `yaml:"query_url"`
}

// Tencentcloudlive Tencentcloudlive
type Tencentcloudlive struct {
	SignOn       bool   `yaml:"sign_on"`
	PushDomain   string `yaml:"push_domain"`
	PlayDomain   string `yaml:"play_domain"`
	Key          string `yaml:"key"`
	SecretID     string `yaml:"secret_id"`
	SecretKey    string `yaml:"secret_key"`
	Region       string `yaml:"region"`
	AppName      string `yaml:"app_name"`
	ImSdkAppid   int    `yaml:"im_sdk_appid"`
	ImKey        string `yaml:"im_key"`
	ImIDentifier string `yaml:"im_identifier"`
}

// Rabbitmq Rabbitmq
type Rabbitmq struct {
	URL          string       `yaml:"url"`
	LiveCallback LiveCallback `yaml:"live_callback"`
	ImCallback   ImCallback   `yaml:"im_callback"`
}

// LiveCallback LiveCallback
type LiveCallback struct {
	Exchange string `yaml:"exchange"`
	Queue    string `yaml:"queue"`
}

// ImCallback ImCallback
type ImCallback struct {
	Exchange string `yaml:"exchange"`
	Queue    string `yaml:"queue"`
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
