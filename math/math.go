package math

import (
	"errors"
	"math"
)

// checkHistograms check if the histograms are correct.
func checkHistograms(hist1, hist2 []float64) error {
	if len(hist1) == 0 || len(hist2) == 0 {
		return errors.New("Could not compare the histograms. The histogram is empty.")
	}
	if len(hist1) != len(hist2) {
		return errors.New("Could not compare the histograms. The slices have different sizes.")
	}
	return nil
}

// chiSquare calculates the distance between two histograms using
// the chi square statistic.
// x^2 = \sum_{i=1}^{n}\frac{(hist1_{i} - hist2_{i})^2}{hist1_{i}}
// References:
// http://file.scirp.org/Html/8-72278_30995.htm
// https://www.google.com/patents/WO2007080817A1?cl=en
func ChiSquare(hist1, hist2 []float64) (float64, error) {

	// Check the histogram sizes
	if err := checkHistograms(hist1, hist2); err != nil {
		return 0.0, err
	}

	var sum float64
	for index := 0; index < len(hist1); index++ {
		numerator := math.Pow(hist1[index]-hist2[index], 2)
		denominator := hist1[index]
		sum += numerator / denominator
	}
	return sum, nil
}

// EuclideanDistance calculates the euclidean distance between two histograms
// by the following formula:
// D = \sqrt{\sum_{i=1}^{n}(hist1_{i} - hist2_{i})^2}
// Reference: http://www.pbarrett.net/techpapers/euclid.pdf
func EuclideanDistance(hist1, hist2 []float64) (float64, error) {

	// Check the histogram sizes
	if err := checkHistograms(hist1, hist2); err != nil {
		return 0.0, err
	}

	var sum float64
	for index := 0; index < len(hist1); index++ {
		sum += math.Pow(hist1[index]-hist2[index], 2)
	}
	return math.Sqrt(sum), nil
}

// NormalizedEuclideanDistance calculates the euclidean distance normalized.
// D = \sqrt{\sum_{i=1}^{n} \frac{(hist1_{i} - hist2_{i})^2}{n}}
// Reference:
// http://www.pbarrett.net/techpapers/euclid.pdf
func NormalizedEuclideanDistance(hist1, hist2 []float64) (float64, error) {

	// Check the histogram sizes
	if err := checkHistograms(hist1, hist2); err != nil {
		return 0.0, err
	}

	var sum float64
	n := float64(len(hist1))
	for index := 0; index < len(hist1); index++ {
		sum += math.Pow(hist1[index]-hist2[index], 2) / n
	}
	return math.Sqrt(sum), nil
}

// AbsoluteValue calculates the absolute values between two histograms.
// D = \sum_{i=1}^{n} \left | hist1_{i} - hist2_{i} \right |
func AbsoluteValue(hist1, hist2 []float64) (float64, error) {

	// Check the histogram sizes
	if err := checkHistograms(hist1, hist2); err != nil {
		return 0.0, err
	}

	var sum float64
	for index := 0; index < len(hist1); index++ {
		sum += math.Abs(hist1[index] - hist2[index])
	}
	return sum, nil
}
