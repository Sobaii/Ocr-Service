package handlers

import (
	"context"
	"encoding/json"
	"log"
	"ocr-service-dev/internal/models"
	pb "ocr-service-dev/internal/proto"
	"ocr-service-dev/internal/services"
	"ocr-service-dev/internal/utils"

	"github.com/google/uuid"
)

type OcrServiceHandler struct {
	pb.UnimplementedOcrServiceServer
}

// test connection
func (h *OcrServiceHandler) TestConnection(ctx context.Context, req *pb.TestRequest) (*pb.TestResponse, error) {
	return &pb.TestResponse{Response: "Connection successful. Request message is: " + req.GetMessage()}, nil
}

func (h *OcrServiceHandler) SearchFileData(ctx context.Context, req *pb.SearchRequest) (*pb.Expenses, error) {

	rclient, err := utils.InitializeRedisClient()
	if err != nil {
		return nil, err
	}
	defer rclient.Close()

	if req.Index == "" && req.Query == "" {
		cmd := rclient.B().FtSearch().Index("FILE_NAME").Query("*").Limit().OffsetNum(0, 1000).Dialect(2).Build()
		n, resp, err := rclient.Do(ctx, cmd).AsFtSearch()
		if err != nil {
			return nil, err
		}

		log.Printf("%d total documents found.", n)
		var expenseItem models.Expense
		var expenseList []*pb.ExtractResponse

		for _, item := range resp {
			jsonData, ok := item.Doc["$"]
			if !ok {
				log.Println("Missing $ key in document")
				continue
			}
			if err := json.Unmarshal([]byte(jsonData), &expenseItem); err != nil {
				log.Println("Error unmarshalling JSON:", err)
				continue
			}
			expenseList = append(expenseList, &pb.ExtractResponse{
				FILE_PAGE:            &pb.ExpenseField{Text: expenseItem.FilePage.Text, Confidence: expenseItem.FilePage.Confidence},
				FILE_NAME:            &pb.ExpenseField{Text: expenseItem.FileName.Text, Confidence: expenseItem.FileName.Confidence},
				INVOICE_RECEIPT_DATE: &pb.ExpenseField{Text: expenseItem.InvoiceReceiptDate.Text, Confidence: expenseItem.InvoiceReceiptDate.Confidence},
				VENDOR_NAME:          &pb.ExpenseField{Text: expenseItem.VendorName.Text, Confidence: expenseItem.VendorName.Confidence},
				VENDOR_ADDRESS:       &pb.ExpenseField{Text: expenseItem.VendorAddress.Text, Confidence: expenseItem.VendorAddress.Confidence},
				TOTAL:                &pb.ExpenseField{Text: expenseItem.Total.Text, Confidence: expenseItem.Total.Confidence},
				SUBTOTAL:             &pb.ExpenseField{Text: expenseItem.Subtotal.Text, Confidence: expenseItem.Subtotal.Confidence},
				TAX:                  &pb.ExpenseField{Text: expenseItem.Tax.Text, Confidence: expenseItem.Tax.Confidence},
				VENDOR_PHONE:         &pb.ExpenseField{Text: expenseItem.VendorPhone.Text, Confidence: expenseItem.VendorPhone.Confidence},
				STREET:               &pb.ExpenseField{Text: expenseItem.Street.Text, Confidence: expenseItem.Street.Confidence},
				GRATUITY:             &pb.ExpenseField{Text: expenseItem.Gratuity.Text, Confidence: expenseItem.Gratuity.Confidence},
				CITY:                 &pb.ExpenseField{Text: expenseItem.City.Text, Confidence: expenseItem.City.Confidence},
				STATE:                &pb.ExpenseField{Text: expenseItem.State.Text, Confidence: expenseItem.State.Confidence},
				COUNTRY:              &pb.ExpenseField{Text: expenseItem.Country.Text, Confidence: expenseItem.Country.Confidence},
				ZIP_CODE:             &pb.ExpenseField{Text: expenseItem.ZipCode.Text, Confidence: expenseItem.ZipCode.Confidence},
				CATEGORY:             &pb.ExpenseField{Text: expenseItem.Category.Text, Confidence: expenseItem.Category.Confidence},
			})
		}
		return &pb.Expenses{
			Info: expenseList,
		}, nil
	}
	return &pb.Expenses{}, nil
}

// extract file data
func (h *OcrServiceHandler) ExtractFileData(ctx context.Context, req *pb.ExtractRequest) (*pb.ExtractResponse, error) {
	log.Printf("Received file data of length: %d", len(req.Binary))

	rclient, err := utils.InitializeRedisClient()
	if err != nil {
		return nil, err
	}
	defer rclient.Close()

	// Get mock data
	expenseItem := services.GetMockData()

	dataDoc, jsonErr := json.Marshal(expenseItem)
	if jsonErr != nil {
		log.Println(jsonErr)
		return nil, err
	}
	docId := uuid.New().String()
	cmd := rclient.B().JsonSet().Key(docId).Path("$").Value(string(dataDoc)).Build()

	cmdErr := rclient.Do(ctx, cmd).Error()

	if cmdErr != nil {
		log.Println(jsonErr)
		return nil, err
	}

	// Return pb.ExtractResponse
	response := &pb.ExtractResponse{
		FILE_PAGE:            &pb.ExpenseField{Text: expenseItem.FilePage.Text, Confidence: expenseItem.FilePage.Confidence},
		FILE_NAME:            &pb.ExpenseField{Text: expenseItem.FileName.Text, Confidence: expenseItem.FileName.Confidence},
		INVOICE_RECEIPT_DATE: &pb.ExpenseField{Text: expenseItem.InvoiceReceiptDate.Text, Confidence: expenseItem.InvoiceReceiptDate.Confidence},
		VENDOR_NAME:          &pb.ExpenseField{Text: expenseItem.VendorName.Text, Confidence: expenseItem.VendorName.Confidence},
		VENDOR_ADDRESS:       &pb.ExpenseField{Text: expenseItem.VendorAddress.Text, Confidence: expenseItem.VendorAddress.Confidence},
		TOTAL:                &pb.ExpenseField{Text: expenseItem.Total.Text, Confidence: expenseItem.Total.Confidence},
		SUBTOTAL:             &pb.ExpenseField{Text: expenseItem.Subtotal.Text, Confidence: expenseItem.Subtotal.Confidence},
		TAX:                  &pb.ExpenseField{Text: expenseItem.Tax.Text, Confidence: expenseItem.Tax.Confidence},
		VENDOR_PHONE:         &pb.ExpenseField{Text: expenseItem.VendorPhone.Text, Confidence: expenseItem.VendorPhone.Confidence},
		STREET:               &pb.ExpenseField{Text: expenseItem.Street.Text, Confidence: expenseItem.Street.Confidence},
		GRATUITY:             &pb.ExpenseField{Text: expenseItem.Gratuity.Text, Confidence: expenseItem.Gratuity.Confidence},
		CITY:                 &pb.ExpenseField{Text: expenseItem.City.Text, Confidence: expenseItem.City.Confidence},
		STATE:                &pb.ExpenseField{Text: expenseItem.State.Text, Confidence: expenseItem.State.Confidence},
		COUNTRY:              &pb.ExpenseField{Text: expenseItem.Country.Text, Confidence: expenseItem.Country.Confidence},
		ZIP_CODE:             &pb.ExpenseField{Text: expenseItem.ZipCode.Text, Confidence: expenseItem.ZipCode.Confidence},
		CATEGORY:             &pb.ExpenseField{Text: expenseItem.Category.Text, Confidence: expenseItem.Category.Confidence},
	}

	return response, nil
}
