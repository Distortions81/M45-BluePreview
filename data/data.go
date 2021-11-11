package data

type ColorRGB8 struct {
	R uint8
	G uint8
	B uint8
}

var RED = ColorRGB8{255, 0, 0}

var GREEN = ColorRGB8{0, 255, 0}
var BLUE = ColorRGB8{0, 0, 255}
var WHITE = ColorRGB8{255, 255, 255}
var BLACK = ColorRGB8{0, 0, 0}
var YELLOW = ColorRGB8{255, 255, 0}
var CYAN = ColorRGB8{0, 255, 255}
var MAGENTA = ColorRGB8{255, 0, 255}
var ORANGE = ColorRGB8{255, 128, 0}
var PURPLE = ColorRGB8{128, 0, 128}
var BROWN = ColorRGB8{128, 64, 0}
var LIGHT_GRAY = ColorRGB8{192, 192, 192}
var DARK_GRAY = ColorRGB8{64, 64, 64}
var LIGHT_RED = ColorRGB8{255, 128, 128}
var LIGHT_GREEN = ColorRGB8{128, 255, 128}
var LIGHT_BLUE = ColorRGB8{128, 128, 255}
var LIGHT_YELLOW = ColorRGB8{55, 255, 128}
var LIGHT_CYAN = ColorRGB8{128, 255, 255}
var LIGHT_MAGENTA = ColorRGB8{128, 128, 255}
var LIGHT_ORANGE = ColorRGB8{255, 192, 128}
var LIGHT_PURPLE = ColorRGB8{192, 128, 255}
var LIGHT_BROWN = ColorRGB8{192, 96, 128}

var DARK_RED = ColorRGB8{128, 0, 0}
var DARK_GREEN = ColorRGB8{0, 128, 0}
var DARK_BLUE = ColorRGB8{0, 0, 128}
var DARK_YELLOW = ColorRGB8{128, 128, 0}
var DARK_CYAN = ColorRGB8{0, 128, 128}
var DARK_MAGENTA = ColorRGB8{128, 0, 128}
var DARK_ORANGE = ColorRGB8{128, 64, 0}
var DARK_PURPLE = ColorRGB8{64, 0, 64}
var DARK_BROWN = ColorRGB8{64, 32, 0}

type Item struct {
	name string
	x    int
	y    int

	color ColorRGB8
}

var ItemData = [...]Item{

	//Default
	{"default", 1, 1, LIGHT_ORANGE},

	//Chests
	{"wooden-chest", 1, 1, LIGHT_ORANGE},
	{"iron-chest", 1, 1, LIGHT_ORANGE},
	{"steel-chest", 1, 1, LIGHT_ORANGE},

	//Belts
	{"transport-belt", 1, 1, LIGHT_GRAY},
	{"fast-transport-belt", 1, 1, LIGHT_GRAY},
	{"express-transport-belt", 1, 1, LIGHT_GRAY},

	//Unders
	{"underground-belt", 1, 1, DARK_GRAY},
	{"fast-underground-belt", 1, 1, DARK_GRAY},
	{"express-underground-belt", 1, 1, DARK_GRAY},

	//Splitters
	{"splitter", 2, 1, DARK_GRAY},
	{"fast-splitter", 2, 1, DARK_GRAY},
	{"express-splitter", 2, 1, DARK_GRAY},

	//Inserters
	{"burner-inserter", 1, 1, WHITE},
	{"inserter", 1, 1, WHITE},
	{"long-handed-inserter", 1, 1, WHITE},
	{"fast-inserter", 1, 1, WHITE},
	{"filter-inserter", 1, 1, WHITE},
	{"stack-inserter", 1, 1, WHITE},
	{"stack-filter-inserter", 1, 1, WHITE},

	//Poles
	{"small-electric-pole", 1, 1, LIGHT_YELLOW},
	{"medium-electric-pole", 1, 1, LIGHT_YELLOW},
	{"big-electric-pole", 2, 2, LIGHT_YELLOW},
	{"substation", 2, 2, LIGHT_YELLOW},

	//Pipes
	{"pipe", 1, 1, DARK_BLUE},
	{"pipe-to-ground", 1, 1, DARK_BLUE},
	{"pump", 1, 2, BLUE},

	//Rails
	{"straight-rail", 2, 2, LIGHT_GRAY},
	{"curved-rail", 2, 2, LIGHT_GRAY},
	{"train-stop", 2, 2, LIGHT_GRAY},
	{"rail-signal", 1, 1, LIGHT_GRAY},
	{"rail-chain-signal", 1, 1, LIGHT_GRAY},

	//Logistics
	{"logistic-chest-active-provider", 1, 1, LIGHT_GRAY},
	{"logistic-chest-passive-provider", 1, 1, LIGHT_GRAY},
	{"logistic-chest-storage", 1, 1, LIGHT_GRAY},
	{"logistic-chest-buffer", 1, 1, LIGHT_GRAY},
	{"logistic-chest-requester", 1, 1, LIGHT_GRAY},
	{"roboport", 4, 4, LIGHT_GRAY},

	//Lamp
	{"small-lamp", 1, 1, LIGHT_YELLOW},

	//Combinators
	{"arithmetic-combinator", 1, 1, LIGHT_GRAY},
	{"decider-combinator", 1, 1, LIGHT_GRAY},
	{"constant-combinator", 1, 1, LIGHT_GRAY},
	{"power-switch", 1, 1, LIGHT_GRAY},
	{"programmable-speaker", 1, 1, LIGHT_GRAY},

	//Generators
	{"boiler", 1, 1, LIGHT_GRAY},
	{"steam-engine", 1, 1, LIGHT_GRAY},
	{"solar-panel", 1, 1, LIGHT_GRAY},
	{"accumulator", 1, 1, LIGHT_GRAY},
	{"nuclear-reactor", 1, 1, LIGHT_GRAY},
	{"heat-pipe", 1, 1, LIGHT_GRAY},
	{"heat-exchanger", 1, 1, LIGHT_GRAY},
	{"steam-turbine", 1, 1, LIGHT_GRAY},

	//Miners
	{"burner-mining-drill", 1, 1, LIGHT_GRAY},
	{"electric-mining-drill", 1, 1, LIGHT_GRAY},
	{"offshore-pump", 1, 1, LIGHT_GRAY},
	{"pumpjack", 1, 1, LIGHT_GRAY},

	//Furnaces
	{"stone-furnace", 1, 1, LIGHT_GRAY},
	{"steel-furnace", 1, 1, LIGHT_GRAY},
	{"electric-furnace", 1, 1, LIGHT_GRAY},

	//Assemblers
	{"assembling-machine-1", 1, 1, LIGHT_GRAY},
	{"assembling-machine-2", 1, 1, LIGHT_GRAY},
	{"assembling-machine-3", 1, 1, LIGHT_GRAY},

	//Refineries
	{"oil-refinery", 1, 1, LIGHT_GRAY},
	{"chemical-plant", 1, 1, LIGHT_GRAY},
	{"centrifuge", 1, 1, LIGHT_GRAY},
	{"lab", 1, 1, LIGHT_GRAY},

	//Late-game
	{"beacon", 1, 1, LIGHT_GRAY},
	{"rocket-silo", 1, 1, LIGHT_GRAY},

	//Walls
	{"stone-wall", 1, 1, LIGHT_GRAY},
	{"gate", 1, 1, LIGHT_GRAY},

	//Turrets
	{"gun-turret", 1, 1, LIGHT_GRAY},
	{"laser-turret", 1, 1, LIGHT_GRAY},
	{"flamethrower-turret", 1, 1, LIGHT_GRAY},
	{"artillery-turret", 1, 1, LIGHT_GRAY},

	//Radar
	{"radar", 1, 1, LIGHT_GRAY},
}
