package file

import (
	"os"
	"testing"
)

type OsMock struct {
	FileInfo os.FileInfo
	Error    error
}

func (o *OsMock) Stat(name string) (os.FileInfo, error) {
	return o.FileInfo, o.Error
}

func TestGetImageDimensions(t *testing.T) {
	t.Run("When image not found return error", func(t *testing.T) {
		expectedErr := "Image was not found"

		_, statErr := os.Stat("")

		if statErr.Error() != expectedErr {
			t.Errorf("Expected %s got %s", expectedErr, statErr)
		}
	})

	t.Run("When image type not valid return error", func(t *testing.T) {})

	t.Run("When image dimensions are obtained return them", func(t *testing.T) {})
}
