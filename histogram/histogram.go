package histogram

import (
	//"errors"
	"math"
)

// GetHistogram function generates a histogram based on the LBP result
func GetHistogram(pixels [][]uint8, gridX, gridY uint8) ([]uint8, error) {
	var hist []uint8

	// Creates the histogram by adding each lbp result in the histogram correct position
	for x := 0; x < len(pixels); x++ {
		for y := 0; y < len(pixels[x]); y++ {
			//hist[lbp[row][col]] += 1
			hist = append(hist, 1)
			// HERE WE NEED TO CREATE THE MULTIPLE HISTOGRAMS AND CONCATENATE IT
		}
	}

	return hist, nil
}

// GetHistogramDist function calculates the distance between two histograms
// using euclidean distance: sum = sqrt((h1(i)-h2(i))^2)
func CalcHistogramDist(hist1, hist2 []uint8) float64 {
	var sum float64
	for index := 0; index < len(hist1); index++ {
		sum += float64((hist1[index] - hist2[index]) * (hist1[index] - hist2[index]))
	}
	return math.Sqrt(sum)
}
