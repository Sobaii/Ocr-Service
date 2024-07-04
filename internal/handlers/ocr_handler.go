package handlers

import (
	"context"
	"log"
	pb "ocr-service-dev/internal/proto"

	"github.com/aws/aws-sdk-go-v2/service/textract"
	"github.com/aws/aws-sdk-go-v2/service/textract/types"
)

type OcrServiceHandler struct {
	pb.UnimplementedOcrServiceServer
	Client *textract.Client
}

// TestConnection tests the connection
func (h *OcrServiceHandler) TestConnection(ctx context.Context, req *pb.TestRequest) (*pb.TestResponse, error) {
	return &pb.TestResponse{Response: "Connection successful. Request message is: " + req.GetMessage()}, nil
}

// ExtractFileData extracts file data from the binary PDF input
func (h *OcrServiceHandler) ExtractFileData(ctx context.Context, req *pb.ExtractRequest) (*pb.ExtractResponse, error) {
	log.Printf("Received file data of length: %d", len(req.Binary))

	input := &textract.AnalyzeExpenseInput{
		Document: &types.Document{
			Bytes: req.GetBinary(),
		},
	}

	result, err := h.Client.AnalyzeExpense(ctx, input)
	if err != nil {
		return nil, err
	}

	response := &pb.ExtractResponse{}

	for _, doc := range result.ExpenseDocuments {
		for _, field := range doc.SummaryFields {
			// Process fields and map them to your response struct
			text := ""
			if field.ValueDetection.Text != nil {
				text = *field.ValueDetection.Text
			}
			var confidence float64
			if field.ValueDetection.Confidence != nil {
				confidence = float64(*field.ValueDetection.Confidence)
			}
			switch *field.Type.Text {
			case "FILE_PAGE":
				response.FILE_PAGE = &pb.ExpenseField{Text: text, Confidence: confidence}
			case "FILE_NAME":
				response.FILE_NAME = &pb.ExpenseField{Text: text, Confidence: confidence}
			case "INVOICE_RECEIPT_DATE":
				response.INVOICE_RECEIPT_DATE = &pb.ExpenseField{Text: text, Confidence: confidence}
			case "VENDOR_NAME":
				response.VENDOR_NAME = &pb.ExpenseField{Text: text, Confidence: confidence}
			case "VENDOR_ADDRESS":
				response.VENDOR_ADDRESS = &pb.ExpenseField{Text: text, Confidence: confidence}
			case "TOTAL":
				response.TOTAL = &pb.ExpenseField{Text: text, Confidence: confidence}
			case "SUBTOTAL":
				response.SUBTOTAL = &pb.ExpenseField{Text: text, Confidence: confidence}
			case "TAX":
				response.TAX = &pb.ExpenseField{Text: text, Confidence: confidence}
			case "VENDOR_PHONE":
				response.VENDOR_PHONE = &pb.ExpenseField{Text: text, Confidence: confidence}
			case "STREET":
				response.STREET = &pb.ExpenseField{Text: text, Confidence: confidence}
			case "GRATUITY":
				response.GRATUITY = &pb.ExpenseField{Text: text, Confidence: confidence}
			case "CITY":
				response.CITY = &pb.ExpenseField{Text: text, Confidence: confidence}
			case "STATE":
				response.STATE = &pb.ExpenseField{Text: text, Confidence: confidence}
			case "COUNTRY":
				response.COUNTRY = &pb.ExpenseField{Text: text, Confidence: confidence}
			case "ZIP_CODE":
				response.ZIP_CODE = &pb.ExpenseField{Text: text, Confidence: confidence}
			case "CATEGORY":
				response.CATEGORY = &pb.ExpenseField{Text: text, Confidence: confidence}
			}
		}
	}

	return response, nil
}
