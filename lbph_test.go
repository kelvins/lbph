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

func TestPredict(t *testing.T) {

	var paths []string
	paths = append(paths, "./dataset/train/1.png")
	paths = append(paths, "./dataset/train/2.png")
	paths = append(paths, "./dataset/train/3.png")

	var labels []string
	labels = append(labels, "rocks")
	labels = append(labels, "grass")
	labels = append(labels, "wood")

	var images []image.Image

	for index := 0; index < len(paths); index++ {
		img, err := loadImage(paths[index])
		if err != nil {
			t.Error(err)
		}
		images = append(images, img)
	}

	err := Train(images, labels)
	if err != nil {
		t.Error(err)
	}

	// Table tests
	var tTable = []struct {
		path  string
		label string
	}{
		{"../dataset/test/1.png", "wood"},
		{"../dataset/test/2.png", "rocks"},
		{"../dataset/test/3.png", "grass"},
	}

	// Test with all values in the table
	for _, pair := range tTable {
		img, _ := loadImage(pair.path)
		lbl, dist, err := Predict(img)
		if err != nil {
			t.Error(err)
		}
		if lbl != pair.label {
			t.Error(
				"Expected label", pair.label,
				"Received label", lbl,
			)
		}
		if dist > 500 || dist < 0 {
			t.Error("Received dist : ", dist)
		}
	}
}
