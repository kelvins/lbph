// Performs face recognition using lbph
package lbph

import (
	"errors"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"math"
	"strconv"

  "github.com/kelvins/lbph/structs"
)

var Data structs.Data

// getSize is responsible for get the width and height from the image
func getSize(img image.Image) (int, int) {
	bounds := img.Bounds()
	return bounds.Max.X, bounds.Max.Y
}

// isGrayscale function is responsible for check if an image is in grayscale.
func isGrayscale(img image.Image) bool {
	// Gets the width and height of the image
	w, h := getSize(img)

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
func checkInputData(images []image.Image) error {
	width, height := getSize(images[0])

	for index := 0; index < len(images); index++ {
		// Check if the image is in grayscale
		if !isGrayscale(images[index]) {
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
func getBinary(value, threshold uint8) string {
	if value >= threshold {
		return "1"
	} else {
		return "0"
	}
}

// Return a 'matrix' containing all pixels from the image passed by parameter
func getPixels(img image.Image) [][]uint8 {
	w, h := getSize(img)

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

// applyLBP applies the LBP operation using radius equal to 1
func applyLBP(img image.Image) ([]int64, error) {
	pixels := getPixels(img)
	w, h := getSize(img)

	var result []int64
	// Convert each pixel to grayscale
	for row := 1; row < w-1; row++ {
		for col := 1; col < h-1; col++ {

			threshold := pixels[row][col]

			binaryResult := ""
			for r := row - 1; r <= row+1; r++ {
				for c := col - 1; c <= col+1; c++ {
					if r != row || c != col {
						binaryResult += getBinary(pixels[r][c], threshold)
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

// getHistogram generates a histogram based on the LBP result
func getHistogram(img image.Image) ([256]int64, error) {
	var histogram [256]int64
	lbp, err := applyLBP(img)
	if err != nil {
		return histogram, errors.New("Error in the LBP operation")
	}
	for index := 0; index < len(lbp); index++ {
		histogram[lbp[index]] += 1
	}
	return histogram, nil
}

// getHistogramDist calculates the distance between two histograms using euclidean distance
// sum = sqrt((h1(i)-h2(i))^2)
func getHistogramDist(hist1, hist2 [256]int64) float64 {
	var sum float64
	for index := 0; index < len(hist1); index++ {
		sum += float64((hist1[index] - hist2[index]) * (hist1[index] - hist2[index]))
	}
	return math.Sqrt(sum)
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

	var histograms [][256]int64
	for index := 0; index < len(images); index++ {
		hist, err := getHistogram(images[index])
		if err != nil {
			return err
		}
		histograms = append(histograms, hist)
	}

	// This conditional must never occurs
	if len(histograms) == 0 {
		return errors.New("None histogram was calculated")
	}

	Data = structs.Data{
	  Images:     images,
		Labels:     labels,
		Histograms: histograms,
	}

	return nil
}

func Predict(img image.Image) (string, float64, error) {
	hist, err := getHistogram(img)
	if err != nil {
		return "", 0.0, errors.New("Could not get the image histogram")
	}
	var min float64
	var i int
	for index := 0; index < len(Data.Histograms); index++ {
		if index == 0 {
			i = index
			min = getHistogramDist(hist, Data.Histograms[index])
		} else {
			x := getHistogramDist(hist, Data.Histograms[index])
			if x < min {
				min = x
				i = index
			}
		}
	}
	return Data.Labels[i], min, nil
}
