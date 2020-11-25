package middlewares

import (
	"github.com/tianxinbaiyun/practice/try/frame/go-micro/api/book/handler"
	"github.com/tianxinbaiyun/practice/try/frame/go-micro/core/config"

	"github.com/gin-gonic/gin"
	"strings"
)

// Auth 用户验证中间件
func Auth(c *gin.Context) {
	//
	//bearer := c.GetHeader("Authorization")
	//if bearer == "" {
	//	c.AbortWithStatusJSON(
	//		http.StatusUnauthorized,
	//		errMsg(handler.InvalidToken, fmt.Errorf("token值为空")),
	//	)
	//	return
	//}
	//arr := strings.Split(bearer, " ")
	//if len(arr) < 2 {
	//	c.AbortWithStatusJSON(
	//		http.StatusUnauthorized,
	//		errMsg(handler.InvalidToken, fmt.Errorf("token格式错误")),
	//	)
	//	return
	//}
	////rpc调用
	//response, err := rpc.AuthSrv.Auth(context.TODO(), &pb.AuthReq{
	//	Token: arr[1],
	//})
	//if err != nil {
	//	c.AbortWithStatusJSON(
	//		http.StatusUnauthorized,
	//		errMsg(handler.InvalidToken, err),
	//	)
	//	return
	//}
	//
	//if response.Code != 0 {
	//	if response.Code == 11 {
	//		c.AbortWithStatusJSON(
	//			http.StatusUnauthorized,
	//			errMsg(handler.InOtherDevices, fmt.Errorf("该账户已在其他设备登陆")),
	//		)
	//		return
	//	}
	//	c.AbortWithStatusJSON(
	//		http.StatusUnauthorized,
	//		errMsg(handler.InvalidToken, fmt.Errorf("token格式错误")),
	//	)
	//	return
	//}
	//c.Set(handler.CUR_UID, response.Uid)
	//
	//c.Next()
}

func errMsg(code int, err error) (msg gin.H) {
	message := handler.AppCodeMsg[code]
	if strings.Index(err.Error(), "Error") == -1 {
		message = err.Error()
	}
	msg = gin.H{
		"code":    code,
		"message": message,
	}

	if config.Cfg.Debug {
		msg["debug"] = error.Error(err)
	}
	return
}
