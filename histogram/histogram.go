package histogram

import (
	"errors"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"math"

	"github.com/kelvins/lbph/lbp"
)

// GetHistogram function generates a histogram based on the LBP result
func GetHistogram(img image.Image) ([256]int64, error) {
	// The histogram size is fixed (256)
	var histogram [256]int64

	// Calculate the LBP operation
	lbp, err := lbp.ApplyLBP(img)

	// Check for errors
	if err != nil {
		return histogram, errors.New("Error in the LBP operation")
	}

	// Creates the histogram by adding each lbp result in the histogram correct position
	for index := 0; index < len(lbp); index++ {
		histogram[lbp[index]] += 1
	}

	return histogram, nil
}

// GetHistogramDist function calculates the distance between two histograms
// using euclidean distance: sum = sqrt((h1(i)-h2(i))^2)
func GetHistogramDist(hist1, hist2 [256]int64) float64 {
	var sum float64
	for index := 0; index < len(hist1); index++ {
		sum += float64((hist1[index] - hist2[index]) * (hist1[index] - hist2[index]))
	}
	return math.Sqrt(sum)
}
