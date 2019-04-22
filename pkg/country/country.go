package country

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/culture"
	"github.com/ironarachne/world/pkg/heraldry"
	"github.com/ironarachne/world/pkg/region"
)

// Country is a geographic and political area
type Country struct {
	Name            string
	DominantCulture culture.Culture
	Government      Government
	Regions         []region.Region
	Capital         string
	Heraldry        heraldry.Heraldry
}

// Generate procedurally generates a country
func Generate() Country {
	regions := []region.Region{}
	country := Country{}

	country.DominantCulture = culture.Generate()
	country.Government = country.getNewMonarchy()
	country.Heraldry = heraldry.GenerateHeraldry()

	size := rand.Intn(10) + 4

	r := region.Region{}

	for i := 0; i < size; i++ {
		r = region.Generate(country.DominantCulture.HomeClimate.Name)
		regions = append(regions, r)
	}

	highestPopulation := 0
	capital := ""

	for _, r := range regions {
		for _, t := range r.Towns {
			if t.Population > highestPopulation {
				highestPopulation = t.Population
				capital = t.Name
			}
		}
	}

	country.Capital = capital

	country.Regions = regions

	return country
}