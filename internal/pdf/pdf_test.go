package pdf

import (
	"errors"
	"testing"
)

type PdfWriterMock struct {
	Error error
}

func (p PdfWriterMock) Write() error {
	return p.Error
}

func TestWritePDF(t *testing.T) {
	t.Run("When PDF file is written then return error", func(t *testing.T) {
		expectedErr := "Error while creating the PDF"
		pdfWriter := &PdfWriterMock{Error: errors.New(expectedErr)}

		writeErr := WritePDF(pdfWriter)

		if writeErr == nil {
			t.Errorf("Expected %s but pdf writing was success", expectedErr)
		}
	})

	t.Run("When output directory is empty return error", func(t *testing.T) {
		expectedErr := "Output directory cannot be empty"

		pdfWriter := &PdfWriterMock{Error: errors.New(expectedErr)}
		writeErr := WritePDF(pdfWriter)

		if writeErr.Error() != expectedErr {
			t.Errorf("Expected %s but got %s", expectedErr, writeErr)
		}

	})

	t.Run("When PDF file is written then do not return error", func(t *testing.T) {
		pdfWriter := &PdfWriterMock{Error: nil}

		writeErr := WritePDF(pdfWriter)

		if writeErr != nil {
			t.Errorf("Expected no error but got %v", writeErr)
		}
	})
}
