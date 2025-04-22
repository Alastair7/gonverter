package input

import (
	"errors"
	"fmt"
	"io"
)

func GetUserInput(message string, r io.Reader) (string, error) {
	if message == "" {
		return "", errors.New("Expected a system message but got an empty text")
	}

	print(message, ": ")

	var userInput string
	_, scanError := fmt.Fscanln(r, &userInput)

	if scanError != nil {
		return "", scanError
	}

	return userInput, nil
}
