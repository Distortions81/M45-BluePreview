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
	"math"
	"os"

	"./data"
)

var outputFile = "output.png"
var inputFile = "input.txt"

const isx = 256
const isy = 256
const scaleup = 4.0
const checkersize = 32.0

func findItem(itemName string) data.Item {
	for _, i := range data.ItemData {
		if i.Name == itemName {
			//fmt.Println("Found item:", i.Name)
			return i
		}
	}

	fmt.Println("Item not found:", itemName)
	return data.Item{"Default", 1, 1, color.RGBA{1, 0, 1, 1}}
}

func main() {

	fmt.Println("Reading input file...")

	// Open the input file
	input, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}

	fmt.Println("Base64 decoding...")
	data, err := base64.StdEncoding.DecodeString(string(input[1:]))
	if err != nil {
		fmt.Println("Error decoding input file:", err)
		return
	}
	//fmt.Println(data)

	fmt.Println("Decompressing...")
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
		X float64
		Y float64
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

	fmt.Println("Unmarshaling JSON...")
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

	fmt.Println("Drawing image...")
	mapimage := image.NewRGBA(image.Rect(0, 0, isx, isy))
	newimage := image.NewRGBA(image.Rect(0, 0, isx, isy))

	//Normalize cordinates
	minx := 0.0
	miny := 0.0

	for _, v := range newbp.BluePrint.Entities {
		if v.Position.X > minx {
			minx = v.Position.X
		}
		if v.Position.Y > miny {
			miny = v.Position.Y
		}
	}

	for _, v := range newbp.BluePrint.Entities {
		if v.Position.X < minx {
			minx = v.Position.X
		}
		if v.Position.Y < miny {
			miny = v.Position.Y
		}
	}

	for _, v := range newbp.BluePrint.Entities {
		objx := int(v.Position.X-minx+0.5) + 2
		objy := int(v.Position.Y-miny+0.5) + 2

		item := findItem(v.Name)

		if item.X > 1 {
			objx = objx - 1
		}
		if item.Y > 1 {
			objy = objy - 1
		}

		x := int(objx * scaleup)
		y := int(objy * scaleup)
		xs := int(item.X * scaleup)
		ys := int(item.Y * scaleup)

		for xo := 0; xo < xs; xo = xo + 1 {
			for yo := 0; yo < ys; yo = yo + 1 {
				mapimage.Set(x+xo, y+yo, item.Color)
			}
		}

		//count = count + 1
	}
	//fmt.Println(count)

	var c uint8 = 0
	size := int(math.Round(checkersize * scaleup))

	for y := 0; y < isy; y++ {
		if y%size == 0 {
			if c == 0 {
				c = 32
			} else {
				c = 0
			}
		}
		for x := 0; x < isx; x++ {

			if x%size == 0 {
				if c == 0 {
					c = 32
				} else {
					c = 0
				}
			}

			//If map has pixel here, draw it, otherwise draw BG
			if mapimage.At(x, y) != (color.RGBA{0, 0, 0, 0}) {
				newimage.Set(x, y, mapimage.At(x, y))
			} else {
				newimage.Set(x, y, color.RGBA{0, c, c, 255})
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
