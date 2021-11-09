package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

var outputFile = "output.png"

func main() {

	image := image.NewRGBA(image.Rect(0, 0, 1024, 1024))
	for y := 0; y < 1024; y++ {
		for x := 0; x < 1024; x++ {
			image.Set(x, y, color.RGBA{uint8(x), uint8(y), 0, 255})
		}
	}

	output, _ := os.Create(outputFile)
	if png.Encode(output, image) != nil {
		panic("Failed to write image")
	}

	output.Close()
	fmt.Println("Wrote image to", outputFile)

}
