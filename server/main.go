package main

import (
	"net/http"

	"github.com/felicepng/moodboard/openai"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func Healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Service is healthy",
	})
}

func main() {
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./web", true)))
	api := r.Group("/api")

	api.GET("/health", Healthcheck)
	api.POST("/images", openai.GenerateImages)

	r.Run(":8080")
}
