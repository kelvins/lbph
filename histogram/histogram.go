package histogram

import (
	"errors"

	"github.com/kelvins/lbph/math"
	"github.com/kelvins/lbph/metric"
)

// Calculate function generates a histogram based on the 'matrix' passed by parameter.
func Calculate(pixels [][]uint8, gridX, gridY uint8) ([]float64, error) {
	var hist []float64

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
			regionHistogram := make([]float64, 256)

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

// Compare function is used to compare two histograms using a selected metric.
// Histogram comparison references:
// http://docs.opencv.org/2.4/doc/tutorials/imgproc/histograms/histogram_comparison/histogram_comparison.html
func Compare(hist1, hist2 []float64, selectedMetric string) (float64, error) {

	switch selectedMetric {
	case metric.ChiSquare:
		return math.ChiSquare(hist1, hist2)
	case metric.EuclideanDistance:
		return math.EuclideanDistance(hist1, hist2)
	case metric.NormalizedEuclideanDistance:
		return math.NormalizedEuclideanDistance(hist1, hist2)
	case metric.Intersection:
		return math.Intersection(hist1, hist2)
	case metric.NormalizedIntersection:
		return math.NormalizedIntersection(hist1, hist2)
	}

	return 0, errors.New("Invalid metric selected to compare the histograms")
}

