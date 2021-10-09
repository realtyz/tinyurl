package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"url-shortener/config"
	"url-shortener/route"
	"url-shortener/store"
)

func main() {
	// 发行模式
	gin.SetMode(gin.ReleaseMode)
	// 获取路由
	r := gin.Default()
	// 初始化路由
	route.InitializeRoutes(r)
	// 初始化存储服务
	store.InitializeStorage()
	// 监听端口
	err := r.Run(":" + config.GetGlobalConfig().Port)

	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}