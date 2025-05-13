package document

import "codeberg.org/go-pdf/fpdf"

func NewPdfDocument() *fpdf.Fpdf {
	pdfClient := fpdf.New("", "mm", "A4", "")

	return pdfClient
}
