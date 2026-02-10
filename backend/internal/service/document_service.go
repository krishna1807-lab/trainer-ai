package service

import (
	"bytes"
	"errors"
	"fmt"
	"sync"

	"github.com/ledongthuc/pdf"
)

var (
	DocumentKnowledge string
	docMutex          sync.RWMutex
)

// ⭐ Extract Text From PDF
func ExtractTextFromPDF(fileBytes []byte) (string, error) {

	if len(fileBytes) == 0 {
		return "", errors.New("empty file")
	}

	reader := bytes.NewReader(fileBytes)

	pdfReader, err := pdf.NewReader(reader, int64(len(fileBytes)))
	if err != nil {
		return "", err
	}

	var text string

	totalPage := pdfReader.NumPage()

	for i := 1; i <= totalPage; i++ {

		page := pdfReader.Page(i)

		if page.V.IsNull() {
			continue
		}

		pageText, err := page.GetPlainText(nil)
		if err != nil {
			continue
		}

		text += pageText + "\n"
	}

	fmt.Println("Extracted text size:", len(text))

	if len(text) == 0 {
		return "", errors.New("no text extracted")
	}

	return text, nil
}

// ⭐ Save Document Knowledge (Thread Safe)
func SaveDocumentText(text string) {

	docMutex.Lock()
	defer docMutex.Unlock()

	DocumentKnowledge += "\n" + text
}

// ⭐ Get Knowledge
func GetDocumentKnowledge() string {

	docMutex.RLock()
	defer docMutex.RUnlock()

	return DocumentKnowledge
}

// ⭐ Optional: Reset Knowledge (Future Feature)
func ResetDocumentKnowledge() {

	docMutex.Lock()
	defer docMutex.Unlock()

	DocumentKnowledge = ""
}
