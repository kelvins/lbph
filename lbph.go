// lbph package provides a texture classification using local binary patterns.
package lbph

import (
	"errors"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/kelvins/lbph/histogram"
	"github.com/kelvins/lbph/lbp"
	"github.com/kelvins/lbph/metric"
)

// TrainingData struct is used to store the input data (images and labels)
// and each calculated histogram.
type TrainingData struct {
	Images     []image.Image
	Labels     []string
	Histograms [][]float64
}

// Parameters struct is used to pass the LBPH parameters.
type Parameters struct {
	Radius    uint8
	Neighbors uint8
	GridX     uint8
	GridY     uint8
}

// trainData struct stores the TrainingData loaded by the user.
// It needs to be a pointer because the first state will be nil.
// This field should not be exported because it is "read only".
var	trainingData = &TrainingData{}

// lbphParameters struct stores the LBPH parameters.
// It is not a pointer, so it will never be nil.
// This field should not be exported because the user cannot change
// the LBPH parameters after training the algorithm. To change the
// parameters we need to call Init that will "reset" the training data.
var	lbphParameters = Parameters{}

// The metric used to compare the histograms in the Predict step.
var Metric string

// init define the default state of some variables.
// It will set the default parameters for the LBPH,
// set the trainingData to nil and define the default
// metric (in this case ChiSquare).
func init() {
	// Define the default LBPH parameters.
	lbphParameters = Parameters{
		Radius:    1,
		Neighbors: 8,
		GridX:     8,
		GridY:     8,
	}

	// As the trainData is a pointer, the initial state can be nil.
	trainingData = nil

	// Use the ChiSquare as the default metric.
	Metric = metric.ChiSquare
}

// Init function is used to set the LBPH parameters based on the Parameters structure.
// It is needed to set the default parameters if something is wrong and
// to reset the trainingData when new parameters are defined.
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

	// Set the LBPH parameters
	lbphParameters = parameters

	// Every time the Init function is called the training data will be
	// reset, so the user needs to train the algorithm again.
	trainingData = nil
}

// GetTrainingData is used to get the trainingData struct.
// The user can use it to access the images, labels and histograms.
func GetTrainingData() TrainingData {
	// Returns the data structure pointed by trainData.
	return *trainingData
}

// checkImagesSizes function is used to check if all images have the same size.
func checkImagesSizes(images []image.Image) error {
	// Check if the slice is empty
	if len(images) == 0 {
		return errors.New("The images slice is empty")
	}
	// Check if the first image is nil
	if images[0] == nil {
		return errors.New("At least one image in the slice is nil")
	}

	// Get the image size from the first image
	defaultWidth, defaultHeight := lbp.GetImageSize(images[0])

	// Check if the size is valid
	// This condition should never happen because
	// we already tested if the image was nil
	if defaultWidth <= 0 || defaultHeight <= 0 {
		return errors.New("At least one image have an invalid size")
	}

	// Check each image in the slice
	for index := 0; index < len(images); index++ {
		// Check if the current image is nil
		if images[index] == nil {
			return errors.New("At least one image in the slice is nil")
		}

		// Get the size from the current image
		width, height := lbp.GetImageSize(images[index])

		// Check if all images have the same size
		if width != defaultWidth || height != defaultHeight {
			return errors.New("One or more images have different sizes")
		}
	}
	// No error has occurred, return nil
	return nil
}

// Train function is used for training the LBPH algorithm based on the
// images and labels passed by parameter. It basically checks the input
// data, calculates the LBP operation and gets the histogram of each image.
func Train(images []image.Image, labels []string) error {
	// Clear the data structure
	trainingData = nil

	// Check if the slices are not empty.
	if len(images) == 0 || len(labels) == 0 {
		return errors.New("At least one of the slices is empty")
	}

	// Check if the images and labels slices have the same size.
	if len(images) != len(labels) {
		return errors.New("The slices have different sizes")
	}

	// Call the CheckImagesSizes from the common package.
	// It will check if all images have the same size.
	err := checkImagesSizes(images)
	if err != nil {
		return err
	}

	// Calculates the LBP operation and gets the histograms for each image.
	var histograms [][]float64
	for index := 0; index < len(images); index++ {
		// Calculate the LBP operation for the current image.
		pixels, err := lbp.Calculate(images[index], lbphParameters.Radius, lbphParameters.Neighbors)
		if err != nil {
			return err
		}

		// Get the histogram from the current image.
		hist, err := histogram.Calculate(pixels, lbphParameters.GridX, lbphParameters.GridY)
		if err != nil {
			return err
		}

		// Store the histogram in the 'matrix' (slice of slice).
		histograms = append(histograms, hist)
	}

	// Store the current data that we are working on.
	trainingData = &TrainingData{
		Images:     images,
		Labels:     labels,
		Histograms: histograms,
	}

	// Everything is ok, return nil.
	return nil
}

// Predict function is used to find the closest image based on the images used in the training step.
func Predict(img image.Image) (string, float64, error) {

	// Check if we have data in the trainingData struct.
	if trainingData == nil {
		return "", 0.0, errors.New("The algorithm was not trained yet")
	}

	// Check if the image passed by parameter is nil.
	if img == nil {
		return "", 0.0, errors.New("The image passed by parameter is nil")
	}

	// If we don't have histograms to compare, probably the Train function was
	// not called or has occurred an error and it was not correctly treated.
	if len(trainingData.Histograms) == 0 {
		return "", 0.0, errors.New("There are no histograms in the trainData")
	}

	// Calculate the LBP operation.
	pixels, err := lbp.Calculate(img, lbphParameters.Radius, lbphParameters.Neighbors)
	if err != nil {
		return "", 0.0, err
	}

	// Calculate the histogram for the image.
	hist, err := histogram.Calculate(pixels, lbphParameters.GridX, lbphParameters.GridY)
	if err != nil {
		return "", 0.0, err
	}

	// Search for the closest histogram based on the histograms calculated in the training step.
	minMaxConfidence, err := histogram.Compare(hist, trainingData.Histograms[0], Metric)
	if err != nil {
		return "", 0.0, err
	}

	minIndex := 0
	for index := 1; index < len(trainingData.Histograms); index++ {
		// Calculate the confidence from the current histogram.
		confidence, err := histogram.Compare(hist, trainingData.Histograms[index], Metric)
		if err != nil {
			return "", 0.0, err
		}

		if Metric == metric.Intersection || Metric == metric.NormalizedIntersection {
			if confidence > minMaxConfidence {
				minMaxConfidence = confidence
				minIndex = index
			}
		} else {
			// If it is closer, save the minConfidence and the index.
			if confidence < minMaxConfidence {
				minMaxConfidence = confidence
				minIndex = index
			}
		}
	}

	// Return the label corresponding to the closest histogram,
	// the confidence (minConfidence) and the error (nil).
	return trainingData.Labels[minIndex], minMaxConfidence, nil
}
