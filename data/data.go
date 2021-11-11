package data

type Item struct {
	name string
	x    int
	y    int
}

var ItemData = [...]Item{

	//Chests
	{"wooden-chest", 1, 1},
	{"iron-chest", 1, 1},
	{"steel-chest", 1, 1},

	//Belts
	{"transport-belt", 1, 1},
	{"fast-transport-belt", 1, 1},
	{"express-transport-belt", 1, 1},

	//Unders
	{"underground-belt", 1, 1},
	{"fast-underground-belt", 1, 1},
	{"express-underground-belt", 1, 1},

	//Splitters
	{"splitter", 2, 1},
	{"fast-splitter", 2, 1},
	{"express-splitter", 2, 1},

	//Inserters
	{"burner-inserter", 1, 1},
	{"inserter", 1, 1},
	{"long-handed-inserter", 1, 1},
	{"fast-inserter", 1, 1},
	{"filter-inserter", 1, 1},
	{"stack-inserter", 1, 1},
	{"stack-filter-inserter", 1, 1},

	//Poles
	{"small-electric-pole", 1, 1},
	{"medium-electric-pole", 1, 1},
	{"big-electric-pole", 1, 1},
	{"substation", 1, 1},

	//Pipes
	{"pipe", 1, 1},
	{"pipe-to-ground", 1, 1},
	{"pump", 1, 2},

	{"", 1, 1},
	{"", 1, 1},
	{"", 1, 1},
	{"", 1, 1},
	{"", 1, 1},
	{"", 1, 1},
	{"", 1, 1},
	{"", 1, 1},
	{"", 1, 1},
	{"", 1, 1},
	{"", 1, 1},
	{"", 1, 1},
	{"", 1, 1},
	{"", 1, 1},
	{"", 1, 1},
	{"", 1, 1}
}
