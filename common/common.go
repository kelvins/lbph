package common

import (
	"errors"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

// getSize is responsible for get the width and height from the image
func GetSize(img image.Image) (int, int) {
	bounds := img.Bounds()
	return bounds.Max.X, bounds.Max.Y
}

// isGrayscale function is responsible for check if an image is in grayscale.
func IsGrayscale(img image.Image) bool {
	// Gets the width and height of the image
	w, h := GetSize(img)

	if w == 0 || h == 0 {
		return false
	}

	// Verifies each pixel (R,G,B)
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			if r != g && g != b {
				return false
			}
		}
	}

	return true
}

// checkInputData function is responsible for check if all images are in
// grayscale and have the same size (width and height)
func CheckInputData(images []image.Image) error {
	width, height := GetSize(images[0])

	for index := 0; index < len(images); index++ {
		// Check if the image is in grayscale
		if !IsGrayscale(images[index]) {
			return errors.New("One or more images are not in grayscale")
		}

		b := images[index].Bounds()

		// Check if all images have the same size
		if b.Max.X != width || b.Max.Y != height {
			return errors.New("One or more images have different sizes")
		}
	}
	return nil
}

// getBinary function return 1 (string) if the value is equal or higher than the
// threshold or 0 (string) otherwise
func GetBinary(value, threshold uint8) string {
	if value >= threshold {
		return "1"
	} else {
		return "0"
	}
}

// Return a 'matrix' containing all pixels from the image passed by parameter
func GetPixels(img image.Image) [][]uint8 {
	w, h := GetSize(img)

	var pixels [][]uint8
	for row := 0; row < w; row++ {
		var r []uint8
		for col := 0; col < h; col++ {
			red, _, _, _ := img.At(row, col).RGBA()
			r = append(r, uint8(red))
		}
		pixels = append(pixels, r)
	}
	return pixels
}
