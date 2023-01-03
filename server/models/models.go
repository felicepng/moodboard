package models

type MoodboardJson struct {
	Theme string `json:"theme"`
}

type GeneratePromptsReq struct {
	Model     string `json:"model"`
	Prompt    string `json:"prompt"`
	MaxTokens int    `json:"max_tokens"`
}
