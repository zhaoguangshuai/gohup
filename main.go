package main

import (
	"fmt"
	"gohub/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化 Gin 实例
	router := gin.New()

	// 初始化路由绑定
	bootstrap.SetupRoute(router)

	// 运行服务，默认为 8080，我们指定端口为 8000
	err := router.Run(":3000")
	if err != nil {
		// 错误处理，端口被占用了或者其他错误
		fmt.Println(err.Error())
	}
}
