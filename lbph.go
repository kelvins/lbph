// Performs face recognition using lbph
package lbph

import (
	"errors"
	"image"
  "fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"strconv"

	"github.com/kelvins/imgproc"
)

var imgs []image.Image
var lbls []string

var histograms [][256]int64

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

func getBinary(value, threshold uint8) string {
	if value >= threshold {
		return "1"
	} else {
		return "0"
	}
}

func getPixels(img image.Image) [][]uint8 {
	bounds := img.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y

	var pixels [][]uint8
	for row := 0; row < w; row++ {
	   var r []uint8
		for col := 0; col < h; col++ {
      red, _, _, _ := img.At(row, col).RGBA()
			r = append(r, uint8(red))
			//r = append(r, newImage.GrayAt(row, col).Y)
		}
		pixels = append(pixels, r)
	}
	return pixels
}

func getHistogram(img image.Image) ([256]int64, error) {
	pixels := getPixels(img)
	bounds := img.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y

	var histogram [256]int64
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
				return histogram, errors.New("Error normalizing the images")
			} else {
				histogram[i] += 1
			}
		}
	}
	return histogram, nil
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
	imgs = images
	lbls = labels
	return nil
}

func sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 1000; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func getHistogramDist(hist1, hist2 [256]int64) float64 {
	var sum float64
	for index := 0; index < len(hist1); index++ {
		sum += float64((hist1[index] - hist2[index]) * (hist1[index] - hist2[index]))
	}
	return sqrt(sum)
}

func Predict(img image.Image) (string, float64, error) {
	hist, err := getHistogram(img)
	if err != nil {
		return "", 0.0, errors.New("Could not get the image histogram")
	}
	var min float64
	var i int
	for index := 0; index < len(histograms); index++ {
		if index == 0 {
			i = index
			min = getHistogramDist(hist, histograms[index])
		} else {
			x := getHistogramDist(hist, histograms[index])
			if x < min {
				min = x
				i = index
			}
		}
	}
	return lbls[i], min, nil
}
