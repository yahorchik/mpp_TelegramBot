package app

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	b "github.com/yahorchik/mpp_TelegramBot/internal/app/bot"
	"github.com/yahorchik/mpp_TelegramBot/internal/database"
	lc "github.com/yahorchik/mpp_TelegramBot/internal/pkg/cache"
	"github.com/yahorchik/mpp_TelegramBot/internal/pkg/config"
	eCache "github.com/yahorchik/mpp_TelegramBot/internal/pkg/events/cache"
	"github.com/yahorchik/mpp_TelegramBot/internal/pkg/events/start"
	"github.com/yahorchik/mpp_TelegramBot/internal/pkg/repositories"
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
	if err := SearchUpdate(ctx, bot, cache, dbconn); err != nil {
		return err
	}
	return nil
}

func SearchUpdate(ctx context.Context, bot *tgbotapi.BotAPI, c *lc.Cache, dbconn *database.DBconn) error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}
	for update := range updates {
		if update.Message != nil { // If we got a message
			err := FindMessage(ctx, update, bot, c, dbconn)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func FindMessage(ctx context.Context, update tgbotapi.Update, bot *tgbotapi.BotAPI, c *lc.Cache, dbconn *database.DBconn) error {
	if update.Message.IsCommand() == true {
		err := findCommand(ctx, update, bot, c, dbconn)
		if err != nil {
			return err
		}
	} else {
		err := eCache.MsgToCache(update.Message, c)
		if err != nil {
			return err
		}
		err = b.SendMessage(bot, update, "Message saved")
		if err != nil {
			return err
		}
	}
	return nil
}

func findCommand(ctx context.Context, update tgbotapi.Update, bot *tgbotapi.BotAPI, c *lc.Cache, dbconn *database.DBconn) error {
	switch update.Message.Command() {
	case "start":
		start.Message(update.Message.Chat.ID, bot)
	case "cache":
		err := eCache.ShowMessage(ctx, update.Message.Chat.ID, update, bot, c, dbconn)
		if err != nil {
			return err
		}
	case "db":
		err := b.SendMessage(bot, update, "DEBUG MOD")
		if err != nil {
			return err
		}
		repositories.GetForDB(ctx, *dbconn)
	default:
		err := b.SendMessage(bot, update, "error: unknown command.")
		if err != nil {
			return err
		}
	}
	return nil
}
