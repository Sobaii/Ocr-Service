package service

import (
	"encoding/json"
	"github.com/yourusername/ocr-service/internal/models"
	"github.com/yourusername/ocr-service/pkg/openai"
)

type OCRService struct {
	openAIClient *openai.Client
}

func NewOCRService(apiKey string) *OCRService {
	return &OCRService{
		openAIClient: openai.NewClient(apiKey),
	}
}

func (s *OCRService) AnalyzeReceipt(imageURL string) (*models.Receipt, error) {
	result, err := s.openAIClient.AnalyzeImage(imageURL)
	if err != nil {
		return nil, err
	}

	var receipt models.Receipt
	err = json.Unmarshal([]byte(result), &receipt)
	if err != nil {
		return nil, err
	}

	return &receipt, nil
}