package file

import (
	"bufio"
	"fmt"
	"net/http"

	"codeberg.org/go-pdf/fpdf"
	"github.com/Alastair7/gonverter/internal/document"
	"github.com/Alastair7/gonverter/internal/utils"
)

type FileHandler struct{}

func (*FileHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	parsingErr := req.ParseMultipartForm(5 << 20)
	if parsingErr != nil {
		panic(parsingErr)
	}

	file, header, formErr := req.FormFile("file")
	if formErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error receiving the file %v", formErr)
		return
	}

	defer file.Close()

	isValid := utils.ValidateFile(header.Filename)

	if !isValid {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "The file extension is not valid")
	}

	// TODO: PDF Logic: Re-escalate image and add it to a PDF
	// Get image width and height and convert them to mm
	width, height := utils.ScaleIfNecessary(0, 0)

	pdfClient := document.NewPdfDocument()
	pdfClient.AddPage()

	fileReader := bufio.NewReader(file)

	// Create a image string with the image stream received from the user
	pdfClient.RegisterImageOptionsReader(header.Filename, fpdf.ImageOptions{}, fileReader)
	pdfClient.ImageOptions("avatar.png", 0, 0,
		width, height, false,
		fpdf.ImageOptions{ImageType: "png", ReadDpi: true, AllowNegativePosition: false}, 0, "")

	pdfErr := pdfClient.OutputFileAndClose("image.pdf")
	if pdfErr != nil {
		panic(pdfErr)
	}
}

func NewFileHandler() *FileHandler {
	return &FileHandler{}
}
