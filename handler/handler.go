package handler

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/OuroborosTFB/Bitcoinbot/commands"
	tb "gopkg.in/tucnak/telebot.v2"
)

func LoadHandler(b *tb.Bot) map[string]func(m *tb.Message) {
	commandMap := make(map[string]func(m *tb.Message))

	commandMap["/price"] = func(m *tb.Message) {
		res, _ := commands.GetPrice()
		b.Send(m.Chat, "Актуальная цена BTC на данный момент в $ "+res)
	}

	commandMap["/convert"] = func(m *tb.Message) {
		re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`) // regex выражение для отлова чисел из строки
		submatchall := re.FindAllString(m.Text, -1)           // поиск по m.Text (текст сообщения) с помощью нашего regex

		value := 0                //значение int
		numberString := "0"       //значение string
		if len(submatchall) > 0 { //если найдено больше 0 чисел, то делаем следующее
			numberString = submatchall[0]         //берём первое число, которое нашлось
			value, _ = strconv.Atoi(numberString) //конвертируем его в int
		}

		res, _ := commands.GetConvertedPrice(value) //вызываем команду для получения цены

		botMessage := strings.Join([]string{"За", numberString + "$", "можно купить:", res + "₿"}, " ") //формируем строку для ответа бота
		b.Send(m.Chat, botMessage)                                                                      //отправляем сообщение от бота
	}
	return commandMap
}
