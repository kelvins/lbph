package common

import (
	"errors"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

// LoadImage function is used to provide an easy way to load an image file
func LoadImage(filePath string) (image.Image, error) {
	// Open the file image
	fImage, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	// Ensure that the image file will be closed
	defer fImage.Close()

	// Convert it to an image "object"
	img, _, err := image.Decode(fImage)
	if err != nil {
		return nil, err
	}

	return img, nil
}

// GetSize function is responsible for get the width and height from an image
func GetSize(img image.Image) (int, int) {
	if img == nil {
		return 0, 0
	}
	bounds := img.Bounds()
	return bounds.Max.X, bounds.Max.Y
}

// CheckInputData function is responsible for check if all images have the same size
func CheckInputData(images []image.Image) error {
	// Check if the slice is empty
	if len(images) == 0 {
		return errors.New("Empty slice")
	}
	// Check if the first image is nil
	if images[0] == nil {
		return errors.New("One or more images are nil")
	}
	// Get the image size from the first image
	width, height := GetSize(images[0])

	// Check if the size is valid
	if width <= 0 && height <= 0 {
		return errors.New("Invalid image sizes")
	}

	// Verifies each image
	for index := 0; index < len(images); index++ {
		// Check if the current image is nil
		if images[index] == nil {
			return errors.New("One or more images are nil")
		}

		// Get the size from the current image
		w, h := GetSize(images[index])

		// Check if all images have the same size
		if w != width || h != height {
			return errors.New("One or more images have different sizes")
		}
	}
	return nil
}

// GetBinary function return 1 (string) if the value is equal or higher than the
// threshold or 0 (string) otherwise
func GetBinary(value, threshold uint8) string {
	if value >= threshold {
		return "1"
	} else {
		return "0"
	}
}

// GetPixels function returns a 'matrix' containing all pixels from the image passed by parameter
func GetPixels(img image.Image) [][]uint8 {
	var pixels [][]uint8
	
	// Check if the image is nil
	if img == nil {
		return pixels
	}

	// Get the image size
	w, h := GetSize(img)

	for x := 0; x < w; x++ {
		var row []uint8
		for y := 0; y < h; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			// Convert the RGB to Grayscale (red*30% + green*59% + blue*11%)
			pixel := (float32(r)*0.299)+(float32(g)*0.587)+(float32(b)*0.114)
			// Convert to uint8 (0-255) and append to the slice
			row = append(row, uint8(pixel))
		}
		// Append the slice to the pixels
		pixels = append(pixels, row)
	}
	return pixels
}
