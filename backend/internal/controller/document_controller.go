package controller

import (
	"io"
	"net/http"
	"trainer-ai/internal/service"
)

func UploadDocumentHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", 405)
		return
	}

	// IMPORTANT
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		http.Error(w, "Cannot parse form", 400)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "File upload failed: "+err.Error(), 400)
		return
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "File read failed", 500)
		return
	}

	text, err := service.ExtractTextFromPDF(fileBytes)
	if err != nil {
		http.Error(w, "PDF extract failed: "+err.Error(), 500)
		return
	}

	service.SaveDocumentText(text)

	w.Write([]byte("PDF uploaded successfully"))
}
