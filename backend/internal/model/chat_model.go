package model

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Prompt string `json:"prompt"`
}

type GroqRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float64   `json:"temperature"`
	MaxTokens   int       `json:"max_tokens"`
}

type GroqResponse struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}
