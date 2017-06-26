// Performs face recognition using lbph
package lbph

import (
	"errors"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/kelvins/lbph/common"
	"github.com/kelvins/lbph/histogram"
	"github.com/kelvins/lbph/structs"
)

var Data structs.Data

// Function used to train the algorithm
func Train(images []image.Image, labels []string) error {

	if len(images) != len(labels) {
		return errors.New("Slices have different sizes")
	}
	if len(images) == 0 {
		return errors.New("Empty vector")
	}

	// Check if the input data is in the correct format
	err := common.CheckInputData(images)
	if err != nil {
		return err
	}

	var histograms [][256]int64
	for index := 0; index < len(images); index++ {
		hist, err := histogram.GetHistogram(images[index])
		if err != nil {
			return err
		}
		histograms = append(histograms, hist)
	}

	// This conditional must never occurs
	if len(histograms) == 0 {
		return errors.New("None histogram was calculated")
	}

	Data = structs.Data{
		Images:     images,
		Labels:     labels,
		Histograms: histograms,
	}

	return nil
}

func Predict(img image.Image) (string, float64, error) {
	hist, err := histogram.GetHistogram(img)
	if err != nil {
		return "", 0.0, errors.New("Could not get the image histogram")
	}
	var min float64
	var i int
	for index := 0; index < len(Data.Histograms); index++ {
		if index == 0 {
			i = index
			min = histogram.GetHistogramDist(hist, Data.Histograms[index])
		} else {
			x := histogram.GetHistogramDist(hist, Data.Histograms[index])
			if x < min {
				min = x
				i = index
			}
		}
	}
	return Data.Labels[i], min, nil
}
