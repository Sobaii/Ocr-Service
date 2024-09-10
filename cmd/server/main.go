package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Sobaii/Ocr-Service/internal/api"
	"github.com/Sobaii/Ocr-Service/internal/service"
	"github.com/Sobaii/Ocr-Service/pkg/openai"
)

func main() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY environment variable is not set")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	openAIClient := openai.NewClient(apiKey)
	ocrService := service.NewOCRService(openAIClient)
	handler := api.NewHandler(ocrService)

	http.HandleFunc("/analyze", handler.AnalyzeReceipt)

	// Simple GET request example
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}