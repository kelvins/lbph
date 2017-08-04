// lbph package provides a texture classification using local binary patterns.
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
	"github.com/kelvins/lbph/metrics"
)

// Store the input data (images and labels) and the calculated histogram.
type TrainDataStruct struct {
	Images     []image.Image
	Labels     []string
	Histograms [][]uint8
}

// Structure used to pass the LBPH parameters.
type Parameters struct {
	Radius    uint8
	Neighbors uint8
	GridX     uint8
	GridY     uint8
}

var (
	// Struct that stores the Data loaded by the user.
	trainData = &TrainDataStruct{}

	// LBPH parameters
	lbphParameters = Parameters{
		Radius:    1,
		Neighbors: 8,
		GridX:     8,
		GridY:     8,
	}
)

// The metric used to compare the histograms in the Predict step
var Metric string

func init() {
	// As the trainData is a pointer, the initial state can be nil.
	trainData = nil

	// Use the ChiSquare as the default metric
	Metric = metrics.ChiSquare
}

// Init function is used to set the LBPH parameters based on the Parameters structure.
// It is needed to set the default parameters if something is wrong and
// to reset the trainData if new parameters are defined
func Init(parameters Parameters) {

	// If some parameter is wrong (== 0) set the default one.
	// As the data type is uint8 we don't need to check if it is lower than 0.
	if parameters.Radius == 0 {
		parameters.Radius = 1
	}

	if parameters.Neighbors == 0 {
		parameters.Neighbors = 8
	}

	if parameters.GridX == 0 {
		parameters.GridX = 8
	}

	if parameters.GridY == 0 {
		parameters.GridY = 8
	}

	lbphParameters = parameters

	// Every time the Init function is called the training data is reset,
	// so the user needs to train the algorithm again.
	trainData = nil
}

// GetTrainData is used to get the trainData struct.
// The user can use it to access the images, labels and histograms.
func GetTrainData() TrainDataStruct {
	// Returns the data structure pointed by trainData
	return *trainData
}

// Train function is used to train the LBPH algorithm based on the images and labels passed by parameter.
func Train(images []image.Image, labels []string) error {
	// Clear the data structure
	trainData = nil

	// Check if the slices are not empty
	if len(images) == 0 || len(labels) == 0 {
		return errors.New("At least one of the slices is empty")
	}

	// Check if the images and labels slices have the same size
	if len(images) != len(labels) {
		return errors.New("The slices have different sizes")
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

		// Get the histogram from the current image
		hist, err := histogram.GetHistogram(pixels, lbphParameters.GridX, lbphParameters.GridY)
		if err != nil {
			return err
		}

		// Store the histogram in the 'matrix'
		histograms = append(histograms, hist)
	}

	// Store the current data that we are working on
	trainData = &TrainDataStruct{
		Images:     images,
		Labels:     labels,
		Histograms: histograms,
	}

	// Everything is ok, return nil
	return nil
}

// Predict function is finds the closest image/group based on the images used in the Train function
func Predict(img image.Image) (string, float64, error) {

	// Check if we have data in the trainData struct
	if trainData == nil {
		return "", 0.0, errors.New("The algorithm was not trained yet")
	}

	// Check if the image passed by parameter is nil
	if img == nil {
		return "", 0.0, errors.New("The image passed by parameter is nil")
	}

	// If we don't have histograms to compare, probably the Train function was
	// not called or has occurred an error and it was not correctly treated
	if len(trainData.Histograms) == 0 {
		return "", 0.0, errors.New("There are no histograms in the trainData")
	}

	// Calculate the LBP operation
	pixels, err := lbp.ApplyLBP(img, lbphParameters.Radius, lbphParameters.Neighbors)
	if err != nil {
		return "", 0.0, err
	}

	// Calculate the histogram for the current image
	hist, err := histogram.GetHistogram(pixels, lbphParameters.GridX, lbphParameters.GridY)
	if err != nil {
		return "", 0.0, err
	}

	// Search for the closest histogram based on the histograms calculated in the Train function
	minValue, err := histogram.CalcHistogramDist(hist, trainData.Histograms[0], Metric)
	if err != nil {
		return "", 0.0, err
	}

	minIndex := 0
	for index := 1; index < len(trainData.Histograms); index++ {
		// Calculate the distance from the current histogram
		dist, err := histogram.CalcHistogramDist(hist, trainData.Histograms[index], Metric)
		if err != nil {
			return "", 0.0, err
		}
		// If it is closer, save the minValue and the index
		if dist < minValue {
			minValue = dist
			minIndex = index
		}
	}

	// Return the label corresponding to the closest histogram,
	// the distance (minValue) and the error (nil)
	return trainData.Labels[minIndex], minValue, nil
}
