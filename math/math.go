package math

import (
	"errors"
	"math"
)

// min function returns the minimum value.
func min(value1, value2 uint8) uint8 {
	if value1 < value2 {
		return value1
	} else {
		return value2
	}
}

// sum function returns the sum of a uint8 slice.
func sum(slice []uint8) float64 {
	var sum float64
	for _, value := range slice {
		sum += float64(value)
	}
	return sum
}

// chiSquare calculates the distance between two histograms using
// the chi square statistic.
// References:
// http://file.scirp.org/Html/8-72278_30995.htm
// https://www.google.com/patents/WO2007080817A1?cl=en
func ChiSquare(hist1, hist2 []uint8) (float64, error) {
	// Check the histogram sizes
	if len(hist1) != len(hist2) {
		return 0, errors.New("Could not compare the histograms. The slices have different sizes.")
	}

	var sum float64
	for index := 0; index < len(hist1); index++ {
		numerator := float64((hist1[index] - hist2[index]) * (hist1[index] - hist2[index]))
		denominator := float64(hist1[index])
		sum += numerator / denominator
	}
	return sum, nil
}

// EuclideanDistance calculates the euclidean distance between two histograms
// by the following formula: sqrt(sum((h1(i)-h2(i))^2))
// Reference: http://www.pbarrett.net/techpapers/euclid.pdf
func EuclideanDistance(hist1, hist2 []uint8) (float64, error) {

	// Check the histogram sizes
	if len(hist1) != len(hist2) {
		return 0, errors.New("Could not compare the histograms. The slices have different sizes.")
	}

	var sum float64
	for index := 0; index < len(hist1); index++ {
		sum += float64((hist1[index] - hist2[index]) * (hist1[index] - hist2[index]))
	}
	return math.Sqrt(sum), nil
}

// NormalizedEuclideanDistance calculates the euclidean distance normalized.
// Reference: http://www.pbarrett.net/techpapers/euclid.pdf
func NormalizedEuclideanDistance(hist1, hist2 []uint8) (float64, error) {

	// Check the histogram sizes
	if len(hist1) != len(hist2) {
		return 0, errors.New("Could not compare the histograms. The slices have different sizes.")
	}

	var sum float64
	for index := 0; index < len(hist1); index++ {
		sum += float64((hist1[index]-hist2[index])*(hist1[index]-hist2[index])) / float64(len(hist1))
	}
	return math.Sqrt(sum), nil
}

// Intersection calculates the intersection between two histograms
// by the following formula: sum(min(h1(i), h2(i)))
// IMPORTANT: This is inversely proportional, higher the intersection higher the similarity.
// References:
// http://blog.datadive.net/histogram-intersection-for-change-detection/
// https://dsp.stackexchange.com/questions/18065/histogram-intersection-with-two-different-bin-sizes
// https://mpatacchiola.github.io/blog/2016/11/12/the-simplest-classifier-histogram-intersection.html
func Intersection(hist1, hist2 []uint8) (float64, error) {
	// Check the histogram sizes
	if len(hist1) != len(hist2) {
		return 0, errors.New("Could not compare the histograms. The slices have different sizes.")
	}

	var sum float64
	for index := 0; index < len(hist1); index++ {
		sum += float64(min(hist1[index], hist2[index]))
	}
	return sum, nil
}

// normalizedIntersection calculates the intersection between two histograms
// and normalizes the result by dividing it by the sum of the hist2
// References:
// https://dsp.stackexchange.com/questions/18065/histogram-intersection-with-two-different-bin-sizes
// https://mpatacchiola.github.io/blog/2016/11/12/the-simplest-classifier-histogram-intersection.html
func NormalizedIntersection(hist1, hist2 []uint8) (float64, error) {
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
