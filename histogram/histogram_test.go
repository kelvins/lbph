package histogram

import (
	"testing"

	"github.com/kelvins/lbph/metric"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	var pixels [][]uint8

	_, err := Calculate(pixels, 1, 1)
	assert.NotNil(t, err)

	row1 := []uint8{255, 255, 255, 255, 255, 255}
	row2 := []uint8{0, 0, 0, 0, 0, 0}
	pixels = append(pixels, row1)
	pixels = append(pixels, row2)
	pixels = append(pixels, row2)
	pixels = append(pixels, row2)
	pixels = append(pixels, row2)
	pixels = append(pixels, row1)

	_, err = Calculate(pixels, 0, 1)
	assert.NotNil(t, err)

	_, err = Calculate(pixels, 1, 0)
	assert.NotNil(t, err)

	expectedHist := make([]float64, 256)
	expectedHist[0] = 24
	expectedHist[255] = 12

	hist, err := Calculate(pixels, 1, 1)
	assert.Nil(t, err)
	assert.Equal(t, expectedHist, hist, "The histograms should be equal")

	expectedHist = make([]float64, 1024)
	expectedHist[0] = 6
	expectedHist[255] = 3
	expectedHist[256] = 6
	expectedHist[511] = 3
	expectedHist[512] = 6
	expectedHist[767] = 3
	expectedHist[768] = 6
	expectedHist[1023] = 3

	hist, err = Calculate(pixels, 2, 2)
	assert.Nil(t, err)
	assert.Equal(t, expectedHist, hist, "The histograms should be equal")
}

func TestCompare(t *testing.T) {
	var hist1 []float64
	var hist2 []float64

	for index := 0; index < 100; index++ {
		hist1 = append(hist1, float64(index))
		hist2 = append(hist2, float64(index))
	}

	confidence, _ := Compare(hist1, hist2, metric.EuclideanDistance)
	assert.Equal(t, 0.0, confidence, "The confidence should be 0")

	hist1 = nil
	hist2 = nil

	for index := 0; index < 100; index++ {
		hist1 = append(hist1, float64(index))
		hist2 = append(hist2, float64(index+1))
	}

	confidence, _ = Compare(hist1, hist2, metric.EuclideanDistance)
	assert.Equal(t, 10.0, confidence, "The confidence should be equal to 10")
}
