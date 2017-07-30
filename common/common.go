package common

import (
	"errors"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

// LoadImage function is used to provide an easy way to load an image file.
func LoadImage(filePath string) (image.Image, error) {
	// Open the image file
	fImage, err := os.Open(filePath)
	// Check if no error has occurred
	if err != nil {
		return nil, err
	}

	// Ensure that the image file will be closed
	defer fImage.Close()

	// Decode it to an image "object" (we don't need the format name so we use "_")
	img, _, err := image.Decode(fImage)
	// Check if no error has occurred
	if err != nil {
		return nil, err
	}

	return img, nil
}

// GetSize function is used to get the width and height from an image.
// If the image is nil it will return 0 width and 0 height
func GetSize(img image.Image) (int, int) {
	if img == nil {
		return 0, 0
	}
	// Get the image bounds
	bounds := img.Bounds()
	// Return the width and height
	return bounds.Max.X, bounds.Max.Y
}

// CheckImagesSizes function is used to check if all images have the same size.
func CheckImagesSizes(images []image.Image) error {
	// Check if the slice is empty
	if len(images) == 0 {
		return errors.New("The slice has no images")
	}
	// Check if the first image is nil
	if images[0] == nil {
		return errors.New("At least one image is nil")
	}

	// Get the image size from the first image
	defaultWidth, defaultHeight := GetSize(images[0])

	// Check if the size is valid
	if defaultWidth <= 0 || defaultHeight <= 0 {
		return errors.New("Invalid image size")
	}

	// Check each image in the slice
	for index := 0; index < len(images); index++ {
		// Check if the current image is nil
		if images[index] == nil {
			return errors.New("At least one image is nil")
		}

		// Get the size from the current image
		width, height := GetSize(images[index])

		// Check if all images have the same size
		if width != defaultWidth || height != defaultHeight {
			return errors.New("One or more images have different sizes")
		}
	}
	// No error has occurred, return nil
	return nil
}

// GetBinaryString function used to get a binary value as a string based on a threshold.
// Return "1" if the value is equal or higher than the threshold or "0" otherwise.
func GetBinaryString(value, threshold uint8) string {
	if value >= threshold {
		return "1"
	} else {
		return "0"
	}
}

// GetPixels function returns a 'matrix' ([][]uint8) containing all pixels from the image passed by parameter.
func GetPixels(img image.Image) [][]uint8 {
	var pixels [][]uint8
	
	// Check if the image is nil
	if img == nil {
		return pixels
	}

	// Get the image size
	width, height := GetSize(img)

	// For each pixel in the image (x, y) convert it to grayscale and store it in the 'matrix'
	for x := 0; x < width; x++ {
		var row []uint8
		for y := 0; y < height; y++ {
			// Get the RGB from the current pixel
			r, g, b, _ := img.At(x, y).RGBA()

			// Convert the RGB to Grayscale (red*30% + green*59% + blue*11%)
			// https://en.wikipedia.org/wiki/Grayscale#Luma_coding_in_video_systems
			pixel := (float32(r) * 0.299) + (float32(g) * 0.587) + (float32(b) * 0.114)

			// Convert the pixel from uin64 to uint8 (0-255) and append it to the slice
			row = append(row, uint8(pixel))
		}
		// Append the row (slice) to the pixels 'matrix'
		pixels = append(pixels, row)
	}

	// Return all pixels
	return pixels
}
