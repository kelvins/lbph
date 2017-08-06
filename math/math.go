package math

import (
	"errors"
	"math"
)

// min function returns the minimum value.
func min(value1, value2 float64) float64 {
	if value1 < value2 {
		return value1
	} else {
		return value2
	}
}

// sum function returns the sum of a float64 slice.
func sum(slice []float64) float64 {
	var sum float64
	for _, value := range slice {
		sum += value
	}
	return sum
}

// abs function returns the absolute value.
func abs(a float64) float64 {
	if a < 0 {
		return a * -1.0
	} else {
		return a
	}
}

// chiSquare calculates the distance between two histograms using
// the chi square statistic.
// x^2 = \sum_{i=1}^{n}\frac{(hist1_{i} - hist2_{i})^2}{hist1_{i}}
// References:
// http://file.scirp.org/Html/8-72278_30995.htm
// https://www.google.com/patents/WO2007080817A1?cl=en
func ChiSquare(hist1, hist2 []float64) (float64, error) {
	// Check the histogram sizes
	if len(hist1) != len(hist2) {
		return 0, errors.New("Could not compare the histograms. The slices have different sizes.")
	}

	var sum float64
	for index := 0; index < len(hist1); index++ {
		numerator := math.Pow(hist1[index] - hist2[index], 2)
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
	if len(hist1) != len(hist2) {
		return 0, errors.New("Could not compare the histograms. The slices have different sizes.")
	}

	var sum float64
	for index := 0; index < len(hist1); index++ {
		sum += math.Pow(hist1[index] - hist2[index], 2)
	}
	return math.Sqrt(sum), nil
}

// NormalizedEuclideanDistance calculates the euclidean distance normalized.
// D = \sqrt{\sum_{i=1}^{n} \frac{(hist1_{i} - hist2_{i})^2}{n}}
// Reference:
// http://www.pbarrett.net/techpapers/euclid.pdf
func NormalizedEuclideanDistance(hist1, hist2 []float64) (float64, error) {

	// Check the histogram sizes
	if len(hist1) != len(hist2) {
		return 0, errors.New("Could not compare the histograms. The slices have different sizes.")
	}

	var sum float64
	n := float64(len(hist1))
	for index := 0; index < len(hist1); index++ {
		sum += math.Pow(hist1[index]-hist2[index], 2) / n
	}
	return math.Sqrt(sum), nil
}

// Intersection calculates the intersection between two histograms
// by the following formula:
// D = \sum_{i=1}^{n} min(hist1_{i}, hist2_{i})
// IMPORTANT: This is inversely proportional, higher the intersection higher the similarity.
// References:
// http://blog.datadive.net/histogram-intersection-for-change-detection/
// https://dsp.stackexchange.com/questions/18065/histogram-intersection-with-two-different-bin-sizes
// https://mpatacchiola.github.io/blog/2016/11/12/the-simplest-classifier-histogram-intersection.html
func Intersection(hist1, hist2 []float64) (float64, error) {
	// Check the histogram sizes
	if len(hist1) != len(hist2) {
		return 0, errors.New("Could not compare the histograms. The slices have different sizes.")
	}

	var sum float64
	for index := 0; index < len(hist1); index++ {
		sum += min(hist1[index], hist2[index])
	}
	return sum, nil
}

// normalizedIntersection calculates the intersection between two histograms
// and normalizes the result by dividing it by the sum of the hist2
// D = \frac{\sum_{i=1}^{n} min(hist1_{i}, hist2_{i})}{max(\sum_{i=1}^{n}hist1_{i},\sum_{i=1}^{n}hist2_{i})}
// References:
// https://dsp.stackexchange.com/questions/18065/histogram-intersection-with-two-different-bin-sizes
// https://mpatacchiola.github.io/blog/2016/11/12/the-simplest-classifier-histogram-intersection.html
func NormalizedIntersection(hist1, hist2 []float64) (float64, error) {
	// Check the histogram sizes
	if len(hist1) != len(hist2) {
		return 0, errors.New("Could not compare the histograms. The slices have different sizes.")
	}

	intersection, _ := Intersection(hist1, hist2)

	sum1 := sum(hist1)
	sum2 := sum(hist2)

	if sum1 > sum2 {
		return intersection / sum1, nil
	} else {
		return intersection / sum2, nil
	}
}

// AbsoluteValueNorm calculates the absolute value normalized
// by the following formula:
// D = \sum_{i=1}^{n} \left | hist1_{i} - hist2_{i} \right |
func AbsoluteValueNorm(hist1, hist2 []float64) (float64, error) {
	// Check the histogram sizes
	if len(hist1) != len(hist2) {
		return 0, errors.New("Could not compare the histograms. The slices have different sizes.")
	}

	var sum float64
	for index := 0; index < len(hist1); index++ {
		sum += abs(hist1[index] - hist2[index])
	}
	return sum, nil
}