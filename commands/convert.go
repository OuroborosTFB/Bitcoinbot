package commands

import (
	"github.com/OuroborosTFB/Bitcoinbot/utils"
)

func GetConvertedPrice(value int) (string, error) {
	price, err := utils.GetConvertedPriceAPI(value)
	if err != nil {

	}
	return price, err
}
