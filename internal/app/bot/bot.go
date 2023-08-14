package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	eCache "github.com/yahorchik/mpp_TelegramBot/internal/pkg/events/cache"
	"github.com/yahorchik/mpp_TelegramBot/internal/pkg/events/start"
	"github.com/yahorchik/mpp_TelegramBot/internal/pkg/repositories"
)

type Message struct {
	Data int
	text string
	user int64
}

func InitBot(token string) (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return bot, err
	}
	bot.Debug = false
	return bot, nil
}

func SearchUpdate(bot *tgbotapi.BotAPI) error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}
	for update := range updates {
		if update.Message != nil { // If we got a message
			err := FindMessage(update, bot)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func FindMessage(update tgbotapi.Update, bot *tgbotapi.BotAPI) error {
	var msg tgbotapi.MessageConfig
	if update.Message.IsCommand() == true {
		err := findCommand(update, bot)
		if err != nil {
			return err
		}
	} else {
		err := eCache.MsgToCache(update.Message)
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

func findCommand(update tgbotapi.Update, bot *tgbotapi.BotAPI) error {
	var msg tgbotapi.MessageConfig
	switch update.Message.Command() {
	case "start":
		start.StartMessage(update.Message.Chat.ID, bot)
	case "cache":
		err := eCache.ShowMessage(update.Message.Chat.ID, update, bot)
		if err != nil {
			return err
		}
	case "db":
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "really pizdec")
		_, err := bot.Send(msg)
		if err != nil {
			return err
		}
		repositories.GetForDB()
	default:
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "error: unknown command.")
		_, err := bot.Send(msg)
		if err != nil {
			return err
		}
	}
	return nil
}
