package main

import (
	"fmt"
	"net/http"
	"trainer-ai/internal/config"
	"trainer-ai/internal/controller"

	"github.com/MadAppGang/httplog"
)

func main() {

	config.LoadEnv()

	port := config.GetEnv("PORT")
	if port == "" {
		port = "8080"
	}

	// Convert HandlerFunc → Handler
	loggedChatHandler := httplog.Logger(
		http.HandlerFunc(controller.ChatHandler),
	)

	http.Handle("/chat", loggedChatHandler)
	http.HandleFunc("/upload-doc", controller.UploadDocumentHandler)

	fmt.Println("🚀 Server running on port", port)

	http.ListenAndServe(":"+port, nil)
}
