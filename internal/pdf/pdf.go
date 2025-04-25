package pdf

import "codeberg.org/go-pdf/fpdf"

type PdfWriter interface {
	WriteNewPDF() error
}

type PdfWriterImpl struct{}

func NewPdfWriter() *PdfWriterImpl {
	return &PdfWriterImpl{}
}

func (p *PdfWriterImpl) WriteNewPDF() error {
	pdf := fpdf.New("", "", "", "")

	pdfErr := pdf.OutputFileAndClose("file.pdf")
	if pdfErr != nil {
		return pdfErr
	}

	return nil
}

func WriteNewPDF(pdfWriter PdfWriter) error {
	return pdfWriter.WriteNewPDF()
}
