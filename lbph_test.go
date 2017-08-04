package lbph

import (
	"image"
	"testing"

	"github.com/kelvins/lbph/common"

	"github.com/stretchr/testify/assert"
)

func TestPredict(t *testing.T) {

	parameters := Parameters{
		Radius:    1,
		Neighbors: 8,
		GridX:     8,
		GridY:     8,
	}

	Init(parameters)

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
		assert.Nil(t, err)
		images = append(images, img)
	}

	err := Train(images, labels)
	assert.Nil(t, err)

	// Table tests
	var tTable = []struct {
		path  string
		label string
	}{
		{"./dataset/test/1.png", "wood"},
		{"./dataset/test/2.png", "rocks"},
		{"./dataset/test/3.png", "grass"},
	}

	Metric = EuclideanDistance

	// Test with all values in the table
	for _, pair := range tTable {
		img, _ := common.LoadImage(pair.path)
		lbl, dist, err := Predict(img)
		assert.Nil(t, err)
		assert.Equal(t, lbl, pair.label, "The labels should be equal")
		if dist > 500 || dist < 0 {
			assert.Equal(t, dist, 250, "The distance should be between 0 and 500")
		}
	}

	labels = []string{"rocks", "grass", "wood"}

	// Test with all values in the table
	for index := 0; index < len(labels); index++ {
		trainData := GetTrainData()
		assert.Equal(t, trainData.Labels[index], labels[index], "The labels should be equal")
	}
}
