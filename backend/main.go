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

	// ⭐ Create Mux
	mux := http.NewServeMux()

	// ⭐ Wrap Chat Handler with Logger
	loggedChatHandler := httplog.Logger(
		http.HandlerFunc(controller.ChatHandler),
	)

	// ⭐ Register Routes
	mux.Handle("/chat", loggedChatHandler)
	mux.HandleFunc("/upload-doc", controller.UploadDocumentHandler)

	fmt.Println("🚀 Server running on port", port)

	// ⭐ Start Server WITH CORS Middleware
	err := http.ListenAndServe(":"+port, corsMiddleware(mux))
	if err != nil {
		fmt.Println("Server Error:", err)
	}
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")

		// ⭐ Important for preflight
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
