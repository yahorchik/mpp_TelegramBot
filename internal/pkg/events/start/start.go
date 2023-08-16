package start

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func Message(id int64, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(id, "Hello! I'm message_cache_bot. Send me a message!")

	if _, err := bot.Send(msg); err != nil {
		log.Println(err)
	}
}
