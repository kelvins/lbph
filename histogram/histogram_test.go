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
	var hist1 [256]int64
	var hist2 [256]int64

	hist1[0] = 15
	hist2[0] = 5
	hist1[255] = 50
	hist2[255] = 40

	dist := CalcHistogramDist(hist1, hist2)
	if !floatEquals(dist, 14.142135623730951) {
		t.Error(
			"Expected distance:", 14.142135623730951,
			"Received distance:", dist,
		)
	}

	hist1[0] = 15
	hist2[0] = 15
	hist1[255] = 50
	hist2[255] = 50

	dist = CalcHistogramDist(hist1, hist2)
	if !floatEquals(dist, 0.0) {
		t.Error(
			"Expected distance:", 0.0,
			"Received distance:", dist,
		)
	}
}
