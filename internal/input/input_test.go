package input

import (
	"strings"
	"testing"
)

func TestGetUserInput(t *testing.T) {
	t.Run("When message is empty return error", func(t *testing.T) {
		message := ""
		expected := "Expected a system message but got an empty text"

		reader := strings.NewReader("test")
		result, resultError := GetUserInput(message, reader)

		if resultError == nil {
			t.Fatalf("Expected %s but got %v",
				expected, result)
		}

		if resultError.Error() != expected {
			t.Fatalf("Expected %s but got %s",
				expected, resultError.Error())
		}
	})

	t.Run("When input is valid return user value", func(t *testing.T) {
		message := "Insert a value:"
		expected := "1"

		r := strings.NewReader(expected)
		result, resultError := GetUserInput(message, r)

		if resultError != nil {
			t.Fatalf("Expected %s but got %s", result, resultError.Error())
		}

		if result != expected {
			t.Fatalf("Expected %s but got %s", expected, result)
		}
	})
}
