package lbph

import (
	"image"
	"os"
	"testing"
)

func loadImage(filePath string) (image.Image, error) {
	// Open the file image
	fImage, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	// Ensure that the image file will be closed
	defer fImage.Close()

	// Convert it to an image "object"
	img, _, err := image.Decode(fImage)

	if err != nil {
		return nil, err
	}

	return img, nil
}

func TestCheckInputData(t *testing.T) {
	// Image is not in grayscale
	var images []image.Image
	img, err := loadImage("./test/3.png")
	if err != nil {
		t.Error(err)
	}
	images = append(images, img)
	err = checkInputData(images)
	if err == nil {
		t.Error("Expected: Image is not in grayscale. Received: nil")
	}
	images = nil

	// Images have different sizes
	var paths []string
	paths = append(paths, "./test/1.png")
	paths = append(paths, "./test/2.png")

	for index := 0; index < len(paths); index++ {
		img, err := loadImage(paths[index])
		if err != nil {
			t.Error(err)
		}
		images = append(images, img)
	}
	err = checkInputData(images)
	if err == nil {
		t.Error("Expected: Images have different sizes. Received: nil")
	}
}
