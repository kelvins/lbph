package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEuclideanDistance(t *testing.T) {
	// Table tests
	var tTable = []struct {
		hist1  []uint8
		hist2  []uint8
		result float64
	}{
		{[]uint8{1,2,3,4,5,6,7,8,9}, []uint8{9,8,7,6,5,4,3,2,1}, 15.491933384829668},
		{[]uint8{1,1,1,1,1,1,1,1,1}, []uint8{9,9,9,9,9,9,9,9,9}, 24.0},
		{[]uint8{9,9,9,9,9,9,9,9,9}, []uint8{1,1,1,1,1,1,1,1,1}, 24.0},
		{[]uint8{9,9,9,9,9,9,9,9,9}, []uint8{9,9,9,9,9,9,9,9,9}, 0.0},
		{[]uint8{1,4,5,4,1,0,0,0,0}, []uint8{0,0,0,0,0,1,3,4,2}, 9.433981132056603},
	}

	// Test with all values in the table
	for _, pair := range tTable {
		distance, err := EuclideanDistance(pair.hist1, pair.hist2)
		assert.Nil(t, err)
		assert.Equal(t, pair.result, distance)
	}
}

func TestNormalizedEuclideanDistance(t *testing.T) {
	// Table tests
	var tTable = []struct {
		hist1  []uint8
		hist2  []uint8
		result float64
	}{
		{[]uint8{1,2,3,4,5,6,7,8,9}, []uint8{9,8,7,6,5,4,3,2,1}, 5.163977794943222},
		{[]uint8{1,1,1,1,1,1,1,1,1}, []uint8{9,9,9,9,9,9,9,9,9}, 8.0},
		{[]uint8{9,9,9,9,9,9,9,9,9}, []uint8{1,1,1,1,1,1,1,1,1}, 8.0},
		{[]uint8{9,9,9,9,9,9,9,9,9}, []uint8{9,9,9,9,9,9,9,9,9}, 0.0},
		{[]uint8{1,4,5,4,1,0,0,0,0}, []uint8{0,0,0,0,0,1,3,4,2}, 3.144660377352201},
	}

	// Test with all values in the table
	for _, pair := range tTable {
		distance, err := NormalizedEuclideanDistance(pair.hist1, pair.hist2)
		assert.Nil(t, err)
		assert.Equal(t, pair.result, distance)
	}
}
