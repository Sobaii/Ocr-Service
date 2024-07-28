package handlers

import (
	"bytes"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	pb "ocr-service-dev/internal/proto"
	"ocr-service-dev/internal/services"
	"strings"
	"time"

	_ "github.com/lib/pq"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	s3Types "github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/aws-sdk-go-v2/service/textract"
	textractTypes "github.com/aws/aws-sdk-go-v2/service/textract/types"
	"github.com/google/uuid"
)

type OcrServiceHandler struct {
	DB *sql.DB
	pb.UnimplementedOcrServiceServer
	TextractClient *textract.Client
	S3Client       *s3.Client
}

func (h *OcrServiceHandler) CreateFolder(ctx context.Context, req *pb.FolderCreationRequest) (*pb.FolderCreationResponse, error) {

	if req.FolderName == "" {
		return &pb.FolderCreationResponse{
			FolderCreated:     false,
			ActionDescription: "Folder name cannot be blank.",
			UserId:            req.UserId,
			FolderName:        req.FolderName,
		}, nil
	}

	// TODO call auth service to validate userid

	// check if folder already exists
	var existingFolderName string
	err := h.DB.QueryRowContext(ctx, "SELECT name FROM folders WHERE clerk_user_id=$1 AND name=$2", req.UserId, req.FolderName).Scan(&existingFolderName)
	if err != nil && err != sql.ErrNoRows {
		// return nil, err
		return &pb.FolderCreationResponse{
			FolderCreated:     false,
			ActionDescription: fmt.Sprintf("Failed to select folder: %s", err.Error()),
		}, nil
	}

	if existingFolderName != "" {
		return &pb.FolderCreationResponse{
			FolderCreated:     false,
			ActionDescription: "Folder already exists.",
			UserId:            req.UserId,
			FolderName:        req.FolderName,
		}, nil
	}

	_, err = h.DB.ExecContext(ctx, "INSERT INTO folders (clerk_user_id, name) VALUES ($1, $2)", req.UserId, req.FolderName)
	if err != nil {
		// return nil, err
		return &pb.FolderCreationResponse{
			FolderCreated:     false,
			ActionDescription: fmt.Sprintf("Failed to insert folder: %s", err.Error()),
		}, nil
	}

	return &pb.FolderCreationResponse{
		FolderCreated:     true,
		ActionDescription: "Folder successfully created.",
		UserId:            req.UserId,
		FolderName:        req.FolderName,
	}, nil
}

func (h *OcrServiceHandler) SearchFolders(ctx context.Context, req *pb.FolderSearchRequest) (*pb.FolderSearchResponse, error) {

	// if empty query, select all folders
	if req.Query == "" {
		rows, err := h.DB.QueryContext(ctx, "SELECT name FROM folders WHERE clerk_user_id=$1", req.UserId)
		if err != nil {
			return &pb.FolderSearchResponse{
				FolderFound:       false,
				ActionDescription: fmt.Sprintf("Failed to select folder: %s", err.Error()),
			}, nil
		}
		folders := []string{}
		for rows.Next() {
			var f string
			err = rows.Scan(&f)
			if err != nil {
				return &pb.FolderSearchResponse{
					FolderFound:       false,
					ActionDescription: fmt.Sprintf("Failed to scan: %s", err.Error()),
				}, nil
			}
			folders = append(folders, f)
		}

		return &pb.FolderSearchResponse{
			FolderFound:       true,
			ActionDescription: "Folders successfully retrieved.",
			Folders:           folders,
		}, nil
	}

	// TODO select folder by name
	return &pb.FolderSearchResponse{}, nil

}

func (h *OcrServiceHandler) SearchFileData(ctx context.Context, req *pb.SearchFileRequest) (*pb.SearchFileResponse, error) {

	var query string
	var rows *sql.Rows
	var err error

	if req.Index == "" && req.Query == "" {
		query = `
            SELECT e.id, e.clerk_user_id, e.object_url, e.preview_url, f.name, ef.field_type, ef.text, ef.confidence
			FROM expenses e
			LEFT JOIN expense_fields ef ON e.id = ef.expense_id
			LEFT JOIN folders f ON e.folder_id = f.id
			WHERE e.clerk_user_id = $1
        `
		rows, err = h.DB.QueryContext(ctx, query, req.UserId)
	} else {
		query = `
			WITH ranked_expenses AS (
				SELECT 
					ef.expense_id, 
					ef.text, 
					similarity(ef.text, $1) AS score,
					ROW_NUMBER() OVER (PARTITION BY ef.expense_id ORDER BY similarity(ef.text, $1) DESC) AS rank
				FROM expense_fields ef
				WHERE ef.field_type = $2
			)
			SELECT e.id, e.clerk_user_id, e.object_url, e.preview_url, f.name, ef.field_type, ef.text, ef.confidence
			FROM ranked_expenses
			RIGHT JOIN expense_fields ef
			ON ranked_expenses.expense_id = ef.expense_id
			right join expenses e
			on ranked_expenses.expense_id = e.id
			right join folders f
			on e.folder_id = f.id
			WHERE e.clerk_user_id = $3
			AND ranked_expenses.rank = 1
			ORDER BY ranked_expenses.score DESC;
		`
		rows, err = h.DB.QueryContext(ctx, query, req.Query, req.Index, req.UserId)
	}

	if err != nil {
		return &pb.SearchFileResponse{
			FileFound:         false,
			ActionDescription: fmt.Sprintf("Failed to query expenses: %s", err.Error()),
		}, nil
	}
	defer rows.Close()

	var expenseMap = make(map[int]*pb.ExpenseItem)
	for rows.Next() {
		var expenseID int
		var clerkUserID, objectURL, previewURL, folderName, fieldType, text string
		var confidence float64

		err := rows.Scan(&expenseID, &clerkUserID, &objectURL, &previewURL, &folderName, &fieldType, &text, &confidence)
		if err != nil {
			return &pb.SearchFileResponse{
				FileFound:         false,
				ActionDescription: fmt.Sprintf("Failed to scan row: %s", err.Error()),
			}, nil
		}

		expense, exists := expenseMap[expenseID]
		if !exists {
			expense = &pb.ExpenseItem{
				FolderName: folderName,
				Data:       &pb.FileExtract{},
			}
			expenseMap[expenseID] = expense
		}

		expense.Data.ObjectUrl = objectURL
		expense.Data.PreviewUrl = previewURL
		expense.Data.ExpenseId = uint32(expenseID)
		switch fieldType {
		case "FILE_PAGE":
			expense.Data.FilePage = &pb.ExpenseField{FieldType: fieldType, Text: text, Confidence: confidence}
		case "FILE_NAME":
			expense.Data.FileName = &pb.ExpenseField{FieldType: fieldType, Text: text, Confidence: confidence}
		case "INVOICE_RECEIPT_DATE":
			expense.Data.InvoiceReceiptDate = &pb.ExpenseField{FieldType: fieldType, Text: text, Confidence: confidence}
		case "VENDOR_NAME":
			expense.Data.VendorName = &pb.ExpenseField{FieldType: fieldType, Text: text, Confidence: confidence}
		case "VENDOR_ADDRESS":
			expense.Data.VendorAddress = &pb.ExpenseField{FieldType: fieldType, Text: text, Confidence: confidence}
		case "TOTAL":
			expense.Data.Total = &pb.ExpenseField{FieldType: fieldType, Text: text, Confidence: confidence}
		case "SUBTOTAL":
			expense.Data.Subtotal = &pb.ExpenseField{FieldType: fieldType, Text: text, Confidence: confidence}
		case "TAX":
			expense.Data.Tax = &pb.ExpenseField{FieldType: fieldType, Text: text, Confidence: confidence}
		case "VENDOR_PHONE":
			expense.Data.VendorPhone = &pb.ExpenseField{FieldType: fieldType, Text: text, Confidence: confidence}
		case "STREET":
			expense.Data.Street = &pb.ExpenseField{FieldType: fieldType, Text: text, Confidence: confidence}
		case "GRATUITY":
			expense.Data.Gratuity = &pb.ExpenseField{FieldType: fieldType, Text: text, Confidence: confidence}
		case "CITY":
			expense.Data.City = &pb.ExpenseField{FieldType: fieldType, Text: text, Confidence: confidence}
		case "STATE":
			expense.Data.State = &pb.ExpenseField{FieldType: fieldType, Text: text, Confidence: confidence}
		case "COUNTRY":
			expense.Data.Country = &pb.ExpenseField{FieldType: fieldType, Text: text, Confidence: confidence}
		case "ZIP_CODE":
			expense.Data.ZipCode = &pb.ExpenseField{FieldType: fieldType, Text: text, Confidence: confidence}
		case "CATEGORY":
			expense.Data.Category = &pb.ExpenseField{FieldType: fieldType, Text: text, Confidence: confidence}
		}
	}

	var expenseList []*pb.ExpenseItem
	for _, expense := range expenseMap {
		expenseList = append(expenseList, expense)
	}

	return &pb.SearchFileResponse{
		FileFound:         true,
		ActionDescription: "Files successfully retrieved.",
		UserId:            req.UserId,
		FolderName:        req.FolderName,
		Expenses: &pb.Expenses{
			Info: expenseList,
		},
	}, nil
}

// extract file data
func (h *OcrServiceHandler) ExtractFileData(ctx context.Context, req *pb.ExtractFileRequest) (*pb.ExtractFileResponse, error) {
	log.Printf("Received file data of length: %d", len(req.Binary))
	docId := uuid.New().String()

	// check if folder exists
	var folderId string
	folderErr := h.DB.QueryRowContext(ctx, "select id from folders where clerk_user_id=$1 and name=$2", req.UserId, req.FolderName).Scan(&folderId)
	if folderErr != nil {
		return &pb.ExtractFileResponse{
			FileExtracted:     false,
			ActionDescription: fmt.Sprintf("Failed to select folder: %s", folderErr.Error()),
		}, nil
	}

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
		Document: &textractTypes.Document{
			Bytes: req.GetBinary(),
		},
	}
	results, err := h.TextractClient.AnalyzeExpense(ctx, input)

	if err != nil {
		return &pb.ExtractFileResponse{
			FileExtracted:     false,
			ActionDescription: fmt.Sprintf("Failed to analyze expense: %s", err.Error()),
		}, nil
	}
	log.Println("Expense data successfully extracted.")

	// response object
	response := &pb.ExtractFileResponse{
		FileExtracted:     true,
		ActionDescription: "File data successfully extracted.",
		UserId:            req.UserId,
		FolderName:        req.FolderName,
		File: &pb.ExpenseItem{
			FolderName: req.FolderName,
			Data:       &pb.FileExtract{},
		},
	}

	// upload the file to S3
	key := fmt.Sprintf("%s/%s/%s.%s", req.UserId, req.FolderName, docId, fileExtension)
	objectURL := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", "sobaii-expenses-us-east-1", key)
	previewURL := objectURL

	_, err = h.S3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String("sobaii-expenses-us-east-1"),
		Key:         aws.String(key),
		ContentType: aws.String(s3ContentType),
		Body:        bytes.NewReader(req.GetBinary()),
	})
	if err != nil {
		return &pb.ExtractFileResponse{
			FileExtracted:     false,
			ActionDescription: fmt.Sprintf("Failed to upload file to S3: %s", err.Error()),
		}, nil
	}
	response.File.Data.ObjectUrl = objectURL

	// update the preview url to be the converted png file
	if s3ContentType == "application/pdf" {
		pngBytes, err := services.ConvertPDFToPNG(req.GetBinary())
		if err != nil {
			return &pb.ExtractFileResponse{
				FileExtracted:     false,
				ActionDescription: fmt.Sprintf("failed to convert PDF to PNG: %s", err.Error()),
			}, nil

		}
		key = fmt.Sprintf("%s/%s/%s-preview.%s", req.UserId, req.FolderName, docId, "png")
		previewURL = fmt.Sprintf("https://%s.s3.amazonaws.com/%s", "sobaii-expenses-us-east-1", key)

		_, err = h.S3Client.PutObject(ctx, &s3.PutObjectInput{
			Bucket:      aws.String("sobaii-expenses-us-east-1"),
			Key:         aws.String(key),
			ContentType: aws.String("image/png"),
			Body:        bytes.NewReader(pngBytes),
		})
		if err != nil {
			return &pb.ExtractFileResponse{
				FileExtracted:     false,
				ActionDescription: fmt.Sprintf("failed to upload file to S3: %s", err.Error()),
			}, nil
		}
	}
	log.Println("Expense successfully uploaded to S3.")
	response.File.Data.PreviewUrl = previewURL

	// Insert the expense record into the database
	var expenseID int
	err = h.DB.QueryRowContext(ctx, `
        INSERT INTO expenses (clerk_user_id, object_url, preview_url, folder_id)
        VALUES ($1, $2, $3, $4)
        RETURNING id
    `, req.UserId, objectURL, previewURL, folderId).Scan(&expenseID)
	if err != nil {
		return &pb.ExtractFileResponse{
			FileExtracted:     false,
			ActionDescription: fmt.Sprintf("failed to insert expense: %s", err.Error()),
		}, nil
	}
	log.Println("Expense record successfully inserted.")
	response.File.Data.ExpenseId = uint32(expenseID)

	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic in document processing: %v", r)
		}
	}()
	log.Printf("Starting to process %d expense documents", len(results.ExpenseDocuments))

	for docIndex, doc := range results.ExpenseDocuments {
		log.Printf("Processing document %d: %d fields found.", docIndex, len(doc.SummaryFields))

		for fieldIndex, field := range doc.SummaryFields {
			log.Printf("Processing field %d", fieldIndex)

			var fieldType string
			var text string
			var confidence float64

			if field.ValueDetection.Text == nil || field.ValueDetection.Confidence == nil {
				continue
			}
			fieldType = *field.Type.Text
			text = *field.ValueDetection.Text
			confidence = float64(*field.ValueDetection.Confidence)

			insertCtx, cancel := context.WithTimeout(ctx, 3*time.Second)

			var insertErr error
			doneChan := make(chan bool)

			go func() {
				_, insertErr = h.DB.ExecContext(insertCtx, `INSERT INTO expense_fields (expense_id, field_type, text, confidence) VALUES ($1, $2, $3, $4)`, expenseID, *field.Type.Text, text, confidence)
				doneChan <- true
			}()

			select {
			case <-doneChan:
				if insertErr != nil {
					log.Printf("Failed to insert expense field: %s", insertErr.Error())
				} else {
					log.Println("Expense field successfully inserted.")
				}
			case <-insertCtx.Done():
				if insertCtx.Err() == context.DeadlineExceeded {
					log.Println("Insertion timed out after 3 seconds. Continuing with next field.")
				}
			}

			// Always cancel the context, regardless of the outcome
			cancel()

			if insertErr != nil {
				log.Printf("failed to insert expense field: %s", insertErr.Error())
				continue
			}

			log.Printf("Text: %s Confidence: %v", text, confidence)
			switch fieldType {
			case "FILE_PAGE":
				response.File.Data.FilePage = &pb.ExpenseField{FieldType: fieldType, Text: text, Confidence: confidence}
			case "FILE_NAME":
				response.File.Data.FileName = &pb.ExpenseField{FieldType: fieldType, Text: text, Confidence: confidence}
			case "INVOICE_RECEIPT_DATE":
				response.File.Data.InvoiceReceiptDate = &pb.ExpenseField{FieldType: fieldType, Text: text, Confidence: confidence}
			case "VENDOR_NAME":
				response.File.Data.VendorName = &pb.ExpenseField{FieldType: fieldType, Text: text, Confidence: confidence}
			case "VENDOR_ADDRESS":
				response.File.Data.VendorAddress = &pb.ExpenseField{FieldType: fieldType, Text: text, Confidence: confidence}
			case "TOTAL":
				response.File.Data.Total = &pb.ExpenseField{FieldType: fieldType, Text: text, Confidence: confidence}
			case "SUBTOTAL":
				response.File.Data.Subtotal = &pb.ExpenseField{FieldType: fieldType, Text: text, Confidence: confidence}
			case "TAX":
				response.File.Data.Tax = &pb.ExpenseField{FieldType: fieldType, Text: text, Confidence: confidence}
			case "VENDOR_PHONE":
				response.File.Data.VendorPhone = &pb.ExpenseField{FieldType: fieldType, Text: text, Confidence: confidence}
			case "STREET":
				response.File.Data.Street = &pb.ExpenseField{FieldType: fieldType, Text: text, Confidence: confidence}
			case "GRATUITY":
				response.File.Data.Gratuity = &pb.ExpenseField{FieldType: fieldType, Text: text, Confidence: confidence}
			case "CITY":
				response.File.Data.City = &pb.ExpenseField{FieldType: fieldType, Text: text, Confidence: confidence}
			case "STATE":
				response.File.Data.State = &pb.ExpenseField{FieldType: fieldType, Text: text, Confidence: confidence}
			case "COUNTRY":
				response.File.Data.Country = &pb.ExpenseField{FieldType: fieldType, Text: text, Confidence: confidence}
			case "ZIP_CODE":
				response.File.Data.ZipCode = &pb.ExpenseField{FieldType: fieldType, Text: text, Confidence: confidence}
			case "CATEGORY":
				response.File.Data.Category = &pb.ExpenseField{FieldType: fieldType, Text: text, Confidence: confidence}
			}
			log.Printf("Finished processing field type: %v", fieldType)
		}
	}

	return response, nil
}

func (h *OcrServiceHandler) ModifyExpenseField(ctx context.Context, req *pb.ModifyExpenseFieldRequest) (*pb.ModifyExpenseFieldResponse, error) {
	// TODO call auth service to validate userid
	var expenseId sql.NullInt32
	var fieldType, fieldText sql.NullString
	var confidence sql.NullFloat64

	query := `
	UPDATE expense_fields
	SET text=$1
	WHERE field_type=$2 AND expense_id=$3

	RETURNING expense_id, field_type, text, confidence
	`
	err := h.DB.QueryRowContext(ctx, query, req.FieldText, req.FieldType, req.ExpenseId).Scan(&expenseId, &fieldType, &fieldText, &confidence)
	if err == sql.ErrNoRows {
		return &pb.ModifyExpenseFieldResponse{
			ExpenseFieldModified: false,
			ActionDescription:    fmt.Sprintf("Failed to update expense field: %s", err.Error()),
		}, nil
	}
	if err != nil {
		return &pb.ModifyExpenseFieldResponse{
			ExpenseFieldModified: false,
			ActionDescription:    fmt.Sprintf("Failed to update expense field: %s", err.Error()),
		}, nil
	}

	return &pb.ModifyExpenseFieldResponse{
		ExpenseFieldModified: true,
		ActionDescription:    "Expense field successfully modified.",
		ExpenseId:            uint32(expenseId.Int32),
		FieldType:            fieldType.String,
		FieldText:            fieldText.String,
		Confidence:           confidence.Float64,
	}, nil
}

func (h *OcrServiceHandler) DeleteExpense(ctx context.Context, req *pb.DeleteExpenseRequest) (*pb.DeleteExpenseResponse, error) {
	// TODO call auth service to validate userid

	var objectUrl, previewUrl string
	var expenseId sql.NullInt32

	err := h.DB.QueryRowContext(ctx, "SELECT object_url, preview_url from expenses where id=$1", req.ExpenseId).Scan(&objectUrl, &previewUrl)
	if err != nil {
		return &pb.DeleteExpenseResponse{
			ExpenseDeleted:    false,
			ActionDescription: fmt.Sprintf("Failed to select expense: %s", err.Error()),
		}, nil
	}

	// retrieve s3 object metadata to check existence
	objectKey := aws.String(extractKeyFromObjectURL(objectUrl))
	_, err = h.S3Client.HeadObject(ctx, &s3.HeadObjectInput{
		Bucket: aws.String("sobaii-expenses-us-east-1"),
		Key:    objectKey,
	})
	if err != nil {
		var notFound *s3Types.NotFound
		if errors.As(err, &notFound) {
			return &pb.DeleteExpenseResponse{
				ExpenseDeleted:    false,
				ActionDescription: fmt.Sprintf("Object does not exist: %v", objectKey),
			}, nil
		}
		log.Printf("Failed to check object existence: %v", err)
		return nil, err // Some other s3 error occurred
	}

	if objectUrl == previewUrl {
		_, err = h.S3Client.DeleteObject(ctx, &s3.DeleteObjectInput{
			Bucket: aws.String("sobaii-expenses-us-east-1"),
			Key:    aws.String(extractKeyFromObjectURL(objectUrl)),
		})

	} else {
		_, err = h.S3Client.DeleteObjects(ctx, &s3.DeleteObjectsInput{
			Bucket: aws.String("sobaii-expenses-us-east-1"),
			Delete: &s3Types.Delete{
				Objects: []s3Types.ObjectIdentifier{
					{Key: objectKey},
					{Key: aws.String(extractKeyFromObjectURL(previewUrl))},
				},
			},
		})
	}
	if err != nil {
		return &pb.DeleteExpenseResponse{
			ExpenseDeleted:    false,
			ActionDescription: fmt.Sprintf("Failed to delete s3 objects: %s", err.Error()),
		}, nil
	}

	err = h.DB.QueryRowContext(ctx, "delete from expenses where id=$1 returning id", req.ExpenseId).Scan(&expenseId)
	if err != nil {
		return &pb.DeleteExpenseResponse{
			ExpenseDeleted:    false,
			ActionDescription: fmt.Sprintf("Failed to delete expense record: %s", err.Error()),
		}, nil
	}

	return &pb.DeleteExpenseResponse{
		ExpenseDeleted:    true,
		ActionDescription: "Expense record successfully deleted.",
		ExpenseId:         uint32(expenseId.Int32),
	}, nil
}

func extractKeyFromObjectURL(objectURL string) string {
	parts := strings.Split(objectURL, ".com/")
	if len(parts) == 2 {
		return parts[1]
	}
	return ""
}
