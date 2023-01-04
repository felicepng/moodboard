package main

import (
	"github.com/felicepng/moodboard/openai"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./web", true)))
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}

func main() {
	r := setupRouter()

	api := r.Group("/api")
	api.POST("/images", openai.GenerateImageUrls)

	r.Run(":8080")
}
