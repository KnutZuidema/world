/*
Package merchant allows for creation of travelling merchants and their wares
*/
package merchant

import (
	"fmt"
	"github.com/ironarachne/world/pkg/geography"

	"github.com/ironarachne/world/pkg/character"
	"github.com/ironarachne/world/pkg/culture"
	"github.com/ironarachne/world/pkg/goods"
	"github.com/ironarachne/world/pkg/profession"
	"github.com/ironarachne/world/pkg/town"
)

// Merchant is a merchant
type Merchant struct {
	Character character.Character
	Goods     []goods.TradeGood
}

const merchantRandomGenerationError = "failed to generate random merchant: %w"

// Generate returns a random merchant
func Generate(originTown town.Town) (Merchant, error) {
	chr, err := character.Random()
	if err != nil {
		err = fmt.Errorf("Could not generate merchant: %w", err)
		return Merchant{}, err
	}
	chr.Profession, _ = profession.ByName("merchant")

	gd := goods.GenerateMerchantGoods(10, 30, originTown.Resources)

	merchant := Merchant{
		Character: chr,
		Goods:     gd,
	}

	return merchant, nil
}

// Random returns a complete random merchant
func Random() (Merchant, error) {
	originArea, err := geography.Generate()
	if err != nil {
		err = fmt.Errorf(merchantRandomGenerationError, err)
		return Merchant{}, err
	}

	originCulture, err := culture.Generate(originArea)
	if err != nil {
		err = fmt.Errorf(merchantRandomGenerationError, err)
		return Merchant{}, err
	}

	originTown, err := town.Generate("metropolis", originArea, originCulture)
	if err != nil {
		err = fmt.Errorf(merchantRandomGenerationError, err)
		return Merchant{}, err
	}

	merchant, err := Generate(originTown)
	if err != nil {
		err = fmt.Errorf(merchantRandomGenerationError, err)
		return Merchant{}, err
	}

	return merchant, nil
}
