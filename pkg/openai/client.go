package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Client struct {
	APIKey string
}

func NewClient(apiKey string) *Client {
	return &Client{APIKey: apiKey}
}

func (c *Client) AnalyzeImage(imageURL string) (string, error) { 
	url := "https://api.openai.com/v1/chat/completions"
	categories := []string{
		"advertising", "meals", "amortization", "insurance", "bank charge", "interest",
		"business taxes, licences & memberships", "franchise fees", "office expense",
		"professional fees", "accounting fees", "brokerage fee", "management and administration",
		"training expense", "rent", "home office", "vehicle rentals", "repairs and maintenance",
		"salary", "sub-contracts", "supplies", "small tools", "computer-related expenses",
		"internet", "property taxes", "travel", "utilities", "telephone and communications",
		"selling expense", "delivery expense", "waste expense", "vehicle expense",
		"general and administrative expense",
	}

	payload := map[string]interface{}{
		"model": "gpt-4o-mini",
		"messages": []map[string]interface{}{
			{
				"role": "user",
				"content": []map[string]interface{}{
					{
						"type": "text",
						"text": fmt.Sprintf("Extract from the first receipt you see: transactionDate, company, currency, vendorAddress, total, subtotal, totalTax, vendorPhone, street, gratuity, city, state, country, zipCode, and category. Be very meticulous about categorizing this receipt into one of the following categories: %s. Return the result as plain JSON. Do not include any markdown or code blocks. Pay very special attention to the subtotal, total tax, and total. Date format: YYYY-MM-DD. Ensure number values are numbers, not strings. No additional text.", categories),
					},
					{
						"type": "image_url",
						"image_url": map[string]interface{}{
							"url":    imageURL,
							"detail": "high",
						},
					},
				},
			},
		},
		"max_tokens": 300,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", err
	}
	
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	log.Printf("Server starting on port %s", body)

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}

	choices, ok := result["choices"].([]interface{})
	if !ok || len(choices) == 0 {
		return "", fmt.Errorf("unexpected response format")
	}

	firstChoice, ok := choices[0].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("unexpected choice format")
	}

	message, ok := firstChoice["message"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("unexpected message format")
	}

	content, ok := message["content"].(string)
	if !ok {
		return "", fmt.Errorf("unexpected content format")
	}

	return content, nil
}