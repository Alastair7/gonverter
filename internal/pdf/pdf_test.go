package pdf

import (
	"errors"
	"testing"
)

type PdfWriterMock struct {
	Error error
}

func (p PdfWriterMock) WriteNewPDF() error {
	return p.Error
}

func TestWriteNewPDF(t *testing.T) {
	t.Run("When PDF file is written then return error", func(t *testing.T) {
		expectedErr := "Error while creating the PDF"
		pdfWriter := &PdfWriterMock{Error: errors.New(expectedErr)}

		writeErr := WriteNewPDF(pdfWriter)

		if writeErr == nil {
			t.Errorf("Expected %s but pdf writing was success", expectedErr)
		}

	})

	t.Run("When PDF file is written then do not return error", func(t *testing.T) {
		pdfWriter := &PdfWriterMock{Error: nil}

		writeErr := WriteNewPDF(pdfWriter)

		if writeErr != nil {
			t.Errorf("Expected no error but got %v", writeErr)
		}

	})
}
