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

// applyLBP applies the LBP operation using radius equal to 1
func ApplyLBP(img image.Image) ([]int64, error) {
	pixels := common.GetPixels(img)
	w, h := common.GetSize(img)

	var result []int64
	// Convert each pixel to grayscale
	for row := 1; row < w-1; row++ {
		for col := 1; col < h-1; col++ {

			threshold := pixels[row][col]

			binaryResult := ""
			for r := row - 1; r <= row+1; r++ {
				for c := col - 1; c <= col+1; c++ {
					if r != row || c != col {
						binaryResult += common.GetBinary(pixels[r][c], threshold)
					}
				}
			}

			i, err := strconv.ParseInt(binaryResult, 2, 32)
			if err != nil {
				return result, errors.New("Error normalizing the images")
			} else {
				result = append(result, i)
			}
		}
	}
	return result, nil
}
