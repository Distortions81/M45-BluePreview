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
	"os"
	"path/filepath"
	"time"

	"bluepreview/data"
)

const NORTH = 0
const NORTH_EAST = 1
const EAST = 2
const SOUTH_EAST = 3
const SOUTH = 4
const SOUTH_WEST = 5
const WEST = 6
const NORTH_WEST = 7

//Code written by CarlOtto81@gmail.com
//MPL-2.0 License
const version = "029"          //increment
const build = "11152021-1257p" //mmddyyyy-hhmm(mst)

var scaleup float64

const checkersize = 32 //Checkerboard BG size, chunks are 32
const maxmag = 32      //Max magnification for small bps
const idealsize = 1024 //Magnify to this size
const tlSpace = 2      //Margins

//Max sizes
const maxInput = 10 * 1024 * 1024 //10MB
const maxJson = 100 * 1024 * 1024 //100MB
const maxImage = 1024 * 16        //16k

type Xy struct {
	X float64 `json:"y"`
	Y float64 `json:"x"`
}
type Ent struct {
	Entity_number int    `json:"entity_number"`
	Name          string `json:"name"`
	Position      Xy     `json:"position"`
	Direction     int    `json:"direction"`
}

type SignalData struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

type Icn struct {
	Signal SignalData `json:"signaldata"`
	Index  int        `json:"index"`
}

type BluePrintData struct {
	Entities []Ent  `json:"entities"`
	Icons    []Icn  `json:"icons"`
	Item     string `json:"item"`
	Label    string `json:"label"`
	Version  int64  `json:"version"`
}
type BluePrintsData struct {
	Blueprint BluePrintData `json:"blueprint"`
}
type BluePrintBookData struct {
	Blueprints []BluePrintsData `json:"blueprints"`
}
type RootData struct {
	Blueprintbook BluePrintBookData `json:"blueprint_book"`
	Blueprint     BluePrintData     `json:"blueprint"`
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

//Int pointer to int
func iAddr(i *int) int {
	newInt := int(*i)

	return newInt
}

func fAddr(f *float64) float64 {
	newFloat := float64(*f)

	return newFloat
}

var outputName string
var inputFile string
var stdinMode bool
var jsonOut bool
var showVersion bool
var showHelp bool
var showTime bool
var showVerbose bool
var showDebug bool
var showChecker bool
var itemSpacing float64

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
	itemSpacingP := flag.Int("space", 1, "draw space around items when magnified")
	minMagP := flag.Float64("minmag", 2, "minimum magnification level")
	flag.Parse()

	outputName = strAddr(outputNameP)
	inputFile = strAddr(inputFileP)
	stdinMode = bAddr(stdinModeP)
	jsonOut = bAddr(jsonOutP)
	showVersion = bAddr(showVersionP)
	showHelp = bAddr(showHelpP)
	showTime = bAddr(showTimeP)
	showVerbose = bAddr(showVerboseP)
	showDebug = bAddr(showDebugP)
	showChecker = bAddr(showCheckerP)
	itemSpacing = float64(iAddr(itemSpacingP))
	scaleup = float64(fAddr(minMagP))

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
			log.Println("ERROR: Opening input file: "+inputFile+":", err)
			os.Exit(1)
		}
	}

	//Max input
	if len(input) > maxInput {
		log.Println("ERROR: Input data too large.")
		os.Exit(1)
	}

	if showVerbose {
		log.Println("Decoding...")
	}
	data, err := base64.StdEncoding.DecodeString(string(input[1:]))
	if err != nil {
		log.Println("ERROR: decoding input:", err)
		os.Exit(1)
	}

	if showVerbose {
		log.Println("Decompressing...")
	}

	r, err := zlib.NewReader(bytes.NewReader(data))
	if err != nil {
		log.Println("ERROR: decompress start failure:", err)
		os.Exit(1)
	}
	enflated, err := ioutil.ReadAll(r)
	if err != nil {
		log.Println("ERROR: Decompress read failure:", err)
		os.Exit(1)
	}

	//Max decompressed size
	if len(enflated) > maxJson {
		log.Println("ERROR: Input data too large.")
		os.Exit(1)
	}

	newbook := RootData{}

	if showVerbose {
		log.Println("Unmarshaling JSON...")
	}
	err = json.Unmarshal(enflated, &newbook)
	if err != nil {
		log.Println("ERROR: JSON Unmarshal failure:", err)
		os.Exit(1)
	}

	if showVerbose {
		log.Println("Drawing image...")
	}
	//Normalize cordinates
	maxx := 0.0
	maxy := 0.0

	var BPList []BluePrintData
	if len(newbook.Blueprintbook.Blueprints) == 0 {
		BPList = append(BPList, newbook.Blueprint)
	} else {
		for _, bp := range newbook.Blueprintbook.Blueprints {
			BPList = append(BPList, bp.Blueprint)
		}
	}

	for a, bp := range BPList {
		fmt.Println("Blueprint found: ", a)
		fmt.Println("Blueprint ent count: ", len(bp.Entities))

		for _, v := range bp.Entities {
			if v.Position.X > maxx {
				maxx = v.Position.X
			}
			if v.Position.Y > maxy {
				maxy = v.Position.Y
			}
		}
		minx := maxx
		miny := maxy
		for _, v := range bp.Entities {
			if v.Position.X < minx {
				minx = v.Position.X
			}
			if v.Position.Y < miny {
				miny = v.Position.Y
			}
		}

		//Size image
		xsize := maxx - minx
		ysize := maxy - miny

		if showVerbose {
			buf := fmt.Sprintf("BP size: %v x %v", int(xsize), int(ysize))
			log.Println(buf)
		}

		//Find largest side
		maxsize := 0.0
		if xsize > ysize {
			maxsize = xsize
		} else {
			maxsize = ysize
		}

		//Scale up to max magnifcation, or ideal size
		for a := scaleup; a <= maxmag; a = a + 1.0 {
			if maxsize*a < idealsize {
				scaleup = float64(a)
			}
		}

		//Offset x/y with margins and scale it up to the magnified size
		imx := (xsize + (tlSpace * 2)) * scaleup
		imy := (ysize + (tlSpace * 2)) * scaleup

		if imx > maxImage || imy > maxImage {
			log.Println("ERROR: Final image would be too large.")
			os.Exit(1)
		}

		if showVerbose {
			buf := fmt.Sprintf("Image size: %d x %d (%2.1fX mag)", int(imx), int(imy), scaleup)
			log.Println(buf)
		}

		//Allocate image and bg
		mapimage := image.NewRGBA(image.Rect(0, 0, int(imx), int(imy)))
		newimage := image.NewRGBA(image.Rect(0, 0, int(imx), int(imy)))

		//Draw map, scaled
		var objx, objy, x, y, xs, ys, xo, yo, ix, iy float64
		count := 0
		for _, v := range bp.Entities {

			//Offset position
			objx = v.Position.X - float64(minx)
			objy = v.Position.Y - float64(miny)
			//Get color for item
			item := findItem(v.Name)

			//Handle item rotation, crudely
			if v.Direction == EAST || v.Direction == WEST {
				ix = item.Y
				iy = item.X
			} else { //north/south, etc
				ix = item.X
				iy = item.Y
			}

			//If item side is larger than 1, deal with recentering it
			if v.Name == "offshore-pump" {
				//offshore pump is a special case apparently
				if v.Direction == EAST {
					objx = objx + 1.0
				} else if v.Direction == WEST {
					objx = objx - 1.0
				} else if v.Direction == NORTH {
					objy = objy - 1.0
				} else if v.Direction == SOUTH {
					objy = objy + 1.0
				}
			} else if v.Name == "straight-rail" {
				//also special case, because of 45 degree angles
				//TODO: handle 45s and curved rail better
				if v.Direction == NORTH_EAST {
					objy = objy - 0.5
					objx = objx + 0.5
				} else if v.Direction == SOUTH_EAST {
					objy = objy + 0.5
					objx = objx + 0.5
				} else if v.Direction == SOUTH_WEST {
					objy = objy + 0.5
					objx = objx - 0.5
				} else if v.Direction == NORTH_WEST {
					objy = objy - 0.5
					objx = objx - 0.5
				}
			}

			//Item x/y and margin offsets
			x = ((objx * scaleup) + (tlSpace * scaleup))
			y = ((objy * scaleup) + (tlSpace * scaleup))

			//Item size, scaled
			xs = ix * scaleup
			ys = iy * scaleup

			//Recenter items with sides larger than 1
			if xs > 1 {
				x = x - (xs / 2.0) + 0.5
			}
			if ys > 1 {
				y = y - (ys / 2.0) + 0.5
			}

			//Color vars for draw loop
			var r, g, b, a uint32

			//Dim image in debug mode, to highlight overdraw
			var newAlpha uint8
			if showDebug {
				newAlpha = 32
			} else {
				newAlpha = 255
			}

			//Draw item spacing if we are magnified, except special cases
			if scaleup > 1.0 && item.Name != "stone-wall" {
				if !item.KeepContinuous {
					xs = xs - itemSpacing
					ys = ys - itemSpacing
				} else {
					if v.Direction == EAST {
						ys = ys - itemSpacing
						y = y + 0.5
					} else if v.Direction == WEST {
						ys = ys - itemSpacing
					} else if v.Direction == NORTH {
						xs = xs - itemSpacing
					} else if v.Direction == SOUTH {
						xs = xs - itemSpacing
						y = y + 0.5
					}
				}
			}

			//Draw item, magnified
			for xo = 0.0; xo < xs; xo = xo + 1 {
				xxo := x + xo //x + offset
				for yo = 0.0; yo < ys; yo = yo + 1 {
					yyo := y + yo //y + offset

					//Get existing color at x/y if debugging
					if showDebug {
						r, g, b, a = (mapimage.At(int(xxo), int(yyo))).RGBA()
					}

					if !showDebug || r == 0 && g == 0 && b == 0 && a == 0 {
						//Get color from item
						r, g, b, _ := item.Color.RGBA()

						//Draw, bitshift to get to 8bit quickly
						mapimage.Set(int(xxo), int(yyo), color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), newAlpha})
					} else {
						//Highlight any overdraw in debug mode
						if showDebug {
							mapimage.Set(int(xxo), int(yyo), color.RGBA{255, 0, 0, 255})
							log.Println(fmt.Sprintf("Overdraw: Name:%v: xy:%v,%v, rgb:%v,%v,%v", v.Name, x+xo, y+yo, r>>8, g>>8, b>>8))
						}
					}
				}
			}

			//Count bp items
			count = count + 1
		}
		if showVerbose {
			log.Println(fmt.Sprintf("Item count: %v", count))
		}

		//Draw checkerboard background
		if showChecker {
			var c uint8
			//scale size to magnified size
			csize := int(checkersize * scaleup)

			//x/y, margin offset, scaled
			yoff := ((tlSpace) * scaleup)
			xoff := ((tlSpace + 1) * scaleup)

			for y := 0.0; y < imy; y++ {
				for x := 0.0; x < imx; x++ {

					//Only draw boarders
					if int(x-xoff)%csize != 0 && int(y-yoff)%csize != 0 {
						c = 32
					} else {
						c = 0
					}

					//If map has pixel here, draw it, otherwise draw BG
					if mapimage.At(int(x), int(y)) != (color.RGBA{0, 0, 0, 0}) {
						newimage.Set(int(x), int(y), mapimage.At(int(x), int(y)))
					} else {
						newimage.Set(int(x), int(y), color.RGBA{c, c, c, 255})
					}
				}
			}
		} else {
			newimage = mapimage
		}

		//If blueprint has a name, use it
		t := time.Now()
		bpname := ""
		if bp.Label != "" {
			bpname = bp.Label + "-"
		}

		//Timestamp option
		cTime := ""
		if showTime {
			cTime = fmt.Sprintf("%v", t.UnixNano())
		}

		//Write json file
		if jsonOut {
			fileName := fmt.Sprintf("%v%v%v.json", outputName, bpname, cTime)
			err = os.WriteFile(fileName, enflated, 0644)
			if err != nil {
				log.Println("ERROR: Failed to write json:", err)
				os.Exit(1)
			}
			log.Println("Wrote json to", fileName)
		}

		//Write the png file
		fileName := fmt.Sprintf("%v%v%v.png", outputName, bpname, cTime)
		output, err := os.Create(fileName)
		if png.Encode(output, newimage) != nil {
			log.Println("ERROR: Failed to write image:", err)
			os.Exit(1)
		}
		log.Println("Wrote image to", fileName)
		output.Close()
	}

}
