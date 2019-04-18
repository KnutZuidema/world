package town

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/character"
	"github.com/ironarachne/world/pkg/climate"
	"github.com/ironarachne/world/pkg/culture"
	"github.com/ironarachne/world/pkg/goods"
)

// Town is a town
type Town struct {
	Name             string
	Population       int
	Category         Category
	Climate          climate.Climate
	Culture          culture.Culture
	Mayor            character.Character
	NotableProducers []goods.Producer
	Exports          []goods.TradeGood
	Imports          []goods.TradeGood
}

func (town Town) generateMayor() character.Character {
	mayor := character.GenerateCharacterOfCulture(town.Culture)
	mayor = mayor.ChangeAge(rand.Intn(30) + 30)

	return mayor
}

func (town Town) generateRandomExports() []goods.TradeGood {
	var exports []goods.TradeGood

	exports = goods.GenerateExportTradeGoods(town.Category.MinExports, town.Category.MaxExports, town.NotableProducers, town.Climate.Resources)

	return exports
}

func (town Town) generateRandomImports() []goods.TradeGood {
	var imports []goods.TradeGood

	foreignClimate := climate.GetForeignClimate(town.Climate)

	imports = goods.GenerateImportTradeGoods(town.Category.MinImports, town.Category.MaxImports, foreignClimate.Resources)

	return imports
}

func generateRandomPopulation(category Category) int {
	sizeIncrement := category.MaxSize - category.MinSize

	return rand.Intn(sizeIncrement) + category.MinSize
}

func (town Town) generateTownName() string {
	name := town.Culture.Language.RandomName()
	return name
}

// Generate generates a random town
func Generate(category string, biome string) Town {
	town := Town{}

	if category == "random" {
		town.Category = getRandomWeightedCategory()
	} else {
		town.Category = getCategoryByName(category)
	}
	if biome == "random" {
		town.Climate = climate.Generate()
	} else {
		town.Climate = climate.GetClimate(biome)
	}

	culture := culture.Generate()
	culture = culture.SetClimate(town.Climate.Name)
	town = SetCulture(culture, town)

	town.Mayor = town.generateMayor()
	town.Name = town.generateTownName()

	possibleProducers := goods.GetPossibleProducers(town.Climate.Resources)
	town.NotableProducers = goods.GetRandomProducers(3, possibleProducers)

	town.Exports = town.generateRandomExports()
	town.Imports = town.generateRandomImports()
	town.Population = generateRandomPopulation(town.Category)

	return town
}

// SetCulture sets the culture of a town and recalculates some things
func SetCulture(culture culture.Culture, town Town) Town {
	newTown := town

	newTown.Culture = culture
	newTown.Name = newTown.Culture.Language.RandomName()
	newTown.Mayor.FirstName = newTown.Culture.Language.RandomGenderedName(town.Mayor.Gender)
	newTown.Mayor.LastName = newTown.Culture.Language.RandomName()

	return newTown
}
