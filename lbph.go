// Performs face recognition using lbph
package lbph

import (
	"image"
  "errors"
	_ "image/gif"
	_ "image/png"
	_ "image/jpeg"

	"github.com/kelvins/imgproc"
)

func checkInputData(images []image.Image) error {
  // Get the image bounds
	bounds := images[0].Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	for index := 0; index < len(images); index++ {
		// Check if the image is in grayscale
		if !imgproc.IsGrayscale(images[index]) {
  		return errors.New("One or more images are not in grayscale")
		}

		b := images[index].Bounds()

		// Check if all images have the same size
		if b.Max.X != width || b.Max.Y != height {
			return errors.New("One or more images have different sizes")
		}
	}
  return nil
}

// Function used to train the algorithm
func Train(images []image.Image, labels []string) error {

	if len(images) != len(labels) {
		return errors.New("Slices have different sizes")
	}
	if len(images) == 0 {
		return errors.New("Empty vector")
	}

  // Check if the input data is in the correct format
  err := checkInputData(images)
  if err != nil {
    return err
  }

	return nil
}

func Predict(image image.Image) (string, float64, error) {
	return "", 0.0, nil
}
