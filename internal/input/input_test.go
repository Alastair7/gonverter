package input

import "testing"

func TestGetUserInput(t *testing.T) {
	t.Run("When message is empty return error", func(t *testing.T) {
		message := ""
		expected := "Expected a system message but got an empty text"

		result, resultError := GetUserInput(message)

		if resultError == nil {
			t.Fatalf("Expected %s but got %v",
				expected, result)
		}

		if resultError.Error() != expected {
			t.Fatalf("Expected %s but got %s",
				expected, resultError.Error())
		}
	})

	t.Run("When input is invalid return error", func(t *testing.T) {
	})

	t.Run("When input is valid return user value", func(t *testing.T) {
	})
}
