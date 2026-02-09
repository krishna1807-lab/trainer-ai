package service

import (
	"bytes"
	"fmt"

	"github.com/ledongthuc/pdf"
)

func ExtractTextFromPDF(fileBytes []byte) (string, error) {

	reader := bytes.NewReader(fileBytes)

	pdfReader, err := pdf.NewReader(reader, int64(len(fileBytes)))
	if err != nil {
		return "", err
	}

	var text string

	totalPage := pdfReader.NumPage()

	for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {

		page := pdfReader.Page(pageIndex)

		if page.V.IsNull() {
			continue
		}

		pageText, err := page.GetPlainText(nil)
		if err != nil {
			continue
		}

		text += pageText + "\n"
	}

	fmt.Println("EXTRACTED TEXT SIZE:", len(text))

	return text, nil
}

var DocumentKnowledge string

func SaveDocumentText(text string) {
	DocumentKnowledge += "\n" + text
	fmt.Println("DocumentKnowledge", DocumentKnowledge)

}

func GetDocumentKnowledge() string {
	return DocumentKnowledge
}
