package climate

import "math/rand"

// Tree is a tree
type Tree struct {
	Name           string
	PluralName     string
	GivesBark      bool
	GivesFruit     bool
	GivesNuts      bool
	IsDeciduous    bool
	IsConiferous   bool
	IsHard         bool
	IsSoft         bool
	IsMedicine     bool
	IsSpice        bool
	MinHumidity    int
	MaxHumidity    int
	MinTemperature int
	MaxTemperature int
}

func (climate Climate) getFilteredTrees() []Tree {
	trees := getAllTrees()
	trees = filterTreesForHumidity(climate.Humidity, trees)
	trees = filterTreesForTemperature(climate.Temperature, trees)
	trees = filterTreesForType(climate, trees)

	return trees
}

func getRandomTrees(amount int, from []Tree) []Tree {
	var plant Tree

	trees := []Tree{}

	if amount > len(from) {
		amount = len(from)
	}

	if len(from) == 0 {
		return trees
	}

	for i := 0; i < amount; i++ {
		plant = from[rand.Intn(len(from)-1)]
		if !isTreeInSlice(plant, trees) {
			trees = append(trees, plant)
		}
	}

	return trees
}

func isTreeInSlice(plant Tree, trees []Tree) bool {
	isIt := false
	for _, a := range trees {
		if a.Name == plant.Name {
			isIt = true
		}
	}

	return isIt
}

func filterTreesForType(climate Climate, trees []Tree) []Tree {
	var filteredTrees []Tree

	for _, a := range trees {
		if (climate.HasDeciduousTrees && a.IsDeciduous) || (climate.HasConiferousTrees && a.IsConiferous) {
			filteredTrees = append(filteredTrees, a)
		}
	}

	return filteredTrees
}

func filterTreesForHumidity(humidity int, trees []Tree) []Tree {
	var filteredTrees []Tree

	for _, a := range trees {
		if a.MinHumidity <= humidity && a.MaxHumidity >= humidity {
			filteredTrees = append(filteredTrees, a)
		}
	}

	return filteredTrees
}

func filterTreesForTemperature(temperature int, trees []Tree) []Tree {
	var filteredTrees []Tree

	for _, a := range trees {
		if a.MinTemperature <= temperature && a.MaxTemperature >= temperature {
			filteredTrees = append(filteredTrees, a)
		}
	}

	return filteredTrees
}

func getAllTrees() []Tree {
	trees := []Tree{
		Tree{
			Name:           "acacia",
			PluralName:     "acacias",
			GivesBark:      true,
			GivesFruit:     false,
			GivesNuts:      false,
			IsDeciduous:    true,
			IsConiferous:   false,
			IsHard:         true,
			IsSoft:         false,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
		},
		Tree{
			Name:           "alder",
			PluralName:     "alders",
			GivesBark:      true,
			GivesFruit:     false,
			GivesNuts:      false,
			IsDeciduous:    true,
			IsConiferous:   false,
			IsHard:         true,
			IsSoft:         false,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
		},
		Tree{
			Name:           "apple",
			PluralName:     "apples",
			GivesBark:      true,
			GivesFruit:     true,
			GivesNuts:      false,
			IsDeciduous:    true,
			IsConiferous:   false,
			IsHard:         true,
			IsSoft:         false,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
		},
		Tree{
			Name:           "ash",
			PluralName:     "ashes",
			GivesBark:      true,
			GivesFruit:     false,
			GivesNuts:      false,
			IsDeciduous:    true,
			IsConiferous:   false,
			IsHard:         true,
			IsSoft:         false,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
		},
		Tree{
			Name:           "aspen",
			PluralName:     "aspens",
			GivesBark:      true,
			GivesFruit:     false,
			GivesNuts:      false,
			IsDeciduous:    true,
			IsConiferous:   false,
			IsHard:         true,
			IsSoft:         false,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
		},
		Tree{
			Name:           "balsa",
			PluralName:     "balsas",
			GivesBark:      true,
			GivesFruit:     false,
			GivesNuts:      false,
			IsDeciduous:    true,
			IsConiferous:   false,
			IsHard:         true,
			IsSoft:         false,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
		},
		Tree{
			Name:           "banana",
			PluralName:     "bananas",
			GivesBark:      true,
			GivesFruit:     true,
			GivesNuts:      false,
			IsDeciduous:    true,
			IsConiferous:   false,
			IsHard:         true,
			IsSoft:         false,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    4,
			MaxHumidity:    10,
			MinTemperature: 7,
			MaxTemperature: 10,
		},
		Tree{
			Name:           "birch",
			PluralName:     "birches",
			GivesBark:      true,
			GivesFruit:     false,
			GivesNuts:      false,
			IsDeciduous:    true,
			IsConiferous:   false,
			IsHard:         true,
			IsSoft:         false,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
		},
		Tree{
			Name:           "black pine",
			PluralName:     "black pines",
			GivesBark:      true,
			GivesFruit:     false,
			GivesNuts:      false,
			IsDeciduous:    false,
			IsConiferous:   true,
			IsHard:         false,
			IsSoft:         true,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
		},
		Tree{
			Name:           "blackwood",
			PluralName:     "blackwoods",
			GivesBark:      true,
			GivesFruit:     false,
			GivesNuts:      false,
			IsDeciduous:    true,
			IsConiferous:   false,
			IsHard:         true,
			IsSoft:         false,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
		},
		Tree{
			Name:           "boxwood",
			PluralName:     "boxwoods",
			GivesBark:      true,
			GivesFruit:     false,
			GivesNuts:      false,
			IsDeciduous:    true,
			IsConiferous:   false,
			IsHard:         true,
			IsSoft:         false,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
		},
		Tree{
			Name:           "cedar",
			PluralName:     "cedars",
			GivesBark:      true,
			GivesFruit:     false,
			GivesNuts:      false,
			IsDeciduous:    false,
			IsConiferous:   true,
			IsHard:         false,
			IsSoft:         true,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
		},
		Tree{
			Name:           "cherry",
			PluralName:     "cherries",
			GivesBark:      true,
			GivesFruit:     true,
			GivesNuts:      false,
			IsDeciduous:    true,
			IsConiferous:   false,
			IsHard:         true,
			IsSoft:         false,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
		},
		Tree{
			Name:           "cinnamon",
			PluralName:     "cinnamons",
			GivesBark:      true,
			GivesFruit:     false,
			GivesNuts:      false,
			IsDeciduous:    true,
			IsConiferous:   false,
			IsHard:         true,
			IsSoft:         false,
			IsMedicine:     false,
			IsSpice:        true,
			MinHumidity:    2,
			MaxHumidity:    8,
			MinTemperature: 0,
			MaxTemperature: 8,
		},
		Tree{
			Name:           "coachwood",
			PluralName:     "coachwoods",
			GivesBark:      true,
			GivesFruit:     false,
			GivesNuts:      false,
			IsDeciduous:    true,
			IsConiferous:   false,
			IsHard:         true,
			IsSoft:         false,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
		},
		Tree{
			Name:           "corkwood",
			PluralName:     "corkwoods",
			GivesBark:      true,
			GivesFruit:     false,
			GivesNuts:      false,
			IsDeciduous:    true,
			IsConiferous:   false,
			IsHard:         true,
			IsSoft:         false,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
		},
		Tree{
			Name:           "cottonwood",
			PluralName:     "cottonwoods",
			GivesBark:      true,
			GivesFruit:     false,
			GivesNuts:      false,
			IsDeciduous:    true,
			IsConiferous:   false,
			IsHard:         true,
			IsSoft:         false,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
		},
		Tree{
			Name:           "crabapple",
			PluralName:     "crabapples",
			GivesBark:      true,
			GivesFruit:     true,
			GivesNuts:      false,
			IsDeciduous:    true,
			IsConiferous:   false,
			IsHard:         true,
			IsSoft:         false,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
		},
		Tree{
			Name:           "cypress",
			PluralName:     "cypress",
			GivesBark:      true,
			GivesFruit:     false,
			GivesNuts:      false,
			IsDeciduous:    false,
			IsConiferous:   true,
			IsHard:         false,
			IsSoft:         true,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
		},
		Tree{
			Name:           "dogwood",
			PluralName:     "dogwoods",
			GivesBark:      true,
			GivesFruit:     false,
			GivesNuts:      false,
			IsDeciduous:    true,
			IsConiferous:   false,
			IsHard:         true,
			IsSoft:         false,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
		},
		Tree{
			Name:           "elm",
			PluralName:     "elms",
			GivesBark:      true,
			GivesFruit:     false,
			GivesNuts:      false,
			IsDeciduous:    true,
			IsConiferous:   false,
			IsHard:         true,
			IsSoft:         false,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 5,
			MaxTemperature: 7,
		},
		Tree{
			Name:           "eucalyptus",
			PluralName:     "eucalyptuses",
			GivesBark:      true,
			GivesFruit:     false,
			GivesNuts:      false,
			IsDeciduous:    true,
			IsConiferous:   false,
			IsHard:         true,
			IsSoft:         false,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
		},
		Tree{
			Name:           "fir",
			PluralName:     "firs",
			GivesBark:      true,
			GivesFruit:     false,
			GivesNuts:      false,
			IsDeciduous:    false,
			IsConiferous:   true,
			IsHard:         false,
			IsSoft:         true,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
		},
		Tree{
			Name:           "hemlock",
			PluralName:     "hemlocks",
			GivesBark:      true,
			GivesFruit:     false,
			GivesNuts:      false,
			IsDeciduous:    false,
			IsConiferous:   true,
			IsHard:         false,
			IsSoft:         true,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
		},
		Tree{
			Name:           "ironwood",
			PluralName:     "ironwoods",
			GivesBark:      true,
			GivesFruit:     false,
			GivesNuts:      false,
			IsDeciduous:    true,
			IsConiferous:   false,
			IsHard:         true,
			IsSoft:         false,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
		},
		Tree{
			Name:           "juniper",
			PluralName:     "junipers",
			GivesBark:      true,
			GivesFruit:     true,
			GivesNuts:      false,
			IsDeciduous:    false,
			IsConiferous:   true,
			IsHard:         false,
			IsSoft:         true,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
		},
		Tree{
			Name:           "kingwood",
			PluralName:     "kingwoods",
			GivesBark:      true,
			GivesFruit:     false,
			GivesNuts:      false,
			IsDeciduous:    true,
			IsConiferous:   false,
			IsHard:         true,
			IsSoft:         false,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
		},
		Tree{
			Name:           "lacewood",
			PluralName:     "lacewoods",
			GivesBark:      true,
			GivesFruit:     false,
			GivesNuts:      false,
			IsDeciduous:    true,
			IsConiferous:   false,
			IsHard:         true,
			IsSoft:         false,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
		},
		Tree{
			Name:           "larch",
			PluralName:     "larches",
			GivesBark:      true,
			GivesFruit:     false,
			GivesNuts:      false,
			IsDeciduous:    false,
			IsConiferous:   true,
			IsHard:         false,
			IsSoft:         true,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
		},
		Tree{
			Name:           "lemon",
			PluralName:     "lemons",
			GivesBark:      true,
			GivesFruit:     true,
			GivesNuts:      false,
			IsDeciduous:    true,
			IsConiferous:   false,
			IsHard:         true,
			IsSoft:         false,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    2,
			MaxHumidity:    9,
			MinTemperature: 0,
			MaxTemperature: 9,
		},
		Tree{
			Name:           "lime",
			PluralName:     "limes",
			GivesBark:      true,
			GivesFruit:     true,
			GivesNuts:      false,
			IsDeciduous:    true,
			IsConiferous:   false,
			IsHard:         true,
			IsSoft:         false,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    2,
			MaxHumidity:    9,
			MinTemperature: 0,
			MaxTemperature: 9,
		},
		Tree{
			Name:           "mahogany",
			PluralName:     "mahoganies",
			GivesBark:      true,
			GivesFruit:     false,
			GivesNuts:      false,
			IsDeciduous:    true,
			IsConiferous:   false,
			IsHard:         true,
			IsSoft:         false,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
		},
		Tree{
			Name:           "mango",
			PluralName:     "mangoes",
			GivesBark:      true,
			GivesFruit:     true,
			GivesNuts:      false,
			IsDeciduous:    true,
			IsConiferous:   false,
			IsHard:         true,
			IsSoft:         false,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    4,
			MaxHumidity:    10,
			MinTemperature: 6,
			MaxTemperature: 10,
		},
		Tree{
			Name:           "maple",
			PluralName:     "maples",
			GivesBark:      true,
			GivesFruit:     false,
			GivesNuts:      false,
			IsDeciduous:    true,
			IsConiferous:   false,
			IsHard:         true,
			IsSoft:         false,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 5,
			MaxTemperature: 7,
		},
		Tree{
			Name:           "oak",
			PluralName:     "oaks",
			GivesBark:      true,
			GivesFruit:     false,
			GivesNuts:      false,
			IsDeciduous:    true,
			IsConiferous:   false,
			IsHard:         true,
			IsSoft:         false,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 5,
			MaxTemperature: 7,
		},
		Tree{
			Name:           "olive",
			PluralName:     "olives",
			GivesBark:      true,
			GivesFruit:     true,
			GivesNuts:      false,
			IsDeciduous:    true,
			IsConiferous:   false,
			IsHard:         true,
			IsSoft:         false,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
		},
		Tree{
			Name:           "papaya",
			PluralName:     "papayas",
			GivesBark:      true,
			GivesFruit:     true,
			GivesNuts:      false,
			IsDeciduous:    true,
			IsConiferous:   false,
			IsHard:         true,
			IsSoft:         false,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    4,
			MaxHumidity:    10,
			MinTemperature: 6,
			MaxTemperature: 10,
		},
		Tree{
			Name:           "pine",
			PluralName:     "pine",
			GivesBark:      true,
			GivesFruit:     false,
			GivesNuts:      true,
			IsDeciduous:    false,
			IsConiferous:   true,
			IsHard:         false,
			IsSoft:         true,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
		},
		Tree{
			Name:           "pineapple",
			PluralName:     "pineapples",
			GivesBark:      true,
			GivesFruit:     true,
			GivesNuts:      false,
			IsDeciduous:    true,
			IsConiferous:   false,
			IsHard:         true,
			IsSoft:         false,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    2,
			MaxHumidity:    9,
			MinTemperature: 6,
			MaxTemperature: 9,
		},
		Tree{
			Name:           "palm",
			PluralName:     "palms",
			GivesBark:      true,
			GivesFruit:     false,
			GivesNuts:      false,
			IsDeciduous:    true,
			IsConiferous:   false,
			IsHard:         true,
			IsSoft:         false,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    2,
			MaxHumidity:    10,
			MinTemperature: 6,
			MaxTemperature: 10,
		},
		Tree{
			Name:           "pear",
			PluralName:     "pears",
			GivesBark:      true,
			GivesFruit:     true,
			GivesNuts:      false,
			IsDeciduous:    true,
			IsConiferous:   false,
			IsHard:         true,
			IsSoft:         false,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
		},
		Tree{
			Name:           "poplar",
			PluralName:     "poplars",
			GivesBark:      true,
			GivesFruit:     false,
			GivesNuts:      false,
			IsDeciduous:    true,
			IsConiferous:   false,
			IsHard:         true,
			IsSoft:         false,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
		},
		Tree{
			Name:           "red oak",
			PluralName:     "red oaks",
			GivesBark:      true,
			GivesFruit:     false,
			GivesNuts:      false,
			IsDeciduous:    true,
			IsConiferous:   false,
			IsHard:         true,
			IsSoft:         false,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
		},
		Tree{
			Name:           "red pine",
			PluralName:     "red pines",
			GivesBark:      true,
			GivesFruit:     false,
			GivesNuts:      false,
			IsDeciduous:    false,
			IsConiferous:   true,
			IsHard:         false,
			IsSoft:         true,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
		},
		Tree{
			Name:           "rosewood",
			PluralName:     "rosewoods",
			GivesBark:      true,
			GivesFruit:     false,
			GivesNuts:      false,
			IsDeciduous:    true,
			IsConiferous:   false,
			IsHard:         true,
			IsSoft:         false,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
		},
		Tree{
			Name:           "sandalwood",
			PluralName:     "sandalwoods",
			GivesBark:      true,
			GivesFruit:     false,
			GivesNuts:      false,
			IsDeciduous:    true,
			IsConiferous:   false,
			IsHard:         true,
			IsSoft:         false,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
		},
		Tree{
			Name:           "spruce",
			PluralName:     "spruces",
			GivesBark:      true,
			GivesFruit:     false,
			GivesNuts:      false,
			IsDeciduous:    false,
			IsConiferous:   true,
			IsHard:         false,
			IsSoft:         true,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
		},
		Tree{
			Name:           "teak",
			PluralName:     "teaks",
			GivesBark:      true,
			GivesFruit:     false,
			GivesNuts:      false,
			IsDeciduous:    true,
			IsConiferous:   false,
			IsHard:         true,
			IsSoft:         false,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
		},
		Tree{
			Name:           "walnut",
			PluralName:     "walnuts",
			GivesBark:      true,
			GivesFruit:     false,
			GivesNuts:      true,
			IsDeciduous:    true,
			IsConiferous:   false,
			IsHard:         true,
			IsSoft:         false,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
		},
		Tree{
			Name:           "white oak",
			PluralName:     "white oaks",
			GivesBark:      true,
			GivesFruit:     false,
			GivesNuts:      false,
			IsDeciduous:    true,
			IsConiferous:   false,
			IsHard:         true,
			IsSoft:         false,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
		},
		Tree{
			Name:           "white pine",
			PluralName:     "white pines",
			GivesBark:      true,
			GivesFruit:     false,
			GivesNuts:      true,
			IsDeciduous:    false,
			IsConiferous:   true,
			IsHard:         false,
			IsSoft:         true,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
		},
		Tree{
			Name:           "willow",
			PluralName:     "willows",
			GivesBark:      true,
			GivesFruit:     false,
			GivesNuts:      false,
			IsDeciduous:    true,
			IsConiferous:   false,
			IsHard:         true,
			IsSoft:         false,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
		},
		Tree{
			Name:           "yew",
			PluralName:     "yews",
			GivesBark:      true,
			GivesFruit:     false,
			GivesNuts:      false,
			IsDeciduous:    false,
			IsConiferous:   true,
			IsHard:         false,
			IsSoft:         true,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
		},
		Tree{
			Name:           "yellow pine",
			PluralName:     "yellow pines",
			GivesBark:      true,
			GivesFruit:     false,
			GivesNuts:      true,
			IsDeciduous:    false,
			IsConiferous:   true,
			IsHard:         false,
			IsSoft:         true,
			IsMedicine:     false,
			IsSpice:        false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
		},
	}

	return trees
}
