# Local Binary Patterns Histograms (LBPH)

[![GoDoc](https://godoc.org/github.com/kelvins/lbph?status.svg)](https://godoc.org/github.com/kelvins/lbph)
[![License: MIT](https://img.shields.io/badge/License-MIT-brightgreen.svg)](LICENSE)

# Summary

1. [Introduction](#introduction)
2. [Step-by-Step](#step-by-step)  
2.1. [Important Notes](#important-notes)
3. [I/O](#io)  
3.1. [Input](#input)  
3.2. [Output](#output)
4. [Usage](#usage)  
4.1. [Installation](#installation)  
4.2. [Usage Example](#usage-example)  
4.3. [Parameters](#parameters)  
5. [References](#references)
6. [How to contribute](#how-to-contribute)  
6.1. [Contributing](#contributing)

## Introduction

Local Binary Patterns (LBP) is a type of visual descriptor used for classification in computer vision. LBP was first described in 1994 and has since been found to be a powerful feature for texture classification. It has further been determined that when LBP is combined with the Histogram of oriented gradients (HOG) descriptor, it improves the detection performance considerably on some datasets.

As LBP is a visual descriptor it can also be used for face recognition tasks, as can be seen in the following Step-by-Step explanation.

## Step-by-Step

In this section, it is shown a step-by-step explanation of the LBPH algorithm:

1. First of all, we need to train the algorithm. To do that we just need to call the `Train` function passing a slice of images and a slice of labels.
2. The `Train` function will check if all images are in grayscale and have the same size.
3. Then, the `Train` function will apply the basic LBP operation by changing each pixel based on its `8` neighbors using a default radius of `1`. The basic LBP operation can be seen in the following image:

![LBP operation](http://i.imgur.com/G4PqJPe.png)

4. After applying the LBP operation we extract the histograms of the grayscale image based on the number of grids (X and Y) passed by parameter. After extracting the histogram of each region, we concatenate all histograms and create a new one that will be used to represent the image.

![Histograms](http://i.imgur.com/3BGk130.png)

5. The images, labels, and histograms are stored in a data structure so we can compare all of it to a new image in the `Predict` function.
6. Now, the algorithm is already trained and we can Predict a new image.
7. To predict a new image we just need to call the `Predict` function passing the image as parameter. The `Predict` function will extract the histogram from the new image and will return the label and distance corresponding to the closest histogram if no error has occurred (e.g. the image is not in grayscale, or image does not have the same size).
8. It uses the normalized euclidean distance to calculate the similarity of the histograms. We can assume that the distance returned by the `Predict` function is the confidence and assume that the algorithm result is correct based on this confidence. The closer to zero is the distance, the greater is the confidence.

### Important Notes

- The similarity between two histograms is calculated using the normalized euclidean distance presented in the following formula:

![Euclidean Distance](http://i.imgur.com/liBbl6u.gif)

- The current LBPH implementation uses a fixed `radius` of `1` and a fixed number of `neighbors` equal to `8`. In the future, we intend to provide an option to the user set these values as parameters.

## I/O

### Input

All input images (for train and test) must have the same size. Different of OpenCV, the images don't need to be in grayscale, because each pixel is automatically converted to grayscale in the LBP process using the following [formula](https://en.wikipedia.org/wiki/Grayscale#Luma_coding_in_video_systems):

```
Y = (0.299 * RED) + (0.587 * GREEN) + (0.114 * BLUE)
```

### Output

The Predict function returns 3 values:

* **label**: The label corresponding to the predicted image.
* **distance**: The distance between the histograms from the input test image and the matched image.
* **err**: Some error that has occurred in the Predict step. If no error occurs it will returns nil.

Using the label you can check if the algorithm has correctly predicted the image. In a real world application, it is not feasible to manually verify all images, so we can use the distance to infer if the algorithm has predicted correctly or not.

## Usage

### Installation

```
$ go get github.com/kelvins/lbph
```

### Usage Example

Usage example:

``` go

package main

import (
	"os"
	"fmt"
	"image"

	"github.com/kelvins/lbph"
	"github.com/kelvins/lbph/common"
)

var trainImagesPaths []string
var trainLabels []string
var trainImages []image.Image

var testImagesPaths []string
var testLabels []string
var testImages []image.Image

func init() {
	trainImagesPaths = append(trainImagesPaths, "./dataset/train/1.png")
	trainImagesPaths = append(trainImagesPaths, "./dataset/train/2.png")
	trainImagesPaths = append(trainImagesPaths, "./dataset/train/3.png")

	trainLabels = append(trainLabels, "rocks")
	trainLabels = append(trainLabels, "grass")
	trainLabels = append(trainLabels, "wood")

	trainImages = loadImages(trainImagesPaths)

	testImagesPaths = append(testImagesPaths, "./dataset/test/1.png")
	testImagesPaths = append(testImagesPaths, "./dataset/test/2.png")
	testImagesPaths = append(testImagesPaths, "./dataset/test/3.png")

	testLabels = append(testLabels, "wood")
	testLabels = append(testLabels, "rocks")
	testLabels = append(testLabels, "grass")

	testImages = loadImages(testImagesPaths)
}

func main() {

	parameters := lbph.Parameters{
		Radius:    1,
		Neighbors: 8,
		GridX:     8,
		GridY:     8,
	}

	lbph.Init(parameters)

	err := lbph.Train(trainImages, trainLabels)
	checkError(err)

	for index := 0; index < len(testImages); index++ {
		label, distance, err := lbph.Predict(testImages[index])
		checkError(err)

		if label == testLabels[index] {
			fmt.Println("Image correctly predicted")
		} else {
			fmt.Println("Image wrongly predicted")
		}

		fmt.Printf("Predicted as %s expected %s\n", label, testLabels[index])
		fmt.Printf("Distance: %f\n\n", distance)
	}
}

func loadImages(paths []string) []image.Image {
	var images []image.Image

	for index := 0; index < len(paths); index++ {
		img, err := common.LoadImage(paths[index])
		checkError(err)
		images = append(images, img)
	}

	return images
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

```

### Parameters

* **Radius**: The radius used for building the Circular Local Binary Pattern. Default value is 1.

* **Neighbors**: The number of sample points to build a Circular Local Binary Pattern from. Keep in mind: the more sample points you include, the higher the computational cost. Default value is 8.

* **GridX**: The number of cells in the horizontal direction. The more cells, the finer the grid, the higher the dimensionality of the resulting feature vector. Default value is 8.

* **GridY**: The number of cells in the vertical direction. The more cells, the finer the grid, the higher the dimensionality of the resulting feature vector. Default value is 8.

## References

* Ahonen, Timo, Abdenour Hadid, and Matti Pietikäinen. "Face recognition with local binary patterns." Computer vision-eccv 2004 (2004): 469-481. Link: https://link.springer.com/chapter/10.1007/978-3-540-24670-1_36

* Face Recognizer module. Open Source Computer Vision Library (OpenCV) Documentation. Version 3.0. Link: http://docs.opencv.org/3.0-beta/modules/face/doc/facerec/facerec_api.html

* Local binary patterns. Wikipedia. Link: https://en.wikipedia.org/wiki/Local_binary_patterns

## How to contribute

Feel free to contribute by commenting, suggesting, creating [issues](https://github.com/kelvins/lbph/issues) or sending pull requests. Any help is welcome.

### Contributing

1. Create an issue (optional)
2. Fork the repo to your Github account
3. Clone the project to your local machine
4. Make your changes
5. Commit your changes (`git commit -am 'Some cool feature'`)
6. Push to the branch (`git push origin master`)
7. Create a new Pull Request
