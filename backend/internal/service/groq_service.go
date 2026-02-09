package service

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"trainer-ai/internal/config"
	"trainer-ai/internal/memory"
	"trainer-ai/internal/model"
)

func CallGroq(prompt string) (string, error) {

	apiKey := config.GetEnv("GROQ_API_KEY")
	modelName := config.GetEnv("GROQ_MODEL")
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
		MaxTokens:   300,
	}

	jsonData, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest(
		"POST",
		"https://api.groq.com/openai/v1/chat/completions",
		bytes.NewBuffer(jsonData),
	)

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var groqResp model.GroqResponse
	json.Unmarshal(body, &groqResp)

	reply := groqResp.Choices[0].Message.Content

	// Save memory
	memory.AddMessage(model.Message{Role: "user", Content: prompt})
	memory.AddMessage(model.Message{Role: "assistant", Content: reply})

	return reply, nil
}
