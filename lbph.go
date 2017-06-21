// Performs face recognition using lbph
package lbph

import (
	"errors"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"strconv"

	"github.com/kelvins/imgproc"
)

var histograms [][]int16

func checkInputData(images []image.Image) error {
	// Get the image bounds
	bounds := images[0].Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	for index := 0; index < len(images); index++ {
		// Check if the image is in grayscale
		if !imgproc.IsGrayscale(images[index]) {
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

func getBinary(value, threshold int) string {
	if value >= threshold {
		return "1"
	} else {
		return "0"
	}
}

// Function used to train the algorithm
func Train(images []image.Image, labels []string) error {

	if len(images) != len(labels) {
		return errors.New("Slices have different sizes")
	}
	if len(images) == 0 {
		return errors.New("Empty vector")
	}

	// Check if the input data is in the correct format
	err := checkInputData(images)
	if err != nil {
		return err
	}

	// STEP 2 - Convert the image
	for index := 0; index < len(images); index++ {
		bounds := images[index].Bounds()
		w, h := bounds.Max.X, bounds.Max.Y
		image := image.NewGray(bounds)

		// Convert each pixel to grayscale
		for x := 0; x < w; x++ {
			for y := 0; y < h; y++ {
				image.Set(x, y, images[index].At(x, y))
			}
		}

		var histogram []int16
		// Convert each pixel to grayscale
		for row := 1; row < w-1; row++ {
			for col := 1; col < h-1; col++ {

				threshold := image.GrayAt(row, col).Y

				binaryResult := ""
				for r := row - 1; r <= row+1; r++ {
					for c := col - 1; c <= col+1; c++ {
						if r != row || c != col {
							binaryResult += getBinary(image.GrayAt(r, c).Y, threshold)
						}
					}
				}

				i, err := strconv.ParseInt(binaryResult, 10, 32)
				if err != nil {
					return errors.New("Error normalizing the images")
				} else {
					fmt.Println(i)
					histogram = append(histogram, i)
				}
			}
		}

		histograms = append(histograms, histogram)
	}

	// This conditional must never occurs
	if len(histograms) == 0 {
		return errors.New("None histogram was calculated")
	}

	return nil
}

func Predict(image image.Image) (string, float64, error) {
	return "", 0.0, nil
}
