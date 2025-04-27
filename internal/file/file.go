package file

import (
	"errors"
	"image"
	"io"
)

type Imager interface {
	GetDimensions(reader io.Reader) (w, h int, error error)
}

type ImageManager struct {
	reader io.Reader
}

func (i *ImageManager) ScaleDownImage(scaleFactor float64) (width, height int, error error) {
	if scaleFactor >= 1 {
		return 0, 0, errors.New("Scale Factor must be less than 1 to scale down")
	}

	originalWidth, originalHeight, imgError := i.getDimensions()

	if imgError != nil {
		return 0, 0, imgError
	}

	resizedWidth := int(float64(originalWidth) * scaleFactor)
	resizedHeight := int(float64(originalHeight) * scaleFactor)

	return resizedWidth, resizedHeight, nil
}

func (i *ImageManager) getDimensions() (w, h int, error error) {
	imageInfo, _, decodingErr := image.DecodeConfig(i.reader)
	if decodingErr != nil {
		return 0, 0, decodingErr
	}

	return imageInfo.Width, imageInfo.Height, nil
}

func NewImageManager(reader io.Reader) *ImageManager {
	return &ImageManager{reader: reader}
}
