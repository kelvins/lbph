package lbp

import (
	"errors"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"strconv"

	"github.com/kelvins/lbph/common"
)

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
	pixels := common.GetPixels(img)

	// Get the image size (width and height)
	width, height := common.GetSize(img)

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
						binaryResult += common.GetBinaryString(pixels[tempX][tempY], threshold)
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
