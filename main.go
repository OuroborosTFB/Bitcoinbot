package main

import (
	"fmt"
	"log"
	"time"

	"github.com/OuroborosTFB/Bitcoinbot/config"
	"github.com/OuroborosTFB/Bitcoinbot/handler"
	"github.com/OuroborosTFB/Bitcoinbot/utils"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	//startBot()
	convertedprice, _ := utils.GetConvertedPriceAPI(55165)
	fmt.Println(convertedprice)

}

func startBot() {

	b, err := tb.NewBot(tb.Settings{

		Token:  config.LoadConfig().Token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	for k, v := range handler.LoadHandler(b) {
		b.Handle(k, v)
		log.Println(k + "✅ Loaded!")
	}

	b.Start()

}
