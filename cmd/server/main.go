package main

import (
	"log"
	"net/http"
	"github.com/yourusername/ocr-service/internal/api"
	"github.com/yourusername/ocr-service/internal/config"
	"github.com/yourusername/ocr-service/internal/service"
)

func main() {
	cfg := config.Load()

	ocrService := service.NewOCRService(cfg.OpenAIAPIKey)
	handler := api.NewHandler(ocrService)

	http.HandleFunc("/analyze", handler.AnalyzeReceipt)

	log.Printf("Server starting on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}