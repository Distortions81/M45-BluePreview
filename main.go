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

	mapimage := image.NewRGBA(image.Rect(0, 0, 1024, 1024))
	newimage := image.NewRGBA(image.Rect(0, 0, 1024, 1024))

	var c uint8 = 0
	size := 16

	for y := 0; y < 1024; y++ {
		if y%size == 0 {
			if c == 0 {
				c = 64
			} else {
				c = 0
			}
		}
		for x := 0; x < 1024; x++ {

			if x%size == 0 {
				if c == 0 {
					c = 64
				} else {
					c = 0
				}
			}

			//If map has pixel here, draw it, otherwise draw BG
			if mapimage.At(x, y) != (color.RGBA{0, 0, 0, 0}) {
				newimage.Set(x, y, mapimage.At(x, y))
			} else {
				newimage.Set(x, y, color.RGBA{c, c, c, 255})
			}
		}
	}

	output, _ := os.Create(outputFile)
	if png.Encode(output, newimage) != nil {
		panic("Failed to write image")
	}

	output.Close()
	fmt.Println("Wrote image to", outputFile)

}
