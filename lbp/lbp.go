package lbp

import (
	"errors"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"strconv"
)

// getBinaryString function used to get a binary value as a string based on a threshold.
// Return "1" if the value is equal or higher than the threshold or "0" otherwise.
func getBinaryString(value, threshold uint8) string {
	if value >= threshold {
		return "1"
	} else {
		return "0"
	}
}

// GetImageSize function is used to get the width and height from an image.
// If the image is nil it will return 0 width and 0 height
func GetImageSize(img image.Image) (int, int) {
	if img == nil {
		return 0, 0
	}
	// Get the image bounds
	bounds := img.Bounds()
	// Return the width and height
	return bounds.Max.X, bounds.Max.Y
}

// GetPixels function returns a 'matrix' ([][]uint8) containing all pixels from the image passed by parameter.
func GetPixels(img image.Image) [][]uint8 {
	var pixels [][]uint8

	// Check if the image is nil
	if img == nil {
		return pixels
	}

	// Get the image size
	width, height := GetImageSize(img)

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

// ApplyLBP applies the LBP operation based on the radius and neighbors passed by parameter
// The radius and neighbors parameters are not in use
func ApplyLBP(img image.Image, radius, neighbors uint8) ([][]uint8, error) {

	var lbpPixels [][]uint8
	// Check the parameters
	if img == nil {
		return lbpPixels, errors.New("The image passed to the ApplyLBP function is nil")
	}
	if radius <= 0 {
		return lbpPixels, errors.New("Invalid radius parameter passed to the ApplyLBP function")
	}
	if neighbors <= 0 {
		return lbpPixels, errors.New("Invalid neighbors parameter passed to the ApplyLBP function")
	}

	// Get the pixels 'matrix' ([][]uint8)
	pixels := GetPixels(img)

	// Get the image size (width and height)
	width, height := GetImageSize(img)

	// For each pixel in the image
	for x := 1; x < width-1; x++ {
		var currentRow []uint8
		for y := 1; y < height-1; y++ {

			// Get the current pixel as the threshold
			threshold := pixels[x][y]

			binaryResult := ""
			// Window based on the radius (3x3)
			for tempX := x - 1; tempX <= x+1; tempX++ {
				for tempY := y - 1; tempY <= y+1; tempY++ {
					// Get the binary for all pixels around the threshold
					if tempX != x || tempY != y {
						binaryResult += getBinaryString(pixels[tempX][tempY], threshold)
					}
				}
			}

			// Convert the binary string to a decimal integer
			dec, err := strconv.ParseUint(binaryResult, 2, 8)
			if err != nil {
				return lbpPixels, errors.New("Error converting binary to uint in the ApplyLBP function")
			} else {
				// Append the decimal do the result slice
				// ParseUint returns a uint64 so we need to convert it to uint8
				currentRow = append(currentRow, uint8(dec))
			}
		}
		// Append the slice to the 'matrix'
		lbpPixels = append(lbpPixels, currentRow)
	}
	return lbpPixels, nil
}
