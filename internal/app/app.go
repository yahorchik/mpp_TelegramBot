package app

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	. "mpp_TelegramBot/internal/app/bot"
	. "mpp_TelegramBot/internal/app/setup"
	"mpp_TelegramBot/internal/pkg/cache"
)

func Run() error {
	token, err := SetupToken()
	if err != nil {
		log.Fatal(err)
	}
	bot, err := InitBot(token)
	if err != nil {
		log.Fatal(err)
	}
	c, err := cache.InitCache()
	if err != nil {
		log.Fatal(err)
	}
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			FindMessage(update, bot, c)
		}
	}
	return nil
}
