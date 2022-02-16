package commands

import (
	"github.com/OuroborosTFB/Bitcoinbot/utils"
)

func GetConvertedPrice() (string, error) {
	price, err := utils.GetConvertedPriceAPI(12)
	if err != nil {

	}
	return price, err
}
