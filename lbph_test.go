package lbph

import (
	"image"
	"testing"
	"os"

	"github.com/kelvins/lbph/metric"

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
		img, err := LoadImage(paths[index])
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

	Metric = metric.EuclideanDistance

	// Test with all values in the table
	for _, pair := range tTable {
		img, _ := LoadImage(pair.path)
		lbl, conf, err := Predict(img)
		assert.Nil(t, err)
		assert.Equal(t, lbl, pair.label, "The labels should be equal")
		if conf > 500 || conf < 0 {
			assert.Equal(t, conf, 250.0, "The confidence should be between 0 and 500")
		}
	}

	labels = []string{"rocks", "grass", "wood"}

	// Test with all values in the table
	for index := 0; index < len(labels); index++ {
		trainData := GetTrainingData()
		assert.Equal(t, trainData.Labels[index], labels[index], "The labels should be equal")
	}
}
