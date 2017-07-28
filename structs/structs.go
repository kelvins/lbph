package structs

import "image"

// Store the input data (images and labels) and the calculated histogram
type Data struct {
	Images     []image.Image
	Labels     []string
	Histograms [][256]int64
}

// Structure used to pass the LBPH parameters
type LBPHParameters struct {
	Radius    uint8
	Neighbors uint8
	GridX     uint8
	GridY     uint8
}