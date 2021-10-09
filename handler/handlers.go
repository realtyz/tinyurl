package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"url-shortener/config"
	"url-shortener/shortener"
	"url-shortener/store"
)

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId string `json:"user_id" binding:"required"`
}

func HandleIndexPage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hey Go URL Shortener 🚀!",
	})
}

func CreateShortUrl(c *gin.Context) {
	var creationRequest UrlCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := shortener.GenerateShortLink(creationRequest.LongUrl, creationRequest.UserId)
	store.SaveUrlMapping(shortUrl, creationRequest.LongUrl, creationRequest.UserId)

	// 获取域名
	host := config.GetGlobalConfig().Domain
	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortUrl,
	})
}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	initialUrl := store.RetrieveInitialUrl(shortUrl)
	c.Redirect(302, initialUrl)
}
