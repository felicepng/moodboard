package openai

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/felicepng/moodboard/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func GenerateImages(c *gin.Context) {
	var moodboard models.MoodboardJson
	if err := c.BindJSON(&moodboard); err != nil {
		fmt.Println("Error occurred unmarshalling moodboard theme:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error occurred unmarshalling moodboard theme",
		})
		return
	}

	GeneratePromptsFromTheme(moodboard.Theme)

	c.JSON(http.StatusOK, gin.H{
		"message": "Images generated",
	})
}

func GeneratePromptsFromTheme(theme string) string {
	fmt.Println("Theme:", theme)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("API_KEY")
	fmt.Println("API Key:", apiKey)

	return ""
}