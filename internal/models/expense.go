package models

// ExpenseField represents a field in an expense with text and confidence
type ExpenseField struct {
	Text       string  `json:"text"`
	Confidence float64 `json:"confidence"`
}

// Expense represents the details of an expense
type Expense struct {
	FilePage           ExpenseField `json:"file_page"`
	FileName           ExpenseField `json:"file_name"`
	InvoiceReceiptDate ExpenseField `json:"invoice_receipt_date"`
	VendorName         ExpenseField `json:"vendor_name"`
	VendorAddress      ExpenseField `json:"vendor_address"`
	Total              ExpenseField `json:"total"`
	Subtotal           ExpenseField `json:"subtotal"`
	Tax                ExpenseField `json:"tax"`
	VendorPhone        ExpenseField `json:"vendor_phone"`
	Street             ExpenseField `json:"street"`
	Gratuity           ExpenseField `json:"gratuity"`
	City               ExpenseField `json:"city"`
	State              ExpenseField `json:"state"`
	Country            ExpenseField `json:"country"`
	ZipCode            ExpenseField `json:"zip_code"`
	Category           ExpenseField `json:"category"`
	ObjectUrl          string       `json:"object_url"`
	PreviewUrl         string       `json:"preview_url"`
}

// FileData represents the data extracted from a file
type FileData struct {
	FolderName string  `json:"folder_name"`
	Data       Expense `json:"data"`
}

// ExtractedFile represents the entire extracted file information
type ExtractedFile struct {
	FileExtracted     bool     `json:"file_extracted"`
	ActionDescription string   `json:"action_description"`
	EmailAddress      string   `json:"email_address"`
	FolderName        string   `json:"folder_name"`
	File              FileData `json:"file"`
}
