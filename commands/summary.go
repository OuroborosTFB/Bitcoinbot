package commands

import (
	"fmt"
	"math"

	"github.com/OuroborosTFB/Bitcoinbot/utils"
)

func GetSummary() (string, string, error) {
	price, err := utils.GetPriceAPI()
	last := price.Last
	buy := price.Buy
	his := ((last - buy) / buy) * 100
	if !math.Signbit(float64(his)) {
		return fmt.Sprintf("%.2f", price.Last), "%" + fmt.Sprintf("%.2f", ((last-buy)/buy)*100), err
	} else {
		return fmt.Sprintf("%.2f", price.Last), "-%" + fmt.Sprintf("%.2f", -1*((last-buy)/buy)*100), err
	}
}
