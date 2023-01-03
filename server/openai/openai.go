package openai

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/felicepng/moodboard/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func GenerateImages(c *gin.Context) {
	var moodboard models.MoodboardJson
	if err := c.BindJSON(&moodboard); err != nil {
		log.Println("Error occurred unmarshalling moodboard theme:", err)
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
	log.Println("Theme:", theme)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("API_KEY")
	log.Println("API Key:", apiKey)

	data := map[string]interface{}{
		"model":  "text-davinci-003",
		"prompt": "Write 8 individual image prompts, numbered from 1 to 8, for a moodboard with the following theme: " + theme,
		"n":      1,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPost, "https://api.openai.com/v1/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Authorization", "Bearer "+apiKey)

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error:", err.Error())
		return ""
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}
	return string([]byte(body))
}
