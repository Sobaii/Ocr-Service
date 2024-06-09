package models

func MockData() Expense {
	return Expense{
		FilePage: ExpenseField{
			Text:       "Page 1",
			Confidence: 0.95,
		},
		FileName: ExpenseField{
			Text:       "invoice.pdf",
			Confidence: 0.98,
		},
		InvoiceReceiptDate: ExpenseField{
			Text:       "2023-05-12",
			Confidence: 0.99,
		},
		VendorName: ExpenseField{
			Text:       "ACME Corp",
			Confidence: 0.97,
		},
		VendorAddress: ExpenseField{
			Text:       "123 Elm Street",
			Confidence: 0.96,
		},
		Total: ExpenseField{
			Text:       "$123.45",
			Confidence: 0.99,
		},
		Subtotal: ExpenseField{
			Text:       "$100.00",
			Confidence: 0.98,
		},
		Tax: ExpenseField{
			Text:       "$23.45",
			Confidence: 0.95,
		},
		VendorPhone: ExpenseField{
			Text:       "+1-234-567-890",
			Confidence: 0.94,
		},
		Street: ExpenseField{
			Text:       "123 Elm Street",
			Confidence: 0.96,
		},
		Gratuity: ExpenseField{
			Text:       "$5.00",
			Confidence: 0.92,
		},
		City: ExpenseField{
			Text:       "Springfield",
			Confidence: 0.95,
		},
		State: ExpenseField{
			Text:       "IL",
			Confidence: 0.95,
		},
		Country: ExpenseField{
			Text:       "USA",
			Confidence: 0.94,
		},
		ZipCode: ExpenseField{
			Text:       "62701",
			Confidence: 0.93,
		},
		Category: ExpenseField{
			Text:       "Office Supplies",
			Confidence: 0.96,
		},
	}
}