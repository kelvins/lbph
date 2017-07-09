# Local Binary Patterns Histograms (LBPH)

[![GoDoc](https://godoc.org/github.com/kelvins/lbph?status.svg)](https://godoc.org/github.com/kelvins/lbph)
[![License: MIT](https://img.shields.io/badge/License-MIT-brightgreen.svg)](LICENSE)

Local Binary Patterns (LBP) is a type of visual descriptor used for classification in computer vision. LBP was first described in 1994 and has since been found to be a powerful feature for texture classification. It has further been determined that when LBP is combined with the Histogram of oriented gradients (HOG) descriptor, it improves the detection performance considerably on some datasets.

## Installation

```
go get github.com/kelvins/lbph
```

## Step-by-Step

In this section, it is shown a step-by-step explanation of the LBPH algorithm:

1. First of all, we need to train the algorithm. To do that we just need to call the `Train` function passing a slice of images and a slice of labels.
2. The `Train` function will check if all images are in grayscale and have the same size.

![LBP operation](http://i.imgur.com/1IEVqnZ.png)

## Important Notes

- The similarity between two histograms is calculated using the normalized euclidean distance presented in the following formula:

![Euclidean Distance](http://i.imgur.com/liBbl6u.gif)

- The current LBPH implementation uses a fixed `radius` of `1` and a fixed number of `neighbors` equal to `8`. In the future, we intend to provide an option to the user set this values as parameters.

## Usage

``` go
package main

import (
	"image"
	"os"
	"fmt"
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

func main() {

	var paths []string
	paths = append(paths, "./dataset/train/1.png")
	paths = append(paths, "./dataset/train/2.png")
	paths = append(paths, "./dataset/train/3.png")

	var labels []string
	labels = append(labels, "1")
	labels = append(labels, "2")
	labels = append(labels, "3")

	var images []image.Image

	for index := 0; index < len(paths); index++ {
		img, err := loadImage(paths[index])
		if err != nil {
			t.Error(err)
		}
		images = append(images, img)
	}

	err := Train(images, labels)
	if err != nil {
		t.Error(err)
	}

	img, err := loadImage("./dataset/test/1.png")
	if err != nil {
		t.Error(err)
	}

	lbl, dist, err := Predict(img)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("Predicted as subject", lbl)
}
```

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
