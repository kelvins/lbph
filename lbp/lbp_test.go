package lbp

import (
	"testing"
	"image"
	"os"

	"github.com/stretchr/testify/assert"
	"github.com/pkg/errors"
)

// LoadImage function is used to provide an easy way to load an image file.
func LoadImage(filePath string) (image.Image, error) {
	// Open the image file
	fImage, err := os.Open(filePath)
	// Check if no error has occurred
	if err != nil {
		return nil, errors.Wrap(err, "Failed to open an image file")
	}

	// Ensure that the image file will be closed
	defer fImage.Close()

	// Decode it to an image "object" (we don't need the format name so we use "_")
	img, _, err := image.Decode(fImage)
	// Check if no error has occurred
	if err != nil {
		return nil, errors.Wrap(err, "Failed decoding the image file")
	}

	return img, nil
}

func TestGetBinary(t *testing.T) {
	// Table tests
	var tTable = []struct {
		value     uint8
		threshold uint8
		result    string
	}{
		{120, 120, "1"},
		{214, 190, "1"},
		{150, 240, "0"},
	}

	// Test with all values in the table
	for _, pair := range tTable {
		result := getBinaryString(pair.value, pair.threshold)
		assert.Equal(t, result, pair.result, "The result should be equal")
	}
}

func TestCalculate(t *testing.T) {
	img, err := LoadImage("../dataset/test/4.png")
	assert.Nil(t, err)

	// Results manually calculated (radius:1 - neighbors:8)
	var expectedLBP [][]uint8
	expectedLBP = append(expectedLBP, []uint8{91, 190, 93, 179})
	expectedLBP = append(expectedLBP, []uint8{238, 245, 255, 206})
	expectedLBP = append(expectedLBP, []uint8{115, 255, 175, 119})
	expectedLBP = append(expectedLBP, []uint8{205, 186, 125, 218})

	pixels, err := Calculate(img, 1, 8)
	assert.Nil(t, err)

	// Check each pixel
	for x := 0; x < len(pixels); x++ {
		for y := 0; y < len(pixels[x]); y++ {
			assert.Equal(t, pixels[x][y], expectedLBP[x][y], "The pixel value should be equal")
		}
	}
}
