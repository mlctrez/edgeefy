package main

import (
	"github.com/mlctrez/edgeefy"
	"image"
	"image/jpeg"
	"image/png"
	"os"
)

func main() {
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

	classic, err := os.Open("samples/classic.jpg")
	if err != nil {
		panic(err)
	}

	img, _, err := image.Decode(classic)
	if err != nil {
		panic(err)
	}

	pixels, err := edgeefy.GrayPixelsFrommImage(img)
	if err != nil {
		panic(err)
	}

	pixels = edgeefy.CannyEdgeDetect(pixels, true, .6, .2)

	grayImage := edgeefy.GrayImageFromGrayPixels(pixels)

	outFile, err := os.Create("output.png")
	if err != nil {
		panic(err)
	}

	err = png.Encode(outFile, grayImage)
	if err != nil {
		panic(err)
	}

	outFile.Close()




}
