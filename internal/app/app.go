package app

import (
	"context"
	b "github.com/yahorchik/mpp_TelegramBot/internal/app/bot"
	"github.com/yahorchik/mpp_TelegramBot/internal/database"
	"github.com/yahorchik/mpp_TelegramBot/internal/pkg/cache"
	"github.com/yahorchik/mpp_TelegramBot/internal/pkg/config"
)

func Run(ctx context.Context) error {
	if err := config.SetupConfig(); err != nil {
		return err
	}
	cache.InitCache()
	err := database.ConnectDB()
	if err != nil {
		return err
	}
	bot, err := b.InitBot(config.Cfg.BotToken)
	if err != nil {
		return err
	}
	if err := b.SearchUpdate(bot); err != nil {
		return err
	}
	return nil
}
