package file

import (
	"fmt"
	"image"
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

	if req.MultipartForm != nil {
		defer func() {
			_ = req.MultipartForm.RemoveAll()
		}()
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

	img, _, decodeErr := image.Decode(file)
	if decodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error decoding the image")
	}

	originalWidth, originalHeight := utils.GetImageDimensionsInMm(img)
	width, height := utils.ScaleIfNecessary(originalWidth, originalHeight)
	imgType := utils.GetImageType(header.Filename)

	pdfClient := document.NewPdfDocument()
	pdfClient.AddPage()

	opts := fpdf.ImageOptions{
		ImageType:             imgType,
		ReadDpi:               true,
		AllowNegativePosition: false,
	}

	pdfClient.RegisterImageOptionsReader(header.Filename, opts, file)

	pdfClient.ImageOptions(header.Filename, 0, 0,
		width, height, false,
		opts, 0, "")

	outputName := utils.GetImageNameWithoutExtension(header.Filename)
	pdfErr := pdfClient.OutputFileAndClose(outputName + ".pdf")
	if pdfErr != nil {
		panic(pdfErr)
	}
}

func NewFileHandler() *FileHandler {
	return &FileHandler{}
}
