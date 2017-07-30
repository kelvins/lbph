// lbph package provides a texture classification using local binary patterns
package lbph

import (
	"errors"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/kelvins/lbph/common"
	"github.com/kelvins/lbph/histogram"
	"github.com/kelvins/lbph/lbp"
)

// Structure used to pass the LBPH parameters
type Parameters struct {
	Radius    uint8
	Neighbors uint8
	GridX     uint8
	GridY     uint8
}

// Store the input data (images and labels) and the calculated histogram
type TrainDataStruct struct {
	Images     []image.Image
	Labels     []string
	Histograms [][]uint8
}

var (
	// Struct that stores the Data loaded by the user
	TrainData TrainDataStruct

	// LBPH parameters
	lbphParameters = Parameters{
		Radius:    1,
		Neighbors: 8,
		GridX:     8,
		GridY:     8,
	}
)

func Init(parameters Parameters) {

	if parameters.Radius <= 0 {
		parameters.Radius = 1
	}

	if parameters.Neighbors <= 0 {
		parameters.Neighbors = 8
	}

	if parameters.GridX <= 0 {
		parameters.GridX = 8
	}

	if parameters.GridY <= 0 {
		parameters.GridY = 8
	}

	lbphParameters = parameters
}

// Train function is used to train the LBPH algorithm
func Train(images []image.Image, labels []string) error {
	// Clear the data structure
	TrainData = TrainDataStruct{}

	// Check if the images and labels slices have the same size
	if len(images) != len(labels) {
		return errors.New("Slices have different sizes")
	}

	// Check if the images slice is not empty
	// As we already checked if the slices have the same size we
	// don't need to check if the labels slice is empty
	if len(images) == 0 {
		return errors.New("Empty vector")
	}

	// Call the CheckInputData from the common package
	// It will check if all images have the same size
	err := common.CheckImagesSizes(images)
	if err != nil {
		return err
	}

	// Call the GetHistogram from the histogram package
	// It will run the LBP operation and generate the histogram for each image
	var histograms [][]uint8
	for index := 0; index < len(images); index++ {
		// Calculate the LBP operation
		pixels, err := lbp.ApplyLBP(images[index], lbphParameters.Radius, lbphParameters.Neighbors)
		if err != nil {
			return err
		}
		hist, err := histogram.GetHistogram(pixels, lbphParameters.GridX, lbphParameters.GridY)
		if err != nil {
			return err
		}
		histograms = append(histograms, hist)
	}

	// Store the current data that we are working on
	TrainData = TrainDataStruct{
		Images:     images,
		Labels:     labels,
		Histograms: histograms,
	}

	// Everything is ok, return nil
	return nil
}

// Predict function is finds the closest image/group based on the images used in the Train function
func Predict(img image.Image) (string, float64, error) {
	// Check if the image passed by parameter is nil
	if img == nil {
		return "", 0.0, errors.New("Image is nil")
	}

	// If we don't have histograms to compare, probably the Train function was
	// not called or has occurred an error and it was not correctly treated
	if len(TrainData.Histograms) == 0 {
		return "", 0.0, errors.New("Could not get the image histogram")
	}

	// Calculate the LBP operation
	pixels, err := lbp.ApplyLBP(img, lbphParameters.Radius, lbphParameters.Neighbors)
	if err != nil {
		return "", 0.0, err
	}
	// Calculate the histogram for the current image
	hist, err := histogram.GetHistogram(pixels, lbphParameters.GridX, lbphParameters.GridY)
	if err != nil {
		return "", 0.0, errors.New("Could not get the image histogram")
	}

	// Search for the closest histogram based on the histograms calculated in the Train function
	minValue := histogram.CalcHistogramDist(hist, TrainData.Histograms[0])
	minIndex := 0
	for index := 1; index < len(TrainData.Histograms); index++ {
		// Calculate the distance from the current histogram
		dist := histogram.CalcHistogramDist(hist, TrainData.Histograms[index])
		// If it is closer, save the minValue and the index
		if dist < minValue {
			minValue = dist
			minIndex = index
		}
	}

	// Return the label corresponding to the closest histogram,
	// the distance (minValue) and the error (nil)
	return TrainData.Labels[minIndex], minValue, nil
}
