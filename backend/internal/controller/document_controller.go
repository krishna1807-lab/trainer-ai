package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"trainer-ai/internal/service"
)

func UploadDocumentHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, false, "Only POST allowed", nil)
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, false, "Cannot parse form", nil)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		writeJSON(w, http.StatusBadRequest, false, "File upload failed", nil)
		return
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, false, "File read failed", nil)
		return
	}

	text, err := service.ExtractTextFromPDF(fileBytes)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, false, "PDF extract failed", nil)
		return
	}

	service.SaveDocumentText(text)

	writeJSON(w, http.StatusOK, true, "PDF uploaded successfully", map[string]interface{}{
		"file_name": header.Filename,
		"text_size": len(text),
	})
}
func writeJSON(w http.ResponseWriter, status int, success bool, message string, data interface{}) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := map[string]interface{}{
		"success": success,
		"message": message,
		"data":    data,
	}

	json.NewEncoder(w).Encode(resp)
}
