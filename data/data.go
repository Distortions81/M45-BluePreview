package data

import "image/color"

const NORTH = 0
const NORTH_EAST = 1
const EAST = 2
const SOUTH_EAST = 3
const SOUTH = 4
const SOUTH_WEST = 5
const WEST = 6
const NORTH_WEST = 7

var ColorRed = color.RGBA{255, 0, 0, 255}
var ColorGreen = color.RGBA{0, 255, 0, 255}
var ColorBlue = color.RGBA{0, 0, 255, 255}
var ColorYellow = color.RGBA{255, 255, 0, 255}
var ColorBlack = color.RGBA{0, 0, 0, 255}
var ColorWhite = color.RGBA{255, 255, 255, 255}
var ColorGray = color.RGBA{128, 128, 128, 255}
var ColorOrange = color.RGBA{255, 165, 0, 255}
var ColorPink = color.RGBA{255, 192, 203, 255}
var ColorPurple = color.RGBA{128, 0, 128, 255}
var ColorSilver = color.RGBA{192, 192, 192, 255}
var ColorTeal = color.RGBA{0, 128, 128, 255}
var ColorMaroon = color.RGBA{128, 0, 0, 255}
var ColorNavy = color.RGBA{0, 0, 128, 255}
var ColorOlive = color.RGBA{128, 128, 0, 255}
var ColorLime = color.RGBA{0, 255, 0, 255}
var ColorFuchsia = color.RGBA{255, 0, 255, 255}
var ColorAqua = color.RGBA{0, 255, 255, 255}
var ColorTransparent = color.RGBA{0, 0, 0, 255}

var ColorLightRed = color.RGBA{255, 192, 192, 255}
var ColorLightGreen = color.RGBA{192, 255, 192, 255}
var ColorLightBlue = color.RGBA{192, 192, 255, 255}
var ColorLightYellow = color.RGBA{255, 255, 192, 255}
var ColorLightGray = color.RGBA{192, 192, 192, 255}
var ColorLightOrange = color.RGBA{255, 224, 192, 255}
var ColorLightPink = color.RGBA{255, 224, 224, 255}
var ColorLightPurple = color.RGBA{192, 192, 255, 255}
var ColorLightSilver = color.RGBA{224, 224, 224, 255}
var ColorLightTeal = color.RGBA{192, 224, 192, 255}
var ColorLightMaroon = color.RGBA{192, 192, 128, 255}
var ColorLightNavy = color.RGBA{192, 192, 128, 255}
var ColorLightOlive = color.RGBA{224, 192, 128, 255}
var ColorLightLime = color.RGBA{192, 255, 192, 255}
var ColorLightFuchsia = color.RGBA{255, 192, 255, 255}
var ColorLightAqua = color.RGBA{192, 255, 255, 255}

type Item struct {
	Name string
	X    float64
	Y    float64

	Color color.RGBA
}

var ItemData = [...]Item{

	//Default
	{"default", 1, 1, ColorOrange},

	//Chests
	{"wooden-chest", 1, 1, ColorLightOrange},
	{"iron-chest", 1, 1, ColorLightOrange},
	{"steel-chest", 1, 1, ColorLightOrange},

	//Belts
	{"transport-belt", 1, 1, ColorLightGray},
	{"fast-transport-belt", 1, 1, ColorLightGray},
	{"express-transport-belt", 1, 1, ColorLightGray},

	//Unders
	{"underground-belt", 1, 1, ColorGray},
	{"fast-underground-belt", 1, 1, ColorGray},
	{"express-underground-belt", 1, 1, ColorGray},

	//Splitters
	{"splitter", 2, 1, ColorWhite},
	{"fast-splitter", 2, 1, ColorWhite},
	{"express-splitter", 2, 1, ColorWhite},

	//Inserters
	{"burner-inserter", 1, 1, ColorOrange},
	{"inserter", 1, 1, ColorOrange},
	{"long-handed-inserter", 1, 1, ColorOrange},
	{"fast-inserter", 1, 1, ColorOrange},
	{"filter-inserter", 1, 1, ColorOrange},
	{"stack-inserter", 1, 1, ColorOrange},
	{"stack-filter-inserter", 1, 1, ColorOrange},

	//Poles
	{"small-electric-pole", 1, 1, ColorRed},
	{"medium-electric-pole", 1, 1, ColorRed},
	{"big-electric-pole", 2, 2, ColorRed},
	{"substation", 2, 2, ColorRed},

	//Pipes
	{"pipe", 1, 1, ColorLightBlue},
	{"pipe-to-ground", 1, 1, ColorBlue},
	{"pump", 1, 2, ColorLightBlue},
	{"storage-tank", 2, 2, ColorLightAqua},

	//Rails
	{"straight-rail", 2, 2, ColorLightGreen},
	{"curved-rail", 2, 2, ColorLightGreen},
	{"train-stop", 2, 2, ColorGreen},
	{"rail-signal", 1, 1, ColorGreen},
	{"rail-chain-signal", 1, 1, ColorGreen},

	//Logistics
	{"logistic-chest-active-provider", 1, 1, ColorLightYellow},
	{"logistic-chest-passive-provider", 1, 1, ColorLightYellow},
	{"logistic-chest-storage", 1, 1, ColorLightYellow},
	{"logistic-chest-buffer", 1, 1, ColorLightYellow},
	{"logistic-chest-requester", 1, 1, ColorLightYellow},
	{"roboport", 4, 4, ColorYellow},

	//Lamp
	{"small-lamp", 1, 1, ColorWhite},

	//Combinators
	{"arithmetic-combinator", 1, 2, ColorLightTeal},
	{"decider-combinator", 1, 2, ColorLightTeal},
	{"constant-combinator", 1, 1, ColorLightTeal},
	{"power-switch", 2, 2, ColorTeal},
	{"programmable-speaker", 1, 1, ColorFuchsia},

	//Generators
	{"boiler", 3, 2, ColorLime},
	{"steam-engine", 3, 5, ColorLightLime},
	{"solar-panel", 3, 3, ColorLightLime},
	{"accumulator", 2, 2, ColorLime},
	{"nuclear-reactor", 5, 5, ColorAqua},
	{"heat-pipe", 1, 1, ColorPink},
	{"heat-exchanger", 3, 2, ColorLightPink},
	{"steam-turbine", 3, 5, ColorLightLime},

	//Miners
	{"burner-mining-drill", 2, 2, ColorSilver},
	{"electric-mining-drill", 3, 3, ColorSilver},
	{"offshore-pump", 1, 2, ColorAqua},
	{"pumpjack", 3, 3, ColorMaroon},

	//Furnaces
	{"stone-furnace", 2, 2, ColorRed},
	{"steel-furnace", 2, 2, ColorRed},
	{"electric-furnace", 3, 3, ColorRed},

	//Assemblers
	{"assembling-machine-1", 3, 3, ColorLightOrange},
	{"assembling-machine-2", 3, 3, ColorLightOrange},
	{"assembling-machine-3", 3, 3, ColorLightOrange},

	//Refineries
	{"oil-refinery", 5, 5, ColorPurple},
	{"chemical-plant", 3, 3, ColorPurple},
	{"centrifuge", 3, 3, ColorPurple},
	{"lab", 3, 3, ColorLightPurple},

	//Late-game
	{"beacon", 3, 3, ColorLightPurple},
	{"rocket-silo", 9, 9, ColorLightPurple},

	//Walls
	{"stone-wall", 1, 1, ColorGray},
	{"gate", 1, 1, ColorLightYellow},

	//Turrets
	{"gun-turret", 2, 2, ColorOrange},
	{"laser-turret", 2, 2, ColorOrange},
	{"flamethrower-turret", 2, 3, ColorOrange},
	{"artillery-turret", 3, 3, ColorOrange},

	//Radar
	{"radar", 3, 3, ColorLightOrange},
}
