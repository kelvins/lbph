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

// ApplyLBP applies the LBP operation using radius equal to 1
// We need to implement a way to apply the LBP based on a different radius passed by parameter
func ApplyLBP(img image.Image) ([]int64, error) {
	// Get the pixels 'matrix' ([][]uint8)
	pixels := common.GetPixels(img)

	// Get the image size (width and height)
	w, h := common.GetSize(img)

	var result []int64
	// For each pixel in the image
	for row := 1; row < w-1; row++ {
		for col := 1; col < h-1; col++ {

			// Get the current pixel as the threshold
			threshold := pixels[row][col]

			binaryResult := ""
			// Image sample 3x3
			for r := row - 1; r <= row+1; r++ {
				for c := col - 1; c <= col+1; c++ {
					// Get the binary for all pixels around the threshold
					if r != row || c != col {
						binaryResult += common.GetBinary(pixels[r][c], threshold)
					}
				}
			}

			// Convert the binary string to a decimal integer
			i, err := strconv.ParseInt(binaryResult, 2, 32)
			if err != nil {
				return result, errors.New("Error normalizing the images")
			} else {
				// Append the decimal do the result slice
				result = append(result, i)
			}
		}
	}
	return result, nil
}
