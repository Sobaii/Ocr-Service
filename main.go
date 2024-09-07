package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sashabaranov/go-openai"
)

var categories = []string{
	"advertising", "meals", "amortization", "insurance", "bank charge", "interest",
	// ... (add all categories here)
}

type ReceiptData struct {
	TransactionDate string  `json:"transactionDate"`
	Company         string  `json:"company"`
	VendorAddress   string  `json:"vendorAddress"`
	Total           float64 `json:"total"`
	Subtotal        float64 `json:"subtotal"`
	TotalTax        float64 `json:"totalTax"`
	VendorPhone     string  `json:"vendorPhone"`
	Street          string  `json:"street"`
	Gratuity        float64 `json:"gratuity"`
	City            string  `json:"city"`
	State           string  `json:"state"`
	Country         string  `json:"country"`
	ZipCode         string  `json:"zipCode"`
	Category        string  `json:"category"`
}

func analyzeFile(client *openai.Client, fileName string) (*ReceiptData, error) {
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4Vision,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: fmt.Sprintf("Extract from the first receipt you see: transactionDate, company, vendorAddress, total, subtotal, totalTax, vendorPhone, street, gratuity, city, state, country, zipCode, and category. Be very meticulous about categorizing this receipt into one of the following categories: %s. Return the result as plain JSON. Do not include any markdown or code blocks. Pay very special attention to the subtotal, total tax, and total. Date format: YYYY-MM-DD. Ensure number values are numbers, not strings. No additional text.", strings.Join(categories, ", ")),
				},
			},
			MultimodalContent: []openai.ChatMessagePart{
				{
					Type: openai.ChatMessagePartTypeImageURL,
					ImageURL: &openai.ChatMessageImageURL{
						URL:    fileName,
						Detail: openai.ImageURLDetailHigh,
					},
				},
			},
		},
	)

	if err != nil {
		return nil, fmt.Errorf("error analyzing receipt: %v", err)
	}

	var receiptData ReceiptData
	err = json.Unmarshal([]byte(resp.Choices[0].Message.Content), &receiptData)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}

	return &receiptData, nil
}

func handleAnalyzeFile(client *openai.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var request struct {
			FileURL string `json:"fileUrl"`
		}

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		result, err := analyzeFile(client, request.FileURL)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error analyzing file: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}

func main() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY environment variable is not set")
	}

	client := openai.NewClient(apiKey)

	http.HandleFunc("/analyze", handleAnalyzeFile(client))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}