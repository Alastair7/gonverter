package main

import (
	"os"

	"github.com/Alastair7/gonverter/internal/input"
)

func main() {
	println("=== GONVERTER ===")
	println("Gonverter converts images or docs to PDF.\n")

	userInput, inputError := input.GetUserInput("Add path to image/document", os.Stdin)
	if inputError != nil {
		println("Error obtaining user input: ", inputError.Error())
	}

	println("User Input: ", userInput)
}
