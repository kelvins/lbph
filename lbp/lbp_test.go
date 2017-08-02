package lbp

import (
	"testing"

	"github.com/kelvins/lbph/common"

	"github.com/stretchr/testify/assert"
)

func TestApplyLBP(t *testing.T) {
	img, err := common.LoadImage("../dataset/test/4.png")
	assert.Nil(t, err)

	// Results manually calculated (radius:1 - neighbors:8)
	var expectedLBP [][]uint8
	expectedLBP = append(expectedLBP, []uint8{91, 190, 93, 179})
	expectedLBP = append(expectedLBP, []uint8{238, 245, 255, 206})
	expectedLBP = append(expectedLBP, []uint8{115, 255, 175, 119})
	expectedLBP = append(expectedLBP, []uint8{205, 186, 125, 218})

	pixels, err := ApplyLBP(img, 1, 8)
	assert.Nil(t, err)

	// Check each pixel
	for x := 0; x < len(pixels); x++ {
		for y := 0; y < len(pixels[x]); y++ {
			assert.Equal(t, pixels[x][y], expectedLBP[x][y], "The pixel value should be equal")
		}
	}
}
