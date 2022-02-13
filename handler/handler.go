package handler

import (
	"github.com/OuroborosTFB/Bitcoinbot/commands"
	tb "gopkg.in/tucnak/telebot.v2"
)

func LoadHandler(b *tb.Bot) map[string]func(m *tb.Message) {
	commandMap := make(map[string]func(m *tb.Message))

	commandMap["/price"] = func(m *tb.Message) {
		res, _ := commands.GetPrice()
		b.Send(m.Chat, "BTC's Current price is: U$S "+res)
	}

	commandMap["/convert"] = func(m *tb.Message) {
		res, g, _ := commands.GetConvertedPrice()
		b.Send(m.Chat, "Cтоимость", "BTC в $ на данный момент"+res)
		b.Send(m.Chat, g)
	}

	return commandMap
}
