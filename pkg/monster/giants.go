package monster

import (
	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/size"
)

func getAllGiants() []Monster {
	return []Monster{
		{
			Name:                 "cyclops",
			PluralName:           "cyclopses",
			SizeCategory:         size.GetCategoryByName("huge"),
			Type:                 "giant",
			IdealTemperature:     5,
			TemperatureTolerance: 5,
			NumAppearing: func() int {
				return 1
			},
		},
		{
			Name:                 "ettin",
			PluralName:           "ettins",
			SizeCategory:         size.GetCategoryByName("large"),
			Type:                 "giant",
			IdealTemperature:     5,
			TemperatureTolerance: 5,
			NumAppearing: func() int {
				num := random.DiceTotal(1, 2)
				return num
			},
		},
		{
			Name:                 "giant",
			PluralName:           "giants",
			SizeCategory:         size.GetCategoryByName("huge"),
			Type:                 "giant",
			IdealTemperature:     5,
			TemperatureTolerance: 5,
			NumAppearing: func() int {
				return 1
			},
		},
		{
			Name:                 "ogre",
			PluralName:           "ogres",
			SizeCategory:         size.GetCategoryByName("large"),
			Type:                 "giant",
			IdealTemperature:     5,
			TemperatureTolerance: 5,
			NumAppearing: func() int {
				return 1
			},
		},
	}
}