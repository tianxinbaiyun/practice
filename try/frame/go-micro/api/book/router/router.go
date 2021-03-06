package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tianxinbaiyun/practice/try/frame/go-micro/core/config"
)

var (
	r        *gin.Engine
	basePath string
)

// Init Init
func Init(serviceName string) *gin.Engine {
	basePath = serviceName
	r = gin.New()
	if config.Cfg.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	apiRouter(r)

	return r
}
