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
	"log"
	"math"
	"os"

	"./data"
)

var outputFile = "output.png"
var inputFile = "input.txt"

var scaleup float64 = 1.0

const checkersize = 32
const maxmag = 32
const idealsize = 512
const tlSpace = 2

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

	//Normalize cordinates
	maxx := 0.0
	maxy := 0.0

	//Find max and min
	for _, v := range newbp.BluePrint.Entities {
		if v.Position.X > maxx {
			maxx = v.Position.X
		}
		if v.Position.Y > maxy {
			maxy = v.Position.Y
		}
	}
	minx := maxx
	miny := maxy
	for _, v := range newbp.BluePrint.Entities {
		if v.Position.X < minx {
			minx = v.Position.X
		}
		if v.Position.Y < miny {
			miny = v.Position.Y
		}
	}

	//Size image
	xsize := int(maxx - minx)
	ysize := int(maxy - miny)

	buf := fmt.Sprintf("BP size: %d x %d", int(xsize), int(ysize))
	log.Println(buf)

	maxsize := 0
	if xsize > ysize {
		maxsize = int(xsize)
	} else {
		maxsize = int(ysize)
	}

	for a := 1; a <= maxmag; a = a + 1 {
		if maxsize*a < idealsize {
			scaleup = float64(a)
		}
	}

	imx := int(xsize+tlSpace*2) * int(scaleup)
	imy := int(ysize+tlSpace*2) * int(scaleup)

	buf = fmt.Sprintf("Image size: %d x %d (%dX mag)", imx, imy, int(scaleup))
	log.Println(buf)

	mapimage := image.NewRGBA(image.Rect(0, 0, imx, imy))
	newimage := image.NewRGBA(image.Rect(0, 0, imx, imy))

	//Draw map, scaled
	var objx, objy, x, y, xs, ys, xo, yo, ix, iy float64
	for _, v := range newbp.BluePrint.Entities {

		//Offset position
		objx = v.Position.X - float64(minx)
		objy = v.Position.Y - float64(miny)
		//Get color for item
		item := findItem(v.Name)

		if v.Direction == 2 || v.Direction == 6 {
			ix = item.Y
			iy = item.X
		} else {
			ix = item.X
			iy = item.Y
		}

		if ix > 1 {
			objx = objx - (ix / 2.0) + 0.5
		}
		if iy > 1 {
			objy = objy - (iy / 2.0) + 0.5
		}

		x = ((objx * scaleup) + (tlSpace * scaleup))
		y = ((objy * scaleup) + (tlSpace * scaleup))

		xs = ix * scaleup
		ys = iy * scaleup
		for xo = 0.0; xo < xs; xo = xo + 1 {
			for yo = 0.0; yo < ys; yo = yo + 1 {
				mapimage.Set(int(x+xo), int(y+yo), item.Color)
			}
		}

		//count = count + 1
	}
	//log.Println(count)

	//Draw checkerboard background, draw map on top
	var c uint8 = 0
	csize := int(math.Round(checkersize * scaleup))

	for y := 0; y < imy; y++ {
		for x := 0; x < imx; x++ {
			yoff := y - ((tlSpace) * int(scaleup))
			xoff := x - ((tlSpace + 1) * int(scaleup))

			if xoff%csize != 0 && yoff%csize != 0 {
				c = 0
			} else {
				c = 16
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
