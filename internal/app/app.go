package app

import (
	"context"
	b "github.com/yahorchik/mpp_TelegramBot/internal/app/bot"
	"github.com/yahorchik/mpp_TelegramBot/internal/database"
	lc "github.com/yahorchik/mpp_TelegramBot/internal/pkg/cache"
	"github.com/yahorchik/mpp_TelegramBot/internal/pkg/config"
)

func Run(ctx context.Context) error {
	//Setup app config
	cfg, err := config.SetupConfig()
	if err != nil {
		return err
	}

	//Init Cache
	//var cache *lc.Cache
	cache := lc.InitCache()

	//Connect to DB
	dbconn, err := database.ConnectDB(cfg)
	if err != nil {
		return err
	}
	bot, err := b.InitBot(cfg.BotToken)
	if err != nil {
		return err
	}
	if err := b.SearchUpdate(ctx, bot, cache, dbconn); err != nil {
		return err
	}
	return nil
}
