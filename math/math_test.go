package math

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChiSquare(t *testing.T) {
	// Table tests
	var tTable = []struct {
		hist1    []float64
		hist2    []float64
		distance float64
	}{
		{[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9}, []float64{9, 8, 7, 6, 5, 4, 3, 2, 1}, 102.8968253968254},
		{[]float64{1, 1, 1, 1, 1, 1, 1, 1, 1}, []float64{9, 9, 9, 9, 9, 9, 9, 9, 9}, 576.0},
		{[]float64{9, 9, 9, 9, 9, 9, 9, 9, 9}, []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}, 64.00000000000001},
		{[]float64{9, 9, 9, 9, 9, 9, 9, 9, 9}, []float64{9, 9, 9, 9, 9, 9, 9, 9, 9}, 0.0},
		{[]float64{1, 4, 5, 4, 1, 0, 0, 0, 0}, []float64{0, 0, 0, 0, 0, 1, 3, 4, 2}, -1.0},
	}

	// Test with all values in the table
	for _, pair := range tTable {
		distance, err := ChiSquare(pair.hist1, pair.hist2)
		assert.Nil(t, err)
		if pair.distance >= 0 {
			assert.Equal(t, pair.distance, distance)
		} else {
			assert.Equal(t, true, math.IsInf(distance, 0))
		}
	}
}

func TestEuclideanDistance(t *testing.T) {
	// Table tests
	var tTable = []struct {
		hist1    []float64
		hist2    []float64
		distance float64
	}{
		{[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9}, []float64{9, 8, 7, 6, 5, 4, 3, 2, 1}, 15.491933384829668},
		{[]float64{1, 1, 1, 1, 1, 1, 1, 1, 1}, []float64{9, 9, 9, 9, 9, 9, 9, 9, 9}, 24.0},
		{[]float64{9, 9, 9, 9, 9, 9, 9, 9, 9}, []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}, 24.0},
		{[]float64{9, 9, 9, 9, 9, 9, 9, 9, 9}, []float64{9, 9, 9, 9, 9, 9, 9, 9, 9}, 0.0},
		{[]float64{1, 4, 5, 4, 1, 0, 0, 0, 0}, []float64{0, 0, 0, 0, 0, 1, 3, 4, 2}, 9.433981132056603},
	}

	// Test with all values in the table
	for _, pair := range tTable {
		distance, err := EuclideanDistance(pair.hist1, pair.hist2)
		assert.Nil(t, err)
		assert.Equal(t, pair.distance, distance)
	}
}

func TestNormalizedEuclideanDistance(t *testing.T) {
	// Table tests
	var tTable = []struct {
		hist1    []float64
		hist2    []float64
		distance float64
	}{
		{[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9}, []float64{9, 8, 7, 6, 5, 4, 3, 2, 1}, 5.163977794943222},
		{[]float64{1, 1, 1, 1, 1, 1, 1, 1, 1}, []float64{9, 9, 9, 9, 9, 9, 9, 9, 9}, 8.0},
		{[]float64{9, 9, 9, 9, 9, 9, 9, 9, 9}, []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}, 8.0},
		{[]float64{9, 9, 9, 9, 9, 9, 9, 9, 9}, []float64{9, 9, 9, 9, 9, 9, 9, 9, 9}, 0.0},
		{[]float64{1, 4, 5, 4, 1, 0, 0, 0, 0}, []float64{0, 0, 0, 0, 0, 1, 3, 4, 2}, 3.144660377352201},
	}

	// Test with all values in the table
	for _, pair := range tTable {
		distance, err := NormalizedEuclideanDistance(pair.hist1, pair.hist2)
		assert.Nil(t, err)
		assert.Equal(t, pair.distance, distance)
	}
}

func TestIntersection(t *testing.T) {
	// Table tests
	var tTable = []struct {
		hist1    []float64
		hist2    []float64
		distance float64
	}{
		{[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9}, []float64{9, 8, 7, 6, 5, 4, 3, 2, 1}, 40.0},
		{[]float64{1, 1, 1, 1, 1, 1, 1, 1, 1}, []float64{9, 9, 9, 9, 9, 9, 9, 9, 9}, 72.0},
		{[]float64{9, 9, 9, 9, 9, 9, 9, 9, 9}, []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}, 72.0},
		{[]float64{9, 9, 9, 9, 9, 9, 9, 9, 9}, []float64{9, 9, 9, 9, 9, 9, 9, 9, 9}, 0.0},
		{[]float64{1, 4, 5, 4, 1, 0, 0, 0, 0}, []float64{0, 0, 0, 0, 0, 1, 3, 4, 2}, 25.0},
	}

	// Test with all values in the table
	for _, pair := range tTable {
		distance, err := Intersection(pair.hist1, pair.hist2)
		assert.Nil(t, err)
		assert.Equal(t, pair.distance, distance)
	}
}

func TestNormalizedIntersection(t *testing.T) {
	// Table tests
	var tTable = []struct {
		hist1    []float64
		hist2    []float64
		distance float64
	}{
		{[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9}, []float64{9, 8, 7, 6, 5, 4, 3, 2, 1}, 0.6153846153846154},
		{[]float64{1, 1, 1, 1, 1, 1, 1, 1, 1}, []float64{9, 9, 9, 9, 9, 9, 9, 9, 9}, 0.8888888888888888},
		{[]float64{9, 9, 9, 9, 9, 9, 9, 9, 9}, []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}, 0.8888888888888888},
		{[]float64{9, 9, 9, 9, 9, 9, 9, 9, 9}, []float64{9, 9, 9, 9, 9, 9, 9, 9, 9}, 0.0},
		{[]float64{1, 4, 5, 4, 1, 0, 0, 0, 0}, []float64{0, 0, 0, 0, 0, 1, 3, 4, 2}, 1.0},
	}

	// Test with all values in the table
	for _, pair := range tTable {
		distance, err := NormalizedIntersection(pair.hist1, pair.hist2)
		assert.Nil(t, err)
		assert.Equal(t, pair.distance, distance)
	}
}
