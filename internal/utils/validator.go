package utils

import (
	"path/filepath"
	"slices"
)

func ValidateFile(filename string) bool {
	validExtensions := []string{".jpg", ".png", ".jpeg"}
	extension := filepath.Ext(filename)

	return slices.Contains(validExtensions, extension)
}
