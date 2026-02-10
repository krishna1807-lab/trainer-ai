package utils

import (
	"encoding/json"
	"net/http"
	"trainer-ai/internal/model"
)

func JSON(w http.ResponseWriter, status int, resp model.APIResponse) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp)
}

func Success(w http.ResponseWriter, message string, data interface{}) {
	JSON(w, http.StatusOK, model.APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func Error(w http.ResponseWriter, status int, message string) {
	JSON(w, status, model.APIResponse{
		Success: false,
		Message: message,
	})
}
