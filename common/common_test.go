package common

import (
	"errors"
	"image"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadImage(t *testing.T) {
	// Table tests
	var tTable = []struct {
		path string
		err  error
	}{
		{"testingInvalid.png", errors.New("Invalid path")},
		{"../dataset/test/1.png", nil},
	}

	for _, pair := range tTable {
		img, err := LoadImage(pair.path)
		if pair.err == nil {
			assert.Nil(t, err)
			assert.NotNil(t, img)
		} else {
			assert.Error(t, err)
			assert.Nil(t, img)
		}
	}
}

func TestGetSize(t *testing.T) {
	// Table tests
	var tTable = []struct {
		path   string
		width  int
		height int
	}{
		{"../dataset/test/1.png", 200, 200},
		{"../dataset/test/2.png", 200, 200},
		{"../dataset/test/3.png", 200, 200},
		{"../dataset/test/4.png", 6, 6},
		{"../dataset/test/5.png", 256, 256},
	}

	// Test with all values in the table
	for _, pair := range tTable {
		img, _ := LoadImage(pair.path)
		width, height := GetSize(img)
		assert.Equal(t, width, pair.width, "The width should be equal")
		assert.Equal(t, height, pair.height, "The height should be equal")
	}
}

func TestCheckImagesSizes(t *testing.T) {
	var images []image.Image

	// Images have different sizes
	var paths []string
	paths = append(paths, "../dataset/test/1.png")
	paths = append(paths, "../dataset/test/4.png")

	for index := 0; index < len(paths); index++ {
		img, _ := LoadImage(paths[index])
		images = append(images, img)
	}

	err := CheckImagesSizes(images)
	assert.NotNil(t, err)

	images = nil

	// No error
	img, _ := LoadImage("../dataset/test/1.png")
	images = append(images, img)
	img, _ = LoadImage("../dataset/test/2.png")
	images = append(images, img)

	err = CheckImagesSizes(images)
	assert.Nil(t, err)
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
		result := GetBinaryString(pair.value, pair.threshold)
		assert.Equal(t, result, pair.result, "The result should be equal")
	}
}

func TestGetPixels(t *testing.T) {
	img, err := LoadImage("../dataset/test/4.png")
	if err != nil {
		t.Error(err)
	}
	pixels := GetPixels(img)

	var expectedPixels [][]uint8
	expectedPixels = append(expectedPixels, []uint8{0, 255, 0, 255, 0, 255})
	expectedPixels = append(expectedPixels, []uint8{255, 255, 255, 255, 255, 0})
	expectedPixels = append(expectedPixels, []uint8{0, 255, 255, 0, 255, 255})
	expectedPixels = append(expectedPixels, []uint8{255, 255, 0, 255, 255, 0})
	expectedPixels = append(expectedPixels, []uint8{0, 255, 255, 255, 255, 255})
	expectedPixels = append(expectedPixels, []uint8{255, 0, 255, 0, 255, 0})

	assert.Equal(t, len(pixels), len(expectedPixels), "The length of the slices should be equal")
	assert.Equal(t, len(pixels[0]), len(expectedPixels[0]), "The length of the slices should be equal")

	for row := 0; row < len(pixels); row++ {
		for col := 0; col < len(pixels[0]); col++ {
			assert.Equal(t, pixels[row][col], expectedPixels[row][col], "The pixel value should be equal")
		}
	}
}
