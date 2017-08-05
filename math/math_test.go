package math

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChiSquare(t *testing.T) {
	// Table tests
	var tTable = []struct {
		hist1      []uint8
		hist2      []uint8
		confidence float64
	}{
		{[]uint8{1, 2, 3, 4, 5, 6, 7, 8, 9}, []uint8{9, 8, 7, 6, 5, 4, 3, 2, 1}, 102.8968253968254},
		{[]uint8{1, 1, 1, 1, 1, 1, 1, 1, 1}, []uint8{9, 9, 9, 9, 9, 9, 9, 9, 9}, 576.0},
		{[]uint8{9, 9, 9, 9, 9, 9, 9, 9, 9}, []uint8{1, 1, 1, 1, 1, 1, 1, 1, 1}, 64.00000000000001},
		{[]uint8{9, 9, 9, 9, 9, 9, 9, 9, 9}, []uint8{9, 9, 9, 9, 9, 9, 9, 9, 9}, 0.0},
		{[]uint8{1, 4, 5, 4, 1, 0, 0, 0, 0}, []uint8{0, 0, 0, 0, 0, 1, 3, 4, 2}, -1.0},
	}

	// Test with all values in the table
	for _, pair := range tTable {
		confidence, err := ChiSquare(pair.hist1, pair.hist2)
		assert.Nil(t, err)
		if pair.confidence >= 0 {
			assert.Equal(t, pair.confidence, confidence)
		} else {
			assert.Equal(t, true, math.IsInf(confidence, 0))
		}
	}
}

func TestEuclideanDistance(t *testing.T) {
	// Table tests
	var tTable = []struct {
		hist1      []uint8
		hist2      []uint8
		confidence float64
	}{
		{[]uint8{1, 2, 3, 4, 5, 6, 7, 8, 9}, []uint8{9, 8, 7, 6, 5, 4, 3, 2, 1}, 15.491933384829668},
		{[]uint8{1, 1, 1, 1, 1, 1, 1, 1, 1}, []uint8{9, 9, 9, 9, 9, 9, 9, 9, 9}, 24.0},
		{[]uint8{9, 9, 9, 9, 9, 9, 9, 9, 9}, []uint8{1, 1, 1, 1, 1, 1, 1, 1, 1}, 24.0},
		{[]uint8{9, 9, 9, 9, 9, 9, 9, 9, 9}, []uint8{9, 9, 9, 9, 9, 9, 9, 9, 9}, 0.0},
		{[]uint8{1, 4, 5, 4, 1, 0, 0, 0, 0}, []uint8{0, 0, 0, 0, 0, 1, 3, 4, 2}, 9.433981132056603},
	}

	// Test with all values in the table
	for _, pair := range tTable {
		confidence, err := EuclideanDistance(pair.hist1, pair.hist2)
		assert.Nil(t, err)
		assert.Equal(t, pair.confidence, confidence)
	}
}

func TestNormalizedEuclideanDistance(t *testing.T) {
	// Table tests
	var tTable = []struct {
		hist1      []uint8
		hist2      []uint8
		confidence float64
	}{
		{[]uint8{1, 2, 3, 4, 5, 6, 7, 8, 9}, []uint8{9, 8, 7, 6, 5, 4, 3, 2, 1}, 5.163977794943222},
		{[]uint8{1, 1, 1, 1, 1, 1, 1, 1, 1}, []uint8{9, 9, 9, 9, 9, 9, 9, 9, 9}, 8.0},
		{[]uint8{9, 9, 9, 9, 9, 9, 9, 9, 9}, []uint8{1, 1, 1, 1, 1, 1, 1, 1, 1}, 8.0},
		{[]uint8{9, 9, 9, 9, 9, 9, 9, 9, 9}, []uint8{9, 9, 9, 9, 9, 9, 9, 9, 9}, 0.0},
		{[]uint8{1, 4, 5, 4, 1, 0, 0, 0, 0}, []uint8{0, 0, 0, 0, 0, 1, 3, 4, 2}, 3.144660377352201},
	}

	// Test with all values in the table
	for _, pair := range tTable {
		confidence, err := NormalizedEuclideanDistance(pair.hist1, pair.hist2)
		assert.Nil(t, err)
		assert.Equal(t, pair.confidence, confidence)
	}
}

func TestIntersection(t *testing.T) {
	// Table tests
	var tTable = []struct {
		hist1      []uint8
		hist2      []uint8
		confidence float64
	}{
		{[]uint8{1, 2, 3, 4, 5, 6, 7, 8, 9}, []uint8{9, 8, 7, 6, 5, 4, 3, 2, 1}, 25},
		{[]uint8{1, 1, 1, 1, 1, 1, 1, 1, 1}, []uint8{9, 9, 9, 9, 9, 9, 9, 9, 9}, 9.0},
		{[]uint8{9, 9, 9, 9, 9, 9, 9, 9, 9}, []uint8{1, 1, 1, 1, 1, 1, 1, 1, 1}, 9.0},
		{[]uint8{9, 9, 9, 9, 9, 9, 9, 9, 9}, []uint8{9, 9, 9, 9, 9, 9, 9, 9, 9}, 81.0},
		{[]uint8{1, 4, 5, 4, 1, 0, 0, 0, 0}, []uint8{0, 0, 0, 0, 0, 1, 3, 4, 2}, 0.0},
	}

	// Test with all values in the table
	for _, pair := range tTable {
		confidence, err := Intersection(pair.hist1, pair.hist2)
		assert.Nil(t, err)
		assert.Equal(t, pair.confidence, confidence)
	}
}

func TestNormalizedIntersection(t *testing.T) {
	// Table tests
	var tTable = []struct {
		hist1      []uint8
		hist2      []uint8
		confidence float64
	}{
		{[]uint8{1, 2, 3, 4, 5, 6, 7, 8, 9}, []uint8{9, 8, 7, 6, 5, 4, 3, 2, 1}, 0.5555555555555556},
		{[]uint8{1, 1, 1, 1, 1, 1, 1, 1, 1}, []uint8{9, 9, 9, 9, 9, 9, 9, 9, 9}, 0.1111111111111111},
		{[]uint8{9, 9, 9, 9, 9, 9, 9, 9, 9}, []uint8{1, 1, 1, 1, 1, 1, 1, 1, 1}, 0.1111111111111111},
		{[]uint8{9, 9, 9, 9, 9, 9, 9, 9, 9}, []uint8{9, 9, 9, 9, 9, 9, 9, 9, 9}, 1.0},
		{[]uint8{1, 4, 5, 4, 1, 0, 0, 0, 0}, []uint8{0, 0, 0, 0, 0, 1, 3, 4, 2}, 0.0},
	}

	// Test with all values in the table
	for _, pair := range tTable {
		confidence, err := NormalizedIntersection(pair.hist1, pair.hist2)
		assert.Nil(t, err)
		assert.Equal(t, pair.confidence, confidence)
	}
}
