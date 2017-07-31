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

func equalSlices(slice1, slice2 []uint8) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for index := 0; index < len(slice1); index++ {
		if slice1[index] != slice2[index] {
			return false
		}
	}
	return true
}

func TestGetHistogram(t *testing.T) {
	var pixels [][]uint8

	_, err := GetHistogram(pixels, 1, 1)
	if err == nil {
		t.Error("Expected an error.")
	}

	row1 := []uint8{255,255,255,255,255,255}
	row2 := []uint8{0,0,0,0,0,0}
	pixels = append(pixels, row1)
	pixels = append(pixels, row2)
	pixels = append(pixels, row2)
	pixels = append(pixels, row2)
	pixels = append(pixels, row2)
	pixels = append(pixels, row1)

	_, err = GetHistogram(pixels, 0, 1)
	if err == nil {
		t.Error("Expected an error.")
	}

	_, err = GetHistogram(pixels, 1, 0)
	if err == nil {
		t.Error("Expected an error.")
	}

	expectedHist := make([]uint8, 256)
	expectedHist[0] = 24
	expectedHist[255] = 12

	hist, err := GetHistogram(pixels, 1, 1)
	if err != nil {
		t.Error(
			"Expected no errors.",
			"Received error:", err,
		)
	}
	if !equalSlices(hist, expectedHist) {
		t.Error("The histograms are different")
	}

	expectedHist = make([]uint8, 1024)
	expectedHist[0] = 6
	expectedHist[255] = 3
	expectedHist[256] = 6
	expectedHist[511] = 3
	expectedHist[512] = 6
	expectedHist[767] = 3
	expectedHist[768] = 6
	expectedHist[1023] = 3

	hist, err = GetHistogram(pixels, 2, 2)

	if err != nil {
		t.Error(
			"Expected no errors.",
			"Received error:", err,
		)
	}
	if !equalSlices(hist, expectedHist) {
		t.Error("The histograms are different")
	}
}

func TestCalcHistogramDist(t *testing.T) {
	var hist1 []uint8
	var hist2 []uint8

	var index uint8
	for index = 0; index < 100; index++ {
		hist1 = append(hist1, uint8(index))
		hist2 = append(hist2, uint8(index))
	}

	dist, _ := CalcHistogramDist(hist1, hist2)
	if !floatEquals(dist, 0.0) {
		t.Error(
			"Expected distance:", 0.0,
			"Received distance:", dist,
		)
	}

	hist1 = nil
	hist2 = nil

	for index = 0; index < 100; index++ {
		hist1 = append(hist1, uint8(index))
		hist2 = append(hist2, uint8(index+1))
	}

	dist, _ = CalcHistogramDist(hist1, hist2)
	if !floatEquals(dist, 10.0) {
		t.Error(
			"Expected distance:", 10.0,
			"Received distance:", dist,
		)
	}
}
