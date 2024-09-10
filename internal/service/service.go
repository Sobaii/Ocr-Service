package service

import (
	"encoding/json"
	"fmt"

	"github.com/Sobaii/Ocr-Service/internal/models"
	"github.com/Sobaii/Ocr-Service/pkg/openai"
)

type OCRService struct {
	openAIClient *openai.Client
}

func NewOCRService(openAIClient *openai.Client) *OCRService {
	return &OCRService{openAIClient: openAIClient}
}

func (s *OCRService) AnalyzeReceipt(imageURL string) (*models.Receipt, error) {
	content, err := s.openAIClient.AnalyzeImage(imageURL)
	if err != nil {
		return nil, fmt.Errorf("failed to analyze image: %w", err)
	}

	var receipt models.Receipt
	err = json.Unmarshal([]byte(content), &receipt)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal receipt data: %w", err)
	}

	if receipt.Total == 0 || receipt.TransactionDate == "" {
		return nil, fmt.Errorf("invalid receipt data: missing required fields")
	}

	return &receipt, nil
}
