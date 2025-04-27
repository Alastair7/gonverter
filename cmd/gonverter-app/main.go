package main

import (
	"os"

	"github.com/Alastair7/gonverter/internal/input"
	"github.com/Alastair7/gonverter/internal/pdf"
)

func main() {
	println("=== GONVERTER ===")
	println("Gonverter converts images or docs to PDF.\n")

	userInput, inputError := input.GetUserInput("Add path to image/document", os.Stdin)
	if inputError != nil {
		println("Error obtaining user input: ", inputError.Error())
	}

	println(userInput)

	pdfManager := pdf.NewPdfManager("P", "mm", "A4", "", "file.pdf")

	pdfManager.NewPDF()
	writeErr := pdfManager.Write()
	if writeErr != nil {
		println("Error writing pdf file: ", writeErr.Error())
	}
}
