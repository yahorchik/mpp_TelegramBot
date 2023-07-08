package app

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	b "github.com/yahorchik/mpp_TelegramBot/internal/app/bot"
	"github.com/yahorchik/mpp_TelegramBot/internal/database"
	"github.com/yahorchik/mpp_TelegramBot/internal/pkg/cache"
	"github.com/yahorchik/mpp_TelegramBot/internal/pkg/config"
	"log"
)

func Run(ctx context.Context) error {
	err := config.SetupConfig()
	if err != nil {
		log.Fatal(err)
	}
	cache.InitCache()
	database.ConnectDB(ctx)
	bot, err := b.InitBot(config.Cfg.BotToken)
	if err != nil {
		log.Fatal(err)
	}
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}
	for update := range updates {
		if update.Message != nil { // If we got a message
			err := b.FindMessage(update, bot)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	return nil
}
