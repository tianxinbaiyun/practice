package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tianxinbaiyun/practice/try/frame/go-micro/api/book/handler/api"
)

func apiRouter(r *gin.Engine) {

	book := new(api.Book)

	router := r.Group(basePath)
	{
		router.GET("", book.List)     //文章列表
		router.GET("/:id", book.Info) //文章详情
	}
}
