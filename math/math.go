package math

import "math"

// min function returns the minimum value
func min(value1, value2 uint8) uint8 {
	if value1 < value2 {
		return value1
	} else {
		return value2
	}
}

// sum function returns the sum of a uint8 slice
func sum(slice []uint8) float64 {
	var sum float64
	for _, value := range slice {
		sum += float64(value)
	}
	return sum
}

// EuclideanDistance calculates the euclidean distance between two histograms
// by the following formula: sqrt(sum((h1(i)-h2(i))^2))
func EuclideanDistance(hist1, hist2 []uint8) float64 {
	var sum float64
	for index := 0; index < len(hist1); index++ {
		sum += float64((hist1[index] - hist2[index]) * (hist1[index] - hist2[index]))
	}
	return math.Sqrt(sum)
}

// Intersection calculates the intersection between two histograms
// by the following formula: sum(min(h1(i), h2(i)))
func Intersection(hist1, hist2 []uint8) float64 {
	var sum float64
	for index := 0; index < len(hist1); index++ {
		sum += float64(min(hist1[index], hist2[index]))
	}
	return sum
}

// normalizedIntersection calculates the intersection between two histograms
// and normalizes the result by dividing it by the sum of the hist2
// https://dsp.stackexchange.com/questions/18065/histogram-intersection-with-two-different-bin-sizes
func NormalizedIntersection(hist1, hist2 []uint8) float64 {
	sum1 := sum(hist1)
	sum2 := sum(hist2)
	var min float64
	if sum1 < sum2 {
		min = sum1
	} else {
		min = sum2
	}
	return Intersection(hist1, hist2) / min
}

// chiSquare calculates the distance between two histograms using
// the chi square statistic. References:
// http://file.scirp.org/Html/8-72278_30995.htm
// https://www.google.com/patents/WO2007080817A1?cl=en
func ChiSquare(hist1, hist2 []uint8) float64 {
	var sum float64
	for index := 0; index < len(hist1); index++ {
		numerator := float64((hist1[index] - hist2[index]) * (hist1[index] - hist2[index]))
		denominator := float64(hist1[index] + hist2[index])
		sum += numerator / denominator
	}
	return sum
}