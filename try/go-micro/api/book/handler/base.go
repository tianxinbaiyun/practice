package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tianxinbaiyun/practice/try/go-micro/core/config"
	"net/http"
)

// 业务逻辑编码
const (
	Success         = 0    // 请求成功
	PublicError     = 400  //一般错误
	InvalidCode     = 9001 // 无效的业务编码
	InvalidParam    = 9002 // 无效的参数
	InvalidUserInfo = 9100 // 无效的用户信息
	InvalidToken    = 9101 // 无效的token
	Unauthorized    = 9102 // 未授权
	NoTemple        = 9103 // 用户不属于任何寺院
	InOtherDevices  = 9104 // 该账户已在其他设备登陆
	SystemError     = 9999 // 业务系统错误
)

// AppCodeMsg 业务逻辑编码对应的消息
var AppCodeMsg = map[int]string{
	Success:         "操作成功",
	InvalidCode:     "无效的业务编码",
	InvalidParam:    "无效的参数",
	InvalidUserInfo: "无效的用户信息",
	InvalidToken:    "无效的token",
	Unauthorized:    "请求未授权",
	NoTemple:        "用户不属于任何寺院",
	SystemError:     "业务系统错误",
	InOtherDevices:  "该账户已在其他设备登陆",
	PublicError:     "一般性错误",
}

// 定义
const (
	CurUID      = "cur_uid"
	CurTempleID = "cur_temple_uid"
	CurSupplyID = "cur_supply_id"
	BizUser     = "biz_user"
	TempleUser  = "temple_user"
)

// PlatformType PlatformType
type PlatformType int

//各后台常量标识
const (
	SxAdmin PlatformType = iota
	TsAdmin
	SpAdmin
)

// Base Base
type Base struct {
}

// Success Success
func (b *Base) Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    Success,
		"message": AppCodeMsg[Success],
		"data":    data,
	})
}

// Failure Failure
func (b *Base) Failure(c *gin.Context, httpCode int, code int, message string, err error) {
	if message == "" {
		message = AppCodeMsg[code]
	}
	msgMap := gin.H{
		"code":    code,
		"message": message,
	}
	if config.Cfg.Debug && err != nil {
		msgMap["bebug"] = error.Error(err)
	}
	c.JSON(httpCode, msgMap)
}
