package utils

import (
	"image"
	"strings"
)

func GetImageDimensionsInMm(img image.Image) (float64, float64) {
	widthPx := img.Bounds().Dx()
	heightPx := img.Bounds().Dy()

	widthMm := PxToMm((float64(widthPx)))
	heightMm := PxToMm((float64(heightPx)))

	return widthMm, heightMm
}

func GetImageType(name string) string {
	extension := strings.Split(name, ".")[1]

	return extension
}

func GetImageNameWithoutExtension(name string) string {
	return strings.Split(name, ".")[0]
}
