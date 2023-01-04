package models

type MoodboardJson struct {
	Theme string `json:"theme"`
}

type GeneratePromptsReq struct {
	Model     string `json:"model"`
	Prompt    string `json:"prompt"`
	MaxTokens int    `json:"max_tokens"`
}

type Result struct {
	FinishReason string `json:"finish_reason"`
	Index        int    `json:"index"`
	Logprobs     *int   `json:"logprobs"`
	Text         string `json:"text"`
}

type GeneratePromptsResp struct {
	Choices []Result               `json:"choices"`
	Created int                    `json:"created"`
	Id      string                 `json:"id"`
	Model   string                 `json:"model"`
	Object  string                 `json:"object"`
	Usage   map[string]interface{} `json:"usage"`
}

type GenerateUrlReq struct {
	Prompt string `json:"prompt"`
	Size   string `json:"size"`
}

type UrlObj struct {
	Url string `json:"url"`
}

type GenerateUrlResp struct {
	Created int      `json:"created"`
	Data    []UrlObj `json:"data"`
}
