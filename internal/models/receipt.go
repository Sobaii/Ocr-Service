package models

type Receipt struct {
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