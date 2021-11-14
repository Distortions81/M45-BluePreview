package main

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"math"
	"os"
	"time"

	"./data"
)

const version = "023"
const build = "111420210146p"

var scaleup float64 = 1.0 //Minimum scale

const checkersize = 32 //Checkerboard BG size, chunks are 32
const maxmag = 32      //Max magnification for small bps
const idealsize = 1024 //Magnify to this size
const tlSpace = 2      //Margins

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

type SignalData struct {
	Type string
	Name string
}

type Icn struct {
	Signal SignalData
	Index  int
}

type BpData struct {
	Entities []Ent
	Icons    []Icn
	Item     string
	Label    string
	Version  int64
}
type Bp struct {
	BluePrint BpData
}

//Look through data.go and find item type
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

func bAddr(b *bool) bool {
	boolVar := *b

	if boolVar {
		return true
	}

	return false

}

func strAddr(str *string) string {
	newString := string(*str)

	return newString
}

func main() {

	inputFileP := flag.String("file", "bp.txt", "blueprint filename")
	outputNameP := flag.String("name", "bp", "bp name")
	stdinModeP := flag.Bool("stdin", false, "look for bp data on stdin")
	jsonOutP := flag.Bool("json", false, "also output json data")
	showVersionP := flag.Bool("version", false, "display version")
	flag.Parse()

	outputName := strAddr(outputNameP)
	inputFile := strAddr(inputFileP)
	stdinMode := bAddr(stdinModeP)
	jsonOut := bAddr(jsonOutP)
	showVersion := bAddr(showVersionP)

	if showVersion {
		fmt.Println("M45-Science FactMap: " + "v" + version + "-" + build)
		os.Exit(0)
	}

	if jsonOut || stdinMode {
		//do stuff
	}

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

	newbp := Bp{}

	log.Println("Unmarshaling JSON...")
	err = json.Unmarshal(enflated, &newbp)
	if err != nil {
		panic(err)
	}

	//Re-Encode json and pretty print it
	//Disabled
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

	//Find largest side
	maxsize := 0
	if xsize > ysize {
		maxsize = int(xsize)
	} else {
		maxsize = int(ysize)
	}

	//Scale up to max magnifcation, or ideal size
	for a := 1; a <= maxmag; a = a + 1 {
		if maxsize*a < idealsize {
			scaleup = float64(a)
		}
	}

	//Offset x/y with margins and scale it up to the magnified size
	imx := int(xsize+tlSpace*2) * int(scaleup)
	imy := int(ysize+tlSpace*2) * int(scaleup)

	buf = fmt.Sprintf("Image size: %d x %d (%dX mag)", imx, imy, int(scaleup))
	log.Println(buf)

	//Allocate image and bg
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

		//Handle item rotation
		if v.Direction == 2 || v.Direction == 6 { //east/west
			ix = item.Y
			iy = item.X
		} else { //north/south, etc
			ix = item.X
			iy = item.Y
		}

		//If item side is larger than 1, deal with recentering it
		if ix > 1 {
			objx = objx - (ix / 2.0) + 0.5
		}
		if iy > 1 {
			objy = objy - (iy / 2.0) + 0.5
		}

		//Item x/y and margin offsets
		x = ((objx * scaleup) + (tlSpace * scaleup))
		y = ((objy * scaleup) + (tlSpace * scaleup))

		//Item size, scaled
		xs = ix * scaleup
		ys = iy * scaleup

		//Draw item, magnified
		for xo = 0.0; xo < xs; xo = xo + 1 {
			for yo = 0.0; yo < ys; yo = yo + 1 {
				mapimage.Set(int(x+xo), int(y+yo), item.Color)
			}
		}

		//count = count + 1
	}
	//log.Println(count)

	//Draw checkerboard background
	var c uint8
	csize := int(math.Round(checkersize * scaleup))

	for y := 0; y < imy; y++ {
		for x := 0; x < imx; x++ {

			//x/y, margin offset, scaled
			yoff := y - ((tlSpace) * int(scaleup))
			xoff := x - ((tlSpace + 1) * int(scaleup))

			//Only draw boarders
			if xoff%csize != 0 && yoff%csize != 0 {
				c = 32
			} else {
				c = 0
			}

			//If map has pixel here, draw it, otherwise draw BG
			if mapimage.At(x, y) != (color.RGBA{0, 0, 0, 0}) {
				newimage.Set(x, y, mapimage.At(x, y))
			} else {
				newimage.Set(x, y, color.RGBA{c, c, c, 255})
			}
		}
	}

	//If blueprint has a name, use it
	t := time.Now()
	bpname := ""
	if newbp.BluePrint.Label != "" {
		bpname = newbp.BluePrint.Label + "-"
	}
	cTime := t.UnixNano()

	//Write json file
	fileName := fmt.Sprintf("%s-%s%d.json", outputName, bpname, cTime)
	err = os.WriteFile(fileName, enflated, 0644)
	if err != nil {
		panic(err)
	}
	log.Println("Wrote json to", fileName)

	//Write the png file
	fileName = fmt.Sprintf("%s-%d.png", bpname, cTime)
	output, _ := os.Create(fileName)
	if png.Encode(output, newimage) != nil {
		panic("Failed to write image")
	}
	log.Println("Wrote image to", fileName)
	output.Close()

}
