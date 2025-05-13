package utils

func ScaleIfNecessary(originalWidth float64, originalHeight float64) (float64, float64) {
	maxWidth := PxToMm(2480)
	maxHeight := PxToMm(3508)

	if originalWidth <= maxWidth && originalHeight <= maxHeight {
		return originalWidth, originalHeight
	}

	widthFactor := float64(maxWidth) / originalWidth
	heightFactor := float64(maxHeight) / originalHeight

	scaleFactor := min(widthFactor, heightFactor)

	return originalWidth * scaleFactor, originalHeight * scaleFactor
}

func PxToMm(px float64) float64 {
	dpi := float64(300)
	return ((px / dpi) * 25.4)
}

func MmToPx(mm float64) float64 {
	dpi := float64(300)

	return (mm / 25.4) * dpi
}
