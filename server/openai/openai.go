package openai

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/felicepng/moodboard/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func GenerateImages(c *gin.Context) {
	var moodboard models.MoodboardJson
	if err := json.NewDecoder(c.Request.Body).Decode(&moodboard); err != nil {
		log.Printf("Error occurred unmarshalling json: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error occurred unmarshalling json",
		})
		return
	}

	prompts, err := GeneratePromptsFromTheme(moodboard.Theme)
	if err != nil {
		log.Printf("Error occurred generating prompts: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error occurred generating prompts",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": prompts,
	})
}

func GeneratePromptsFromTheme(theme *string) (string, error) {
	if theme != nil {
		log.Printf("Theme: %s", *theme)
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("API_KEY")
	log.Printf("API Key: %s", apiKey)

	// data := map[string]interface{}{
	// 	"model":  "text-davinci-003",
	// 	"prompt": "Write 8 individual image prompts, numbered from 1 to 8, for a moodboard with the following theme: " + theme,
	// 	"n":      1,
	// }

	// jsonData, err := json.Marshal(data)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// req, err := http.NewRequest(http.MethodPost, "https://api.openai.com/v1/completions", bytes.NewBuffer(jsonData))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// req.Header.Add("Authorization", "Bearer "+apiKey)

	// client := &http.Client{
	// 	Timeout: time.Second * 10,
	// }
	// resp, err := client.Do(req)
	// if err != nil {
	// 	log.Println("Error:", err.Error())
	// 	return "", err
	// }
	// defer resp.Body.Close()

	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Println("Error while reading the response bytes:", err)
	// }
	// return string([]byte(body)), nil

	return "", nil
}
