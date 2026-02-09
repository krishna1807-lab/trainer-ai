package controller

import (
	"encoding/json"
	"net/http"
	"trainer-ai/internal/model"
	"trainer-ai/internal/service"
)

func ChatHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", 405)
		return
	}

	var req model.ChatRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", 400)
		return
	}

	reply, err := service.CallGroq(req.Prompt)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"response": reply,
	})
}
