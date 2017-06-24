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

func TestGetSize(t *testing.T) {
	// Table tests
	var tTable = []struct {
		path   string
		width  int
		height int
	}{
		{"./test/1.png", 200, 200},
		{"./test/2.png", 6, 6},
		{"./test/3.png", 256, 256},
	}

	// Test with all values in the table
	for _, pair := range tTable {
		img, _ := loadImage(pair.path)
		width, height := getSize(img)
		if width != pair.width {
			t.Error(
				"Expected: ", pair.width,
				"Received: ", width,
			)
		}
		if height != pair.height {
			t.Error(
				"Expected: ", pair.height,
				"Received: ", height,
			)
		}
	}
}

func TestIsGrayscale(t *testing.T) {
	// Table tests
	var tTable = []struct {
		path string
		res  bool
	}{
		{"./test/1.png", true},
		{"./test/2.png", true},
		{"./test/3.png", false},
	}

	// Test with all values in the table
	for _, pair := range tTable {
		img, _ := loadImage(pair.path)
		res := isGrayscale(img)
		if res != pair.res {
			t.Error(
				"Expected: ", pair.res,
				"Received: ", res,
			)
		}
	}
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
	images = nil

	// No error
	img, err = loadImage("./test/1.png")
	if err != nil {
		t.Error(err)
	}
	images = append(images, img)
	err = checkInputData(images)
	if err != nil {
		t.Error("Expected: nil. Received: ", err)
	}
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
		result := getBinary(pair.value, pair.threshold)
		if result != pair.result {
			t.Error(
				"Expected: ", pair.result,
				"Received: ", result,
			)
		}
	}
}

func TestPredict(t *testing.T) {

	var paths []string
	paths = append(paths, "./train/1.png")
	paths = append(paths, "./train/2.png")
	paths = append(paths, "./train/3.png")

	var labels []string
	labels = append(labels, "1")
	labels = append(labels, "2")
	labels = append(labels, "3")

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

	img, err := loadImage("./test/1.png")
	if err != nil {
		t.Error(err)
	}

	lbl, dist, err := Predict(img)
	if err != nil {
		t.Error(err)
	}
	if lbl != "3" {
		t.Error("Expected label == 3 - Received", lbl)
	}
	if dist > 500 || dist < 0 {
		t.Error("Received dist : ", dist)
	}
}
