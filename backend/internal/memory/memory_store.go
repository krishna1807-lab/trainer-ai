package memory

import (
	"os"
	"strconv"
	"trainer-ai/internal/model"
)

var ChatHistory []model.Message

func AddMessage(msg model.Message) {

	maxMemory, _ := strconv.Atoi(os.Getenv("MAX_MEMORY"))

	ChatHistory = append(ChatHistory, msg)

	if len(ChatHistory) > maxMemory {
		ChatHistory = ChatHistory[len(ChatHistory)-maxMemory:]
	}
}

func GetHistory() []model.Message {
	return ChatHistory
}
