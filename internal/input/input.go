package input

import "errors"

func GetUserInput(message string) (any, error) {
	if message == "" {
		return nil, errors.New("Expected a system message but got an empty text")
	}
	return "", nil
}
