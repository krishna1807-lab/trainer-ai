package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"trainer-ai/internal/config"
	"trainer-ai/internal/memory"
	"trainer-ai/internal/model"
)

func CallGroq(prompt string) (string, error) {

	if prompt == "" {
		return "", errors.New("prompt cannot be empty")
	}

	apiKey := config.GetEnv("GROQ_API_KEY")
	modelName := config.GetEnv("GROQ_MODEL")

	if apiKey == "" {
		return "", errors.New("missing GROQ_API_KEY")
	}

	docText := GetDocumentKnowledge()

	systemPrompt := `
You are an Enterprise Trainer AI.

You MUST answer using ONLY the document knowledge provided below.

If answer is not present in document, reply:
"Not found in uploaded document".

-------------------
DOCUMENT KNOWLEDGE:
` + docText + `
-------------------
`

	// ⭐ Build Messages
	messages := []model.Message{
		{
			Role:    "system",
			Content: systemPrompt,
		},
	}

	messages = append(messages, memory.GetHistory()...)
	messages = append(messages, model.Message{
		Role:    "user",
		Content: prompt,
	})

	reqBody := model.GroqRequest{
		Model:       modelName,
		Messages:    messages,
		Temperature: 0.7,
		MaxTokens:   500,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(
		"POST",
		"https://api.groq.com/openai/v1/chat/completions",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 60 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var groqResp model.GroqResponse
	err = json.Unmarshal(body, &groqResp)
	if err != nil {
		return "", err
	}

	if len(groqResp.Choices) == 0 {
		return "", errors.New("empty response from groq")
	}

	reply := groqResp.Choices[0].Message.Content

	// ⭐ Save Memory
	memory.AddMessage(model.Message{Role: "user", Content: prompt})
	memory.AddMessage(model.Message{Role: "assistant", Content: reply})

	return reply, nil
}
