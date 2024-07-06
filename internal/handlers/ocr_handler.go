package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"ocr-service-dev/internal/models"
	pb "ocr-service-dev/internal/proto"
	"ocr-service-dev/internal/utils"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/textract"
	"github.com/aws/aws-sdk-go-v2/service/textract/types"
	"github.com/google/uuid"
)

type OcrServiceHandler struct {
	pb.UnimplementedOcrServiceServer
	TextractClient *textract.Client
	S3Client       *s3.Client
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
				ObjectUrl:            expenseItem.ObjectUrl,
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
	docId := uuid.New().String()

	rclient, err := utils.InitializeRedisClient()
	if err != nil {
		return nil, err
	}
	defer rclient.Close()

	// upload req.GetBinary() to s3 using docId as the identifier
	fileType := "pdf"
	key := fmt.Sprintf("%s.%s", docId, fileType)
	objectURL := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", "sobaii-expenses-us-east-1", key)
	_, err = h.S3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String("sobaii-expenses-us-east-1"),
		Key:         aws.String(key),
		ContentType: aws.String("application/pdf"),
		Body:        bytes.NewReader(req.GetBinary()),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to upload file to S3: %w", err)
	}

	input := &textract.AnalyzeExpenseInput{
		Document: &types.Document{
			Bytes: req.GetBinary(),
		},
	}
	results, err := h.TextractClient.AnalyzeExpense(ctx, input)

	if err != nil {
		return nil, err
	}

	response := &pb.ExtractResponse{}
	for _, doc := range results.ExpenseDocuments {
		for _, field := range doc.SummaryFields {
			text := ""
			var confidence float64

			if field.ValueDetection.Text == nil || field.ValueDetection.Confidence == nil {
				continue
			}
			text = *field.ValueDetection.Text
			confidence = float64(*field.ValueDetection.Confidence)

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
	response.ObjectUrl = objectURL

	dataDoc, jsonErr := json.Marshal(response)
	if jsonErr != nil {
		log.Println(jsonErr)
		return nil, err
	}
	cmd := rclient.B().JsonSet().Key(docId).Path("$").Value(string(dataDoc)).Build()

	cmdErr := rclient.Do(ctx, cmd).Error()

	if cmdErr != nil {
		log.Println(jsonErr)
		return nil, err
	}

	return response, nil
}
