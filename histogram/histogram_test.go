package histogram

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHistogram(t *testing.T) {
	var pixels [][]uint8

	_, err := GetHistogram(pixels, 1, 1)
	assert.NotNil(t, err)

	row1 := []uint8{255, 255, 255, 255, 255, 255}
	row2 := []uint8{0, 0, 0, 0, 0, 0}
	pixels = append(pixels, row1)
	pixels = append(pixels, row2)
	pixels = append(pixels, row2)
	pixels = append(pixels, row2)
	pixels = append(pixels, row2)
	pixels = append(pixels, row1)

	_, err = GetHistogram(pixels, 0, 1)
	assert.NotNil(t, err)

	_, err = GetHistogram(pixels, 1, 0)
	assert.NotNil(t, err)

	expectedHist := make([]uint8, 256)
	expectedHist[0] = 24
	expectedHist[255] = 12

	hist, err := GetHistogram(pixels, 1, 1)
	assert.Nil(t, err)
	assert.Equal(t, hist, expectedHist, "The histograms should be equal")

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
	assert.Nil(t, err)
	assert.Equal(t, hist, expectedHist, "The histograms should be equal")
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
	assert.Equal(t, dist, 0.0, "The distance should be 0")

	hist1 = nil
	hist2 = nil

	for index = 0; index < 100; index++ {
		hist1 = append(hist1, uint8(index))
		hist2 = append(hist2, uint8(index+1))
	}

	dist, _ = CalcHistogramDist(hist1, hist2)
	assert.Equal(t, dist, 10.0, "The distance should be equal to 10")
}
