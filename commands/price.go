package commands

import (
	"fmt"

	"github.com/OuroborosTFB/Bitcoinbot/utils"
)

func GetPrice() (string, error) {
	price, err := utils.GetPriceAPI()
	return fmt.Sprintf("%.2f", price.Buy), err
}
