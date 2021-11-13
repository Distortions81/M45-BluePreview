package main

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"encoding/json"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"math"
	"os"

	"./data"
)

var outputFile = "output.png"
var inputFile = "input.txt"

const isx = 4096
const isy = 4096
const scaleup = 4.0
const checkersize = 32.0
const tlSpace = 4.0

func findItem(itemName string) data.Item {
	for _, i := range data.ItemData {
		if i.Name == itemName {
			//log.Println("Found item:", i.Name)
			return i
		}
	}

	log.Println("Item not found:", itemName)
	return data.Item{"Default", 1, 1, color.RGBA{1, 0, 1, 1}}
}

func main() {

	log.SetFlags(log.Lmicroseconds | log.Lshortfile)
	log.Println("Reading input file...")

	// Open the input file
	input, err := os.ReadFile(inputFile)
	if err != nil {
		log.Println("Error opening input file:", err)
		return
	}

	log.Println("Base64 decoding...")
	data, err := base64.StdEncoding.DecodeString(string(input[1:]))
	if err != nil {
		log.Println("Error decoding input file:", err)
		return
	}
	//log.Println(data)

	log.Println("Decompressing...")
	r, err := zlib.NewReader(bytes.NewReader(data))
	if err != nil {
		panic(err)
	}
	enflated, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}
	//log.Println(string(enflated))

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

	log.Println("Unmarshaling JSON...")
	err = json.Unmarshal(enflated, &newbp)
	if err != nil {
		panic(err)
	}

	if 1 == 2 {
		str, err := json.Marshal(newbp)
		if err != nil {
			panic(err)
		}
		log.Println(string(str))
	}

	log.Println("Drawing image...")
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
		objx := int(v.Position.X - minx + 0.5)
		objy := int(v.Position.Y - miny + 0.5)

		item := findItem(v.Name)

		if item.X > 1 {
			objx = objx - 1
		}
		if item.Y > 1 {
			objy = objy - 1
		}

		x := int(objx*scaleup) + ((tlSpace + 2) * scaleup)
		y := int(objy*scaleup) + (tlSpace * scaleup)
		xs := int(item.X * scaleup)
		ys := int(item.Y * scaleup)

		for xo := 0; xo < xs; xo = xo + 1 {
			for yo := 0; yo < ys; yo = yo + 1 {
				mapimage.Set(x+xo, y+yo, item.Color)
			}
		}

		//count = count + 1
	}
	//log.Println(count)

	var c uint8 = 0
	size := int(math.Round(checkersize * scaleup))

	for y := 0; y < isy; y++ {
		if (y-(tlSpace*scaleup))%size == 0 {
			if c == 0 {
				c = 1
			} else {
				c = 0
			}
		}
		for x := 0; x < isx; x++ {

			if (x-((tlSpace+1)*scaleup))%size == 0 {
				if c == 0 {
					c = 16
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
	log.Println("Wrote image to", outputFile)

}
