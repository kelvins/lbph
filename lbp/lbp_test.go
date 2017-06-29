package lbp

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

func TestApplyLBP(t *testing.T) {
	img, err := loadImage("../dataset/test/2.png")
	if err != nil {
		t.Error(err)
	}

	expectedlbp := []int64{93, 183, 91, 173, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	lbp, err := ApplyLBP(img)
	if err != nil {
		t.Error(err)
	}
	if len(lbp) == len(expectedlbp) {
		for index := 0; index < len(lbp); index++ {
			/*if lbp[index] != expectedlbp[index] {
				t.Error(
					"Expected value: ", expectedlbp[index],
					"Received value: ", lbp[index],
				)
			}*/
		}
	} else {
		t.Error("Different size")
	}
}
