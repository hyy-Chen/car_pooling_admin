package main

import (
	"car_pooling_admini/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建Gin路由引擎
	r := gin.Default()

	// 配置路由
	routers.AdminRoutersInit(r)

	// 启动Gin服务
	r.Run(":8080")
}
