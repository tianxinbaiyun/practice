package config

// 定义
const (
	RMQADDR      = "amqp://guest:guest@192.168.15.131:5672/"
	EXCHANGENAME = "syslog_direct"
	CONSUMERCNT  = 4
)

// 定义
var (
	RoutingKeys [4]string = [4]string{"info", "debug", "warn", "error"}
)
