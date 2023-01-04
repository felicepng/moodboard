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
	"reflect"
	"strings"
	"time"

	"github.com/felicepng/moodboard/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const AI_MODEL = "text-davinci-003"
const IMAGES_COUNT = 8
const PROMPT_MAX_WORDS = 8
const IMAGES_SIZE = "256x256"

var urlChan = make(chan string)
var errorChan = make(chan error)

func GenerateImageUrls(c *gin.Context) {
	var moodboard models.MoodboardJson
	if err := json.NewDecoder(c.Request.Body).Decode(&moodboard); err != nil || moodboard == (models.MoodboardJson{}) {
		log.Println("Error occurred unmarshalling json")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error occurred unmarshalling json",
		})
		return
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file\n")
	}
	var API_KEY = os.Getenv("API_KEY")

	prompts, err := GeneratePromptsFromTheme(moodboard.Theme, API_KEY)
	if err != nil {
		log.Printf("Error occurred generating prompts: %s\n", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error occurred generating prompts",
		})
		return
	}

	prompts = strings.Trim(prompts, "\n")
	promptsArr := strings.Split(prompts, "|")
	for _, prompt := range promptsArr {
		go GenerateUrlFromPrompt(prompt, API_KEY)
	}

	urls := make([]string, 0)
	done := make(chan struct{})
	go func() {
		for {
			select {
			case url := <-urlChan:
				urls = append(urls, url)
			case <-done:
				close(urlChan)
				return
			}
		}
	}()

loop:
	for {
		select {
		case <-errorChan:
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error occurred generating urls",
			})
			return
		default:
			if len(urls) == len(promptsArr) {
				done <- struct{}{}
				break loop
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"urls": urls,
	})
}

func GeneratePromptsFromTheme(theme string, apiKey string) (string, error) {
	data, err := json.Marshal(models.GeneratePromptsReq{
		Model:     AI_MODEL,
		Prompt:    fmt.Sprintf("Write %d detailed image prompts, each having maximum %d words, for a moodboard with the theme: %s. The prompts should not be numbered, instead separated by one '|' character", IMAGES_COUNT, PROMPT_MAX_WORDS, theme),
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
	req.Header.Add("Authorization", "Bearer "+apiKey)

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
	if err := json.Unmarshal(body, &v); err != nil || reflect.DeepEqual(v, models.GeneratePromptsResp{}) {
		log.Println("Invalid response")
		return "", errors.New("invalid response")
	}

	return v.Choices[0].Text, nil
}

func GenerateUrlFromPrompt(prompt string, apiKey string) {
	data, err := json.Marshal(models.GenerateUrlReq{
		Prompt: prompt + ", aesthetic theme",
		Size:   IMAGES_SIZE,
	})
	if err != nil {
		errorChan <- err
		log.Printf("Error marshalling data: %v\n", err)
		return
	}

	req, err := http.NewRequest(http.MethodPost, "https://api.openai.com/v1/images/generations", bytes.NewBuffer(data))
	if err != nil {
		errorChan <- err
		log.Printf("Error creating request: %v\n", err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+apiKey)

	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		errorChan <- err
		log.Printf("Error sending request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		errorChan <- err
		log.Printf("Error reading response bytes: %v\n", err)
		return
	}

	var v models.GenerateUrlResp
	if err := json.Unmarshal(body, &v); err != nil || reflect.DeepEqual(v, models.GenerateUrlResp{}) {
		errorChan <- errors.New("invalid response")
		log.Println("Invalid response")
		return
	}

	urlChan <- v.Data[0].Url
}
