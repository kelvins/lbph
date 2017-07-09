package lbph

import (
	"image"
	"testing"

	"github.com/kelvins/lbph/common"
)

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
		img, err := common.LoadImage(paths[index])
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
		{"./dataset/test/1.png", "wood"},
		{"./dataset/test/2.png", "rocks"},
		{"./dataset/test/3.png", "grass"},
	}

	// Test with all values in the table
	for _, pair := range tTable {
		img, _ := common.LoadImage(pair.path)
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
		if dist > 3000 || dist < 0 {
			t.Error("Received dist : ", dist)
		}
	}
}
