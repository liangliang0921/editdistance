package main

import (
	"./api"
	"github.com/gin-gonic/gin"
	// "net/http"
)

// 根据请求处理函数，所有请求相关方都在context中，输出响应hello word
// func WebRoot(context *gin.Context) {
//	context.String(http.StatusOK, "hello world")
//}

func main() {

	// 初始化引擎
	router := gin.Default()

	// 注册一个路由和处理函数
	v1 := router.Group("/")
	{
		v1.POST("/editdistance", api.GetShortestEditDistance)
	}
	// engine.Any("/", WebRoot)

	// 绑定端口号，然后启动应用
	router.Run(":9200")

}
