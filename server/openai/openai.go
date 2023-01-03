package openai

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/felicepng/moodboard/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const AI_MODEL = "text-davinci-003"
const IMAGES_COUNT = 8
const PROMPT_MAX_WORDS = 8

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
		log.Printf("Error occurred generating prompts: %s\n", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error occurred generating prompts",
		})
		return
	}

	// TODO: generate images from prompts

	c.JSON(http.StatusOK, gin.H{
		"prompts": prompts,
	})
}

func GeneratePromptsFromTheme(theme string) (string, error) {
	log.Printf("Theme: %s\n", theme)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file\n")
	}

	data, err := json.Marshal(models.GeneratePromptsReq{
		Model:     AI_MODEL,
		Prompt:    fmt.Sprintf("Write %d image prompts, each having maximum %d words, for a moodboard with the theme: %s. The prompts should not be numbered, instead separated by one '|' character", IMAGES_COUNT, PROMPT_MAX_WORDS, theme),
		MaxTokens: IMAGES_COUNT*PROMPT_MAX_WORDS*3 + 40,
	})
	if err != nil {
		log.Printf("Error marshalling data: %v\n", err)
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, "https://api.openai.com/v1/completions", bytes.NewBuffer(data))
	if err != nil {
		log.Printf("Error creating request: %v\n", err)
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+os.Getenv("API_KEY"))

	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending request: %v\n", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response bytes: %v\n", err)
		return "", err
	}

	var v models.GeneratePromptsResp
	if err := json.Unmarshal(body, &v); err != nil || v.Choices == nil {
		log.Println("Invalid response")
		return "", errors.New("invalid response")
	}

	return v.Choices[0].Text, nil
}
