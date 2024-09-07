package api

import (
	"encoding/json"
	"net/http"
	"github.com/yourusername/ocr-service/internal/service"
)

type Handler struct {
	ocrService *service.OCRService
}

func NewHandler(ocrService *service.OCRService) *Handler {
	return &Handler{ocrService: ocrService}
}

func (h *Handler) AnalyzeReceipt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		ImageURL string `json:"imageUrl"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	receipt, err := h.ocrService.AnalyzeReceipt(request.ImageURL)
	if err != nil {
		http.Error(w, "Error analyzing receipt", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(receipt)
}