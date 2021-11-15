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
	"path/filepath"
	"time"

	"./data"
)

const version = "025"
const build = "11152021-1255"

var scaleup float64 = 1.0 //Minimum scale

const checkersize = 32 //Checkerboard BG size, chunks are 32
const maxmag = 32      //Max magnification for small bps
const idealsize = 1024 //Magnify to this size
const tlSpace = 2      //Margins

//Max sizes
const maxInput = 10 * 1024 * 1024 //10MB
const maxJson = 100 * 1024 * 1024 //100MB
const maxImage = 1024 * 16        //16k

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
			return i
		}
	}

	log.Println("ERROR: Item not found:", itemName)
	return data.Item{Name: "Default", X: 1, Y: 1, Color: color.RGBA{1, 0, 1, 1}}
}

//bool pointer to bool
func bAddr(b *bool) bool {
	boolVar := *b

	if boolVar {
		return boolVar
	}

	return false

}

//String pointer to string
func strAddr(str *string) string {
	newString := string(*str)

	return newString
}

func main() {

	log.SetFlags(log.Lmicroseconds | log.Lshortfile)

	//Lanch params
	inputFileP := flag.String("file", "bp.txt", "filename of input")
	outputNameP := flag.String("name", "bp", "blueprint name")
	stdinModeP := flag.Bool("stdin", false, "look for bp data on stdin")
	jsonOutP := flag.Bool("json", false, "also output json data")
	showVersionP := flag.Bool("version", false, "display version")
	showHelpP := flag.Bool("help", false, "display help")
	showTimeP := flag.Bool("time", false, "put time (unix nano) in filenames")
	showVerboseP := flag.Bool("verbose", false, "verbose output (progress)")
	showDebugP := flag.Bool("debug", false, "debug output")
	showCheckerP := flag.Bool("checker", true, "show checkerboard background")
	flag.Parse()

	outputName := strAddr(outputNameP)
	inputFile := strAddr(inputFileP)
	stdinMode := bAddr(stdinModeP)
	jsonOut := bAddr(jsonOutP)
	showVersion := bAddr(showVersionP)
	showHelp := bAddr(showHelpP)
	showTime := bAddr(showTimeP)
	showVerbose := bAddr(showVerboseP)
	showDebug := bAddr(showDebugP)
	showChecker := bAddr(showCheckerP)

	//Debug mode also enables verbose
	if showDebug {
		showVerbose = true
	}

	//Launch param help
	if showHelp {
		execName := filepath.Base(os.Args[0])
		log.Println("Usage: " + execName + " [options]")
		log.Println("Options:")
		flag.PrintDefaults()
		os.Exit(0)
	}

	//Show version and build
	if showVersion {
		log.Println("M45-Science FactMap: " + "v" + version + "-" + build)
		os.Exit(0)
	}

	log.SetFlags(log.Lmicroseconds | log.Lshortfile)

	var input []byte
	var err error
	if stdinMode {
		//TODO
		//read standard input, and don't attempt to read input file.
		log.Println("Not implemented yet.")
		os.Exit(1)
	} else {

		if showVerbose {
			log.Println("Reading input file...")
		}

		// Open the input file
		input, err = os.ReadFile(inputFile)
		if err != nil {
			log.Println("Error opening input file: "+inputFile+"\n", err)
			os.Exit(1)
		}
	}

	//Max input
	if len(input) > maxInput {
		log.Println("Input data too large.")
		os.Exit(1)
	}

	if showVerbose {
		log.Println("Decoding...")
	}
	data, err := base64.StdEncoding.DecodeString(string(input[1:]))
	if err != nil {
		log.Println("Error decoding input:", err)
		os.Exit(1)
	}

	if showVerbose {
		log.Println("Decompressing...")
	}

	r, err := zlib.NewReader(bytes.NewReader(data))
	if err != nil {
		panic(err)
	}
	enflated, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}

	//Max decompressed size
	if len(enflated) > maxJson {
		log.Println("Input data too large.")
		os.Exit(1)
	}

	newbp := Bp{}

	if showVerbose {
		log.Println("Unmarshaling JSON...")
	}
	err = json.Unmarshal(enflated, &newbp)
	if err != nil {
		panic(err)
	}

	if showVerbose {
		log.Println("Drawing image...")
	}
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

	if showVerbose {
		buf := fmt.Sprintf("BP size: %v x %v", int(xsize), int(ysize))
		log.Println(buf)
	}

	//Find largest side
	maxsize := 0
	if xsize > ysize {
		maxsize = int(xsize)
	} else {
		maxsize = int(ysize)
	}

	if xsize > maxImage {
		log.Println("Final image would be too large.")
		os.Exit(1)
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

	if showVerbose {
		buf := fmt.Sprintf("Image size: %v x %v (%vX mag)", imx, imy, int(scaleup))
		log.Println(buf)
	}

	//Allocate image and bg
	mapimage := image.NewRGBA(image.Rect(0, 0, imx, imy))
	newimage := image.NewRGBA(image.Rect(0, 0, imx, imy))

	//Draw map, scaled
	var objx, objy, x, y, xs, ys, xo, yo, ix, iy float64
	count := 0
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
		if v.Name == "offshore-pump" {
			if v.Direction == 2 { //east
				objx = objx + 1.0
			} else if v.Direction == 6 { //west
				objx = objx - 1.0
			} else if v.Direction == 0 { //north
				objy = objy - 1.0
			} else if v.Direction == 4 { //south
				objy = objy + 1.0
			}
		} else if v.Name == "straight-rail" {
			if v.Direction == 1 { //ne
				objy = objy - 0.5
				objx = objx + 0.5
			} else if v.Direction == 3 { //se
				objy = objy + 0.5
				objx = objx + 0.5
			} else if v.Direction == 5 { //sw
				objy = objy + 0.5
				objx = objx - 0.5
			} else if v.Direction == 7 { //nw
				objy = objy - 0.5
				objx = objx - 0.5
			}
		} else {
			if ix > 1 {
				objx = objx - (ix / 2.0) + 0.5
			}
			if iy > 1 {
				objy = objy - (iy / 2.0) + 0.5
			}
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
				r, g, b, a := (mapimage.At(int(x+xo), int(y+yo))).RGBA()
				if r == 0 && g == 0 && b == 0 && a == 0 {
					r, g, b, _ := item.Color.RGBA()
					mapimage.Set(int(x+xo), int(y+yo), color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), 255})
				} else {
					if showDebug {
						if showDebug {
							mapimage.Set(int(x+xo), int(y+yo), color.RGBA{254, 0, 0, 255})
						}
						log.Println("Error:", v.Name, "at", x+xo, ",", y+yo, "overdraw! Color:", r>>8, g>>8, b>>8)
					}
				}
			}
		}

		count = count + 1
	}
	if showDebug {
		buf := fmt.Sprintf("Item count: %v", count)
		log.Println(buf)
	}

	//Draw checkerboard background
	if showChecker {
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
	} else {
		newimage = mapimage
	}

	//If blueprint has a name, use it
	t := time.Now()
	bpname := ""
	if newbp.BluePrint.Label != "" {
		bpname = newbp.BluePrint.Label + "-"
	}

	cTime := ""
	if showTime {
		cTime = fmt.Sprintf("%v", t.UnixNano())
	}

	//Write json file
	if jsonOut {
		fileName := fmt.Sprintf("%v%v%v.json", outputName, bpname, cTime)
		err = os.WriteFile(fileName, enflated, 0644)
		if err != nil {
			panic(err)
		}
		log.Println("Wrote json to", fileName)
	}

	//Write the png file
	fileName := fmt.Sprintf("%v%v%v.png", outputName, bpname, cTime)
	output, _ := os.Create(fileName)
	if png.Encode(output, newimage) != nil {
		panic("Failed to write image")
	}
	log.Println("Wrote image to", fileName)
	output.Close()

}
