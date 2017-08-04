package histogram

import (
	"errors"

	"github.com/kelvins/lbph/math"
	"github.com/kelvins/lbph/metrics"
)

// GetHistogram function generates a histogram based on the 'matrix' passed by parameter.
func GetHistogram(pixels [][]uint8, gridX, gridY uint8) ([]uint8, error) {
	var hist []uint8

	// Check the pixels 'matrix'
	if len(pixels) == 0 {
		return hist, errors.New("The pixels slice passed to the GetHistogram function is empty")
	}

	// Get the 'matrix' dimensions
	rows := len(pixels)
	cols := len(pixels[0])

	// Check the grid (X and Y)
	if gridX <= 0 || int(gridX) >= cols {
		return hist, errors.New("Invalid grid X passed to the GetHistogram function")
	}
	if gridY <= 0 || int(gridX) >= rows {
		return hist, errors.New("Invalid grid Y passed to the GetHistogram function")
	}

	// Get the size (width and height) of each region
	gridWidth := cols/int(gridX)
	gridHeight := rows/int(gridY)

	// Calculates the histogram of each grid
	for gX := 0; gX < int(gridX); gX++ {
		for gY := 0; gY < int(gridY); gY++ {
			// Create a slice with empty 256 positions
			regionHistogram := make([]uint8, 256)

			// Define the start and end positions for the following loop
			startPosX := gX*gridWidth
			startPosY := gY*gridHeight
			endPosX := (gX+1)*gridWidth
			endPosY := (gY+1)*gridHeight

			// Make sure that no pixel has been leave at the end
			if gX == int(gridX)-1 {
				endPosX = cols
			}
			if gY == int(gridY)-1 {
				endPosY = rows
			}

			// Creates the histogram for the current region
			for x := startPosX; x < endPosX; x++ {
				for y := startPosY; y < endPosY; y++ {
					// Make sure we are trying to access a valid position
					if x < len(pixels) {
						if y < len(pixels[x]) {
							if int(pixels[x][y]) < len(regionHistogram) {
								regionHistogram[pixels[x][y]] += 1
							}
						}
					}
				}
			}
			// Concatenate two slices
			hist = append(hist, regionHistogram...)
		}
	}

	return hist, nil
}

// GetHistogramDist function calculates the distance between two histograms.
// Histogram comparison:
// http://docs.opencv.org/2.4/doc/tutorials/imgproc/histograms/histogram_comparison/histogram_comparison.html
func CalcHistogramDist(hist1, hist2 []uint8, metric string) (float64, error) {
	// Check the histogram sizes
	if len(hist1) != len(hist2) {
		return 0, errors.New("Could not calculate the histogram distance. The slices have different sizes")
	}

	switch metric {
	case metrics.ChiSquare:
		return math.ChiSquare(hist1, hist2), nil
	case metrics.EuclideanDistance:
		return math.EuclideanDistance(hist1, hist2), nil
	case metrics.Intersection:
		return math.Intersection(hist1, hist2), nil
	case metrics.NormalizedIntersection:
		return math.NormalizedIntersection(hist1, hist2), nil
	}

	return 0, errors.New("Invalid metric selected to calculate the histogram distance")
}

