package main

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"os"
)

var outputFile = "output.png"
var inputFile = "input.txt"

func main() {

	fmt.Println("Reading input file...")

	// Open the input file
	input, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}

	data, err := base64.StdEncoding.DecodeString(string(input[1:]))
	if err != nil {
		fmt.Println("Error decoding input file:", err)
		return
	}
	//fmt.Println(data)

	r, err := zlib.NewReader(bytes.NewReader(data))
	if err != nil {
		panic(err)
	}
	enflated, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(enflated))

	type Xy struct {
		X float32
		Y float32
	}
	type Ent struct {
		Entity_number int
		Name          string
		Position      Xy
		Direction     int
	}
	type BpData struct {
		Entities []Ent
		Item     string
		Label    string
		Version  int64
	}
	type Bp struct {
		BluePrint BpData
	}

	newbp := Bp{}

	err = json.Unmarshal(enflated, &newbp)
	if err != nil {
		panic(err)
	}

	if 1 == 2 {
		str, err := json.Marshal(newbp)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(str))
	}

	mapimage := image.NewRGBA(image.Rect(0, 0, 1024, 1024))
	newimage := image.NewRGBA(image.Rect(0, 0, 1024, 1024))

	count := 0
	for _, v := range newbp.BluePrint.Entities {
		x := int(v.Position.X)
		y := int(v.Position.Y)
		mapimage.Set(x, y, color.RGBA{255, 255, 255, 255})

		count = count + 1
	}
	fmt.Println(count)

	var c uint8 = 0
	size := 16

	for y := 0; y < 1024; y++ {
		if y%size == 0 {
			if c == 0 {
				c = 8
			} else {
				c = 0
			}
		}
		for x := 0; x < 1024; x++ {

			if x%size == 0 {
				if c == 0 {
					c = 8
				} else {
					c = 0
				}
			}

			//If map has pixel here, draw it, otherwise draw BG
			if mapimage.At(x, y) != (color.RGBA{0, 0, 0, 0}) {
				newimage.Set(x, y, mapimage.At(x, y))
			} else {
				newimage.Set(x, y, color.RGBA{c, 0, 0, 255})
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
