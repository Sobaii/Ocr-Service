package models

// ExpenseField represents a field in an expense with text and confidence
type ExpenseField struct {
	Text       string  `json:"text"`
	Confidence float64 `json:"confidence"`
}

// Expense represents the details of an expense
type Expense struct {
	FilePage            ExpenseField `json:"FILE_PAGE"`
	FileName            ExpenseField `json:"FILE_NAME"`
	InvoiceReceiptDate  ExpenseField `json:"INVOICE_RECEIPT_DATE"`
	VendorName          ExpenseField `json:"VENDOR_NAME"`
	VendorAddress       ExpenseField `json:"VENDOR_ADDRESS"`
	Total               ExpenseField `json:"TOTAL"`
	Subtotal            ExpenseField `json:"SUBTOTAL"`
	Tax                 ExpenseField `json:"TAX"`
	VendorPhone         ExpenseField `json:"VENDOR_PHONE"`
	Street              ExpenseField `json:"STREET"`
	Gratuity            ExpenseField `json:"GRATUITY"`
	City                ExpenseField `json:"CITY"`
	State               ExpenseField `json:"STATE"`
	Country             ExpenseField `json:"COUNTRY"`
	ZipCode             ExpenseField `json:"ZIP_CODE"`
	Category            ExpenseField `json:"CATEGORY"`
}

