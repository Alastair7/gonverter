package pdf

import (
	"errors"

	"codeberg.org/go-pdf/fpdf"
)

type PdfWriter interface {
	Write() error
}

type PdfCreator interface {
	NewPDF()
}

type PdfManager struct {
	*fpdf.Fpdf
	Orientation string
	Unit        string
	Size        string
	FontDir     string
	OutputDir   string
}

func NewPdfManager(orientation string, unit string, size string, fontDir string, outputDir string) *PdfManager {
	return &PdfManager{
		Orientation: orientation,
		Unit:        unit,
		Size:        size,
		FontDir:     fontDir,
		OutputDir:   outputDir,
	}
}

func (p *PdfManager) NewPDF() {
	pdf := fpdf.New(p.Orientation,
		p.Unit,
		p.Size,
		p.FontDir)

	p.Fpdf = pdf
}

func (p *PdfManager) Write() error {
	pdfManager := p.Fpdf

	if p.OutputDir == "" {
		return errors.New("Output directory cannot be empty")
	}

	pdfManager.AddPage()
	/*
		Scale Factor = dimension of new shape / dimension of original shape
		Scale Up = Larger dimension / Smaller dimension
		Scale Down = Smaller dimension / Larger dimension
	*/
	pdfManager.ImageOptions("image2.png", 0, 0, 300, 170, false, fpdf.ImageOptions{ImageType: "png", ReadDpi: true, AllowNegativePosition: false}, 0, "")

	pdfErr := pdfManager.OutputFileAndClose(p.OutputDir)
	if pdfErr != nil {
		return pdfErr
	}

	return nil
}

func WritePDF(pdfWriter PdfWriter) error {

	return pdfWriter.Write()
}
