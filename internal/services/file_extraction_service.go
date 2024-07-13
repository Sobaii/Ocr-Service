package services

import (
	"context"
	"log"
	"ocr-service-dev/internal/models"

	"github.com/aws/aws-sdk-go-v2/service/textract"
	"github.com/aws/aws-sdk-go-v2/service/textract/types"
)

// @tonyqiu123 todo
func ExtractExpenseData(fileBytes []byte, ctx context.Context, textractClient *textract.Client) (models.Expense, error) {
	input := &textract.AnalyzeExpenseInput{
		Document: &types.Document{
			Bytes: fileBytes,
		},
	}
	results, err := textractClient.AnalyzeExpense(ctx, input)

	if err != nil {
		return models.Expense{}, err
	}

	log.Println(results)
	return models.Expense{}, nil
}
