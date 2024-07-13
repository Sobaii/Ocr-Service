package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"ocr-service-dev/internal/models"
	pb "ocr-service-dev/internal/proto"
	"ocr-service-dev/internal/services"
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

func (h *OcrServiceHandler) CreateFolder(ctx context.Context, req *pb.FolderCreationRequest) (*pb.FolderCreationResponse, error) {

	rclient, err := utils.InitializeRedisClient()
	if err != nil {
		return nil, err
	}
	defer rclient.Close()

	userHashKey := fmt.Sprintf("user:%s", req.EmailAddress)
	foldersKey := fmt.Sprintf("%s:folders", userHashKey)

	userExistsCmd := rclient.B().Exists().Key(userHashKey).Build()
	existsResp, err := rclient.Do(ctx, userExistsCmd).AsInt64()
	if err != nil {
		return nil, err
	}
	if existsResp == 0 {
		// create a new user hash
		createUserCmd := rclient.B().Hset().Key(userHashKey).FieldValue().FieldValue("name", req.FullName).Build()

		err := rclient.Do(ctx, createUserCmd).Error()
		if err != nil {
			return nil, err
		}
	}

	// check if folder exists
	folderExistsCmd := rclient.B().Sismember().Key(foldersKey).Member(req.FolderName).Build()
	folderExistsResp, err := rclient.Do(ctx, folderExistsCmd).AsInt64()
	if err != nil {
		return nil, err
	}
	if folderExistsResp == 1 {
		return &pb.FolderCreationResponse{
			FolderCreated:     false,
			ActionDescription: "Folder already exists. Try again with a different folder name.",
			EmailAddress:      "",
			FullName:          "",
			FolderName:        "",
		}, nil
	}

	// add folder
	addFolderCmd := rclient.B().Sadd().Key(foldersKey).Member(req.FolderName).Build()
	if _, err := rclient.Do(ctx, addFolderCmd).AsInt64(); err != nil {
		return nil, err
	}

	return &pb.FolderCreationResponse{
		FolderCreated:     true,
		ActionDescription: "Folder successfully created.",
		EmailAddress:      req.EmailAddress,
		FullName:          req.FullName,
		FolderName:        req.FolderName,
	}, nil
}

func (h *OcrServiceHandler) SearchFolders(ctx context.Context, req *pb.FolderSearchRequest) (*pb.FolderSearchResponse, error) {

	rclient, err := utils.InitializeRedisClient()
	if err != nil {
		return nil, err
	}
	defer rclient.Close()

	userHashKey := fmt.Sprintf("user:%s", req.EmailAddress)
	foldersKey := fmt.Sprintf("%s:folders", userHashKey)

	userExistsCmd := rclient.B().Exists().Key(userHashKey).Build()
	existsResp, err := rclient.Do(ctx, userExistsCmd).AsInt64()
	if err != nil {
		return nil, err
	}
	if existsResp == 0 {
		return &pb.FolderSearchResponse{
			FolderFound:       false,
			ActionDescription: "Invalid user email.",
			Folders:           make([]string, 0),
		}, err
	}

	// TODO SearchRequest params

	retrieveFoldersCmd := rclient.B().Smembers().Key(foldersKey).Build()
	foldersResp, err := rclient.Do(ctx, retrieveFoldersCmd).AsStrSlice()
	if err != nil {
		return nil, err
	}

	return &pb.FolderSearchResponse{
		FolderFound:       true,
		ActionDescription: "Folders successfully retrieved.",
		Folders:           foldersResp,
	}, nil
}

func (h *OcrServiceHandler) SearchFileData(ctx context.Context, req *pb.SearchFileRequest) (*pb.SearchFileResponse, error) {

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
		var expenseItem models.ExtractedFile
		var expenseList []*pb.ExpenseItem

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
			expenseList = append(expenseList, &pb.ExpenseItem{
				FolderName: expenseItem.FolderName,
				Data: &pb.FileExtract{
					FilePage:           &pb.ExpenseField{Text: expenseItem.File.Data.FilePage.Text, Confidence: expenseItem.File.Data.FilePage.Confidence},
					FileName:           &pb.ExpenseField{Text: expenseItem.File.Data.FileName.Text, Confidence: expenseItem.File.Data.FileName.Confidence},
					InvoiceReceiptDate: &pb.ExpenseField{Text: expenseItem.File.Data.InvoiceReceiptDate.Text, Confidence: expenseItem.File.Data.InvoiceReceiptDate.Confidence},
					VendorName:         &pb.ExpenseField{Text: expenseItem.File.Data.VendorName.Text, Confidence: expenseItem.File.Data.VendorName.Confidence},
					VendorAddress:      &pb.ExpenseField{Text: expenseItem.File.Data.VendorAddress.Text, Confidence: expenseItem.File.Data.VendorAddress.Confidence},
					Total:              &pb.ExpenseField{Text: expenseItem.File.Data.Total.Text, Confidence: expenseItem.File.Data.Total.Confidence},
					Subtotal:           &pb.ExpenseField{Text: expenseItem.File.Data.Subtotal.Text, Confidence: expenseItem.File.Data.Subtotal.Confidence},
					Tax:                &pb.ExpenseField{Text: expenseItem.File.Data.Tax.Text, Confidence: expenseItem.File.Data.Tax.Confidence},
					VendorPhone:        &pb.ExpenseField{Text: expenseItem.File.Data.VendorPhone.Text, Confidence: expenseItem.File.Data.VendorPhone.Confidence},
					Street:             &pb.ExpenseField{Text: expenseItem.File.Data.Street.Text, Confidence: expenseItem.File.Data.Street.Confidence},
					Gratuity:           &pb.ExpenseField{Text: expenseItem.File.Data.Gratuity.Text, Confidence: expenseItem.File.Data.Gratuity.Confidence},
					City:               &pb.ExpenseField{Text: expenseItem.File.Data.City.Text, Confidence: expenseItem.File.Data.City.Confidence},
					State:              &pb.ExpenseField{Text: expenseItem.File.Data.State.Text, Confidence: expenseItem.File.Data.State.Confidence},
					Country:            &pb.ExpenseField{Text: expenseItem.File.Data.Country.Text, Confidence: expenseItem.File.Data.Country.Confidence},
					ZipCode:            &pb.ExpenseField{Text: expenseItem.File.Data.ZipCode.Text, Confidence: expenseItem.File.Data.ZipCode.Confidence},
					Category:           &pb.ExpenseField{Text: expenseItem.File.Data.Category.Text, Confidence: expenseItem.File.Data.Category.Confidence},
					ObjectUrl:          expenseItem.File.Data.ObjectUrl,
					PreviewUrl:         expenseItem.File.Data.PreviewUrl,
				},
			})
		}
		return &pb.SearchFileResponse{
			FileFound:         true,
			ActionDescription: "Files successfully retrieved.",
			EmailAddress:      req.EmailAddress,
			FolderName:        req.FolderName,
			Expenses: &pb.Expenses{
				Info: expenseList,
			},
		}, nil
	}
	return &pb.SearchFileResponse{}, nil
}

// extract file data
func (h *OcrServiceHandler) ExtractFileData(ctx context.Context, req *pb.ExtractFileRequest) (*pb.ExtractFileResponse, error) {
	log.Printf("Received file data of length: %d", len(req.Binary))
	docId := uuid.New().String()

	rclient, err := utils.InitializeRedisClient()
	if err != nil {
		return nil, err
	}
	defer rclient.Close()

	// upload req.GetBinary() to s3 using docId as the identifier
	var fileExtension string
	var s3ContentType string
	switch req.MimeType {
	case pb.MimeType_APPLICATION_PDF:
		fileExtension = "pdf"
		s3ContentType = "application/pdf"
	case pb.MimeType_IMAGE_JPEG:
		fileExtension = "jpeg"
		s3ContentType = "image/jpeg"
	case pb.MimeType_IMAGE_PNG:
		fileExtension = "png"
		s3ContentType = "image/png"
	default:
		return nil, fmt.Errorf("unsupported MIME type")
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

	response := &pb.ExtractFileResponse{
		FileExtracted:     true,
		ActionDescription: "File data successfully extracted.",
		EmailAddress:      req.EmailAddress,
		FolderName:        req.FolderName,
		File: &pb.ExpenseItem{
			FolderName: req.FolderName,
			Data:       &pb.FileExtract{},
		},
	}
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
				response.File.Data.FilePage = &pb.ExpenseField{Text: text, Confidence: confidence}
			case "FILE_NAME":
				response.File.Data.FileName = &pb.ExpenseField{Text: text, Confidence: confidence}
			case "INVOICE_RECEIPT_DATE":
				response.File.Data.InvoiceReceiptDate = &pb.ExpenseField{Text: text, Confidence: confidence}
			case "VENDOR_NAME":
				response.File.Data.VendorName = &pb.ExpenseField{Text: text, Confidence: confidence}
			case "VENDOR_ADDRESS":
				response.File.Data.VendorAddress = &pb.ExpenseField{Text: text, Confidence: confidence}
			case "TOTAL":
				response.File.Data.Total = &pb.ExpenseField{Text: text, Confidence: confidence}
			case "SUBTOTAL":
				response.File.Data.Subtotal = &pb.ExpenseField{Text: text, Confidence: confidence}
			case "TAX":
				response.File.Data.Tax = &pb.ExpenseField{Text: text, Confidence: confidence}
			case "VENDOR_PHONE":
				response.File.Data.VendorPhone = &pb.ExpenseField{Text: text, Confidence: confidence}
			case "STREET":
				response.File.Data.Street = &pb.ExpenseField{Text: text, Confidence: confidence}
			case "GRATUITY":
				response.File.Data.Gratuity = &pb.ExpenseField{Text: text, Confidence: confidence}
			case "CITY":
				response.File.Data.City = &pb.ExpenseField{Text: text, Confidence: confidence}
			case "STATE":
				response.File.Data.State = &pb.ExpenseField{Text: text, Confidence: confidence}
			case "COUNTRY":
				response.File.Data.Country = &pb.ExpenseField{Text: text, Confidence: confidence}
			case "ZIP_CODE":
				response.File.Data.ZipCode = &pb.ExpenseField{Text: text, Confidence: confidence}
			case "CATEGORY":
				response.File.Data.Category = &pb.ExpenseField{Text: text, Confidence: confidence}
			}
		}
	}

	key := fmt.Sprintf("%s/%s/%s.%s", req.EmailAddress, req.FolderName, docId, fileExtension)
	objectURL := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", "sobaii-expenses-us-east-1", key)
	_, err = h.S3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String("sobaii-expenses-us-east-1"),
		Key:         aws.String(key),
		ContentType: aws.String(s3ContentType),
		Body:        bytes.NewReader(req.GetBinary()),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to upload file to S3: %w", err)
	}

	response.File.Data.ObjectUrl = objectURL

	if s3ContentType == "application/pdf" {
		pngBytes, err := services.ConvertPDFToPNG(req.GetBinary())
		if err != nil {
			return nil, fmt.Errorf("failed to convert PDF to PNG: %w", err)
		}
		key = fmt.Sprintf("%s/%s/%s-preview.%s", req.EmailAddress, req.FolderName, docId, "png")
		objectURL = fmt.Sprintf("https://%s.s3.amazonaws.com/%s", "sobaii-expenses-us-east-1", key)

		_, err = h.S3Client.PutObject(ctx, &s3.PutObjectInput{
			Bucket:      aws.String("sobaii-expenses-us-east-1"),
			Key:         aws.String(key),
			ContentType: aws.String("image/png"),
			Body:        bytes.NewReader(pngBytes),
		})
		if err != nil {
			return nil, fmt.Errorf("failed to upload file to S3: %w", err)
		}
	}

	response.File.Data.PreviewUrl = objectURL

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
