package start

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func StartMessage(id int64, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(id, "Hello! I'm message_cache_bot. Send me a message!")
	_, err := bot.Send(msg)
	if err != nil {
		log.Fatal(err)
	}
}
