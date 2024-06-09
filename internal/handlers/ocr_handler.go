package handlers

import (
	"context"
	"io"
	"log"
	pb "ocr-service-dev/internal/proto"
	"ocr-service-dev/internal/services"
)

type OcrServiceHandler struct {
	pb.UnimplementedOcrServiceServer
}

// test connection
func (h *OcrServiceHandler) TestConnection(ctx context.Context, req *pb.TestRequest) (*pb.TestResponse, error) {
	return &pb.TestResponse{Response: "Connection successful. Request message is: " + req.GetMessage()}, nil
}

func (h *OcrServiceHandler) ExtractFileData(stream pb.OcrService_ExtractFileDataServer) error {
	var finalRequest *pb.ExtractRequest
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		finalRequest = req
	}
	log.Println(finalRequest.String())

	// Get mock data
	mockData := services.GetMockData()

	// Convert mockData (models.Expense) to pb.ExtractResponse
	response := &pb.ExtractResponse{
		FILE_PAGE:            &pb.ExpenseField{Text: mockData.FilePage.Text, Confidence: mockData.FilePage.Confidence},
		FILE_NAME:            &pb.ExpenseField{Text: mockData.FileName.Text, Confidence: mockData.FileName.Confidence},
		INVOICE_RECEIPT_DATE: &pb.ExpenseField{Text: mockData.InvoiceReceiptDate.Text, Confidence: mockData.InvoiceReceiptDate.Confidence},
		VENDOR_NAME:          &pb.ExpenseField{Text: mockData.VendorName.Text, Confidence: mockData.VendorName.Confidence},
		VENDOR_ADDRESS:       &pb.ExpenseField{Text: mockData.VendorAddress.Text, Confidence: mockData.VendorAddress.Confidence},
		TOTAL:                &pb.ExpenseField{Text: mockData.Total.Text, Confidence: mockData.Total.Confidence},
		SUBTOTAL:             &pb.ExpenseField{Text: mockData.Subtotal.Text, Confidence: mockData.Subtotal.Confidence},
		TAX:                  &pb.ExpenseField{Text: mockData.Tax.Text, Confidence: mockData.Tax.Confidence},
		VENDOR_PHONE:         &pb.ExpenseField{Text: mockData.VendorPhone.Text, Confidence: mockData.VendorPhone.Confidence},
		STREET:               &pb.ExpenseField{Text: mockData.Street.Text, Confidence: mockData.Street.Confidence},
		GRATUITY:             &pb.ExpenseField{Text: mockData.Gratuity.Text, Confidence: mockData.Gratuity.Confidence},
		CITY:                 &pb.ExpenseField{Text: mockData.City.Text, Confidence: mockData.City.Confidence},
		STATE:                &pb.ExpenseField{Text: mockData.State.Text, Confidence: mockData.State.Confidence},
		COUNTRY:              &pb.ExpenseField{Text: mockData.Country.Text, Confidence: mockData.Country.Confidence},
		ZIP_CODE:             &pb.ExpenseField{Text: mockData.ZipCode.Text, Confidence: mockData.ZipCode.Confidence},
		CATEGORY:             &pb.ExpenseField{Text: mockData.Category.Text, Confidence: mockData.Category.Confidence},
	}

	return stream.SendAndClose(response)
}