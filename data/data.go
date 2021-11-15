package data

import "image/color"

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

	Color          color.RGBA
	KeepContinuous bool
}

var ItemData = [...]Item{

	//Default
	{"default", 1, 1, ColorOrange, false},

	//Chests
	{"wooden-chest", 1, 1, ColorLightOrange, false},
	{"iron-chest", 1, 1, ColorLightOrange, false},
	{"steel-chest", 1, 1, ColorLightOrange, false},

	//Belts
	{"transport-belt", 1, 1, ColorLightGray, true},
	{"fast-transport-belt", 1, 1, ColorLightGray, true},
	{"express-transport-belt", 1, 1, ColorLightGray, true},

	//Unders
	{"underground-belt", 1, 1, ColorGray, false},
	{"fast-underground-belt", 1, 1, ColorGray, false},
	{"express-underground-belt", 1, 1, ColorGray, false},

	//Splitters
	{"splitter", 2, 1, ColorWhite, false},
	{"fast-splitter", 2, 1, ColorWhite, false},
	{"express-splitter", 2, 1, ColorWhite, false},

	//Inserters
	{"burner-inserter", 1, 1, ColorOrange, false},
	{"inserter", 1, 1, ColorOrange, false},
	{"long-handed-inserter", 1, 1, ColorOrange, false},
	{"fast-inserter", 1, 1, ColorOrange, false},
	{"filter-inserter", 1, 1, ColorOrange, false},
	{"stack-inserter", 1, 1, ColorOrange, false},
	{"stack-filter-inserter", 1, 1, ColorOrange, false},

	//Poles
	{"small-electric-pole", 1, 1, ColorRed, false},
	{"medium-electric-pole", 1, 1, ColorRed, false},
	{"big-electric-pole", 2, 2, ColorRed, false},
	{"substation", 2, 2, ColorRed, false},

	//Pipes
	{"pipe", 1, 1, ColorLightBlue, true},
	{"pipe-to-ground", 1, 1, ColorBlue, false},
	{"pump", 1, 2, ColorLightBlue, false},
	{"offshore-pump", 1, 2, ColorAqua, false},
	{"storage-tank", 2, 2, ColorLightAqua, false},

	//Rails
	{"straight-rail", 1, 2, ColorLightGreen, true},
	{"curved-rail", 2, 2, ColorLightGreen, true},
	{"train-stop", 2, 2, ColorGreen, false},
	{"rail-signal", 1, 1, ColorGreen, false},
	{"rail-chain-signal", 1, 1, ColorGreen, false},

	//Logistics
	{"logistic-chest-active-provider", 1, 1, ColorLightYellow, false},
	{"logistic-chest-passive-provider", 1, 1, ColorLightYellow, false},
	{"logistic-chest-storage", 1, 1, ColorLightYellow, false},
	{"logistic-chest-buffer", 1, 1, ColorLightYellow, false},
	{"logistic-chest-requester", 1, 1, ColorLightYellow, false},
	{"roboport", 4, 4, ColorYellow, false},

	//Lamp
	{"small-lamp", 1, 1, ColorWhite, false},

	//Combinators
	{"arithmetic-combinator", 1, 2, ColorLightTeal, false},
	{"decider-combinator", 1, 2, ColorLightTeal, false},
	{"constant-combinator", 1, 1, ColorLightTeal, false},
	{"power-switch", 2, 2, ColorTeal, false},
	{"programmable-speaker", 1, 1, ColorFuchsia, false},

	//Generators
	{"boiler", 3, 2, ColorLime, false},
	{"steam-engine", 3, 5, ColorLightLime, false},
	{"solar-panel", 3, 3, ColorLightLime, false},
	{"accumulator", 2, 2, ColorLime, false},
	{"nuclear-reactor", 5, 5, ColorAqua, false},
	{"heat-pipe", 1, 1, ColorPink, false},
	{"heat-exchanger", 3, 2, ColorLightPink, false},
	{"steam-turbine", 3, 5, ColorLightLime, false},

	//Miners
	{"burner-mining-drill", 2, 2, ColorSilver, false},
	{"electric-mining-drill", 3, 3, ColorSilver, false},
	{"pumpjack", 3, 3, ColorMaroon, false},

	//Furnaces
	{"stone-furnace", 2, 2, ColorRed, false},
	{"steel-furnace", 2, 2, ColorRed, false},
	{"electric-furnace", 3, 3, ColorRed, false},

	//Assemblers
	{"assembling-machine-1", 3, 3, ColorLightOrange, false},
	{"assembling-machine-2", 3, 3, ColorLightOrange, false},
	{"assembling-machine-3", 3, 3, ColorLightOrange, false},

	//Refineries
	{"oil-refinery", 5, 5, ColorPurple, false},
	{"chemical-plant", 3, 3, ColorPurple, false},
	{"centrifuge", 3, 3, ColorPurple, false},
	{"lab", 3, 3, ColorLightPurple, false},

	//Late-game
	{"beacon", 3, 3, ColorLightPurple, false},
	{"rocket-silo", 9, 9, ColorLightPurple, false},

	//Walls
	{"stone-wall", 1, 1, ColorGray, false},
	{"gate", 1, 1, ColorLightYellow, false},

	//Turrets
	{"gun-turret", 2, 2, ColorOrange, false},
	{"laser-turret", 2, 2, ColorOrange, false},
	{"flamethrower-turret", 2, 3, ColorOrange, false},
	{"artillery-turret", 3, 3, ColorOrange, false},

	//Radar
	{"radar", 3, 3, ColorLightOrange, false},
}
