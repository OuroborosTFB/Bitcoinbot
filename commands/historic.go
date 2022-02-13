package commands

import (
	"fmt"
	"math"

	"github.com/OuroborosTFB/Bitcoinbot/utils"
	tb "gopkg.in/tucnak/telebot.v2"
)

func GetHistoric() (string, *tb.Animation, error) {
	price, err := utils.GetApiCall()
	last := price.Last
	buy := price.Buy
	his := ((last - buy) / buy) * 100
	if !math.Signbit(float64(his)) {
		g := &tb.Animation{File: tb.FromURL("https://i.pinimg.com/originals/e4/38/99/e4389936b099672128c54d25c4560695.gif")}
		return "%" + fmt.Sprintf("%.2f", ((last-buy)/buy)*100), g, err
	} else {
		g := &tb.Animation{File: tb.FromURL("http://www.brainlesstales.com/bitcoin-assets/images/fan-versions/2015-01-osEroUI.gif")}
		return "-%" + fmt.Sprintf("%.2f", -1*((last-buy)/buy)*100), g, err
	}
}
