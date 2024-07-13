package services

import (
	"bytes"
	"fmt"
	"image/png"
	"os"

	"github.com/gen2brain/go-fitz"
)

func ConvertPDFToPNG(pdfBytes []byte) ([]byte, error) {
	// Create a temporary file to write the byte array
	tmpFile, err := os.CreateTemp("", "temp-*.pdf")
	if err != nil {
		return nil, fmt.Errorf("error creating temp file: %w", err)
	}
	defer os.Remove(tmpFile.Name())

	// Write the byte array to the temp file
	if _, err := tmpFile.Write(pdfBytes); err != nil {
		return nil, fmt.Errorf("error writing to temp file: %w", err)
	}
	tmpFile.Close()

	// Open the PDF document
	doc, err := fitz.New(tmpFile.Name())
	if err != nil {
		return nil, fmt.Errorf("error opening PDF document: %w", err)
	}
	defer doc.Close()

	// Render the first page of the PDF to an image
	img, err := doc.Image(0)
	if err != nil {
		return nil, fmt.Errorf("error getting image from page: %w", err)
	}

	// Create a buffer to write the PNG image
	buf := new(bytes.Buffer)
	err = png.Encode(buf, img)
	if err != nil {
		return nil, fmt.Errorf("error encoding image to PNG: %w", err)
	}

	return buf.Bytes(), nil
}
