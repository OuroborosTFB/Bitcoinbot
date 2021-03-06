package main

import (
	"log"
	"time"

	"github.com/OuroborosTFB/Bitcoinbot/config"
	"github.com/OuroborosTFB/Bitcoinbot/handler"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	startBot()
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
