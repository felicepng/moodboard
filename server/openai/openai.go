package openai

import "github.com/gin-gonic/gin"

func GenerateImages(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Images generated",
	})
}
