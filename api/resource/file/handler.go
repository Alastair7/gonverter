package file

import (
	"fmt"
	"net/http"

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
}

func NewFileHandler() *FileHandler {
	return &FileHandler{}
}
