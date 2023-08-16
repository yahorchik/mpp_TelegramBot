package bot

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/yahorchik/mpp_TelegramBot/internal/database"
	lc "github.com/yahorchik/mpp_TelegramBot/internal/pkg/cache"
	eCache "github.com/yahorchik/mpp_TelegramBot/internal/pkg/events/cache"
	"github.com/yahorchik/mpp_TelegramBot/internal/pkg/events/start"
	"github.com/yahorchik/mpp_TelegramBot/internal/pkg/repositories"
)

type Message struct {
	Data int
	Text string
	User int64
}

func InitBot(botToken string) (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {

		return nil, err
	}
	bot.Debug = false
	return bot, nil
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
	var msg tgbotapi.MessageConfig
	if update.Message.IsCommand() == true {
		err := findCommand(ctx, update, bot, c, dbconn)
		if err != nil {
			return err
		}
	} else {
		fmt.Println(&c)
		err := eCache.MsgToCache(update.Message, c)
		if err != nil {
			return err
		}
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Message saved.")
		_, err = bot.Send(msg)
		if err != nil {
			return err
		}
	}
	return nil
}

func findCommand(ctx context.Context, update tgbotapi.Update, bot *tgbotapi.BotAPI, c *lc.Cache, dbconn *database.DBconn) error {
	var msg tgbotapi.MessageConfig
	switch update.Message.Command() {
	case "start":
		start.Message(update.Message.Chat.ID, bot)
	case "cache":
		err := eCache.ShowMessage(ctx, update.Message.Chat.ID, update, bot, c, dbconn)
		if err != nil {
			return err
		}
	case "db":
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "really pizdec")
		_, err := bot.Send(msg)
		if err != nil {
			return err
		}
		repositories.GetForDB(ctx, *dbconn)
	default:
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "error: unknown command.")
		_, err := bot.Send(msg)
		if err != nil {
			return err
		}
	}
	return nil
}
