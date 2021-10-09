package route

import (
	"github.com/gin-gonic/gin"
	"url-shortener/handler"
)

// InitializeRoutes 初始化路由/*
func InitializeRoutes(router *gin.Engine)  {
	router.GET("/", handler.HandleIndexPage)
	router.POST("/create-short-url", handler.CreateShortUrl)
	router.GET("/:shortUrl", handler.HandleShortUrlRedirect)
}