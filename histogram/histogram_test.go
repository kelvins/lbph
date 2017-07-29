package histogram

import (
	"testing"
)

var EPSILON float64 = 0.00000001
func floatEquals(a, b float64) bool {
	if (a - b) < EPSILON && (b - a) < EPSILON {
		return true
	}
	return false
}

func TestCalcHistogramDist(t *testing.T) {
	var hist1 []uint8
	var hist2 []uint8

	var index uint8
	for index = 0; index < 100; index++ {
		hist1 = append(hist1, uint8(index))
		hist2 = append(hist2, uint8(index))
	}

	dist := CalcHistogramDist(hist1, hist2)
	if !floatEquals(dist, 0.0) {
		t.Error(
			"Expected distance:", 0.0,
			"Received distance:", dist,
		)
	}

	for index = 0; index < 100; index++ {
		hist1 = append(hist1, uint8(index))
		hist2 = append(hist2, uint8(index+1))
	}

	dist = CalcHistogramDist(hist1, hist2)
	if !floatEquals(dist, 10.0) {
		t.Error(
			"Expected distance:", 10.0,
			"Received distance:", dist,
		)
	}
}
