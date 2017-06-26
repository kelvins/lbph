package structs

import "image"

// Store the input data (images and labels) and the calculated histogram
type Data struct {
	Images     []image.Image
	Labels     []string
	Histograms [][256]int64
}
