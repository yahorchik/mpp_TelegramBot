package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	eCache "github.com/yahorchik/mpp_TelegramBot/internal/pkg/events/cache"
	"github.com/yahorchik/mpp_TelegramBot/internal/pkg/events/start"
	"log"
)

type Message struct {
	Data int
	text string
	user int64
}

func InitBot(token string) (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = false
	return bot, nil
}

func FindMessage(update tgbotapi.Update, bot *tgbotapi.BotAPI) error {
	var msg tgbotapi.MessageConfig
	if update.Message.IsCommand() == true {
		err := findCommand(update, bot)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := eCache.MsgToCache(update.Message)
		if err != nil {
			log.Fatal(err)
		}
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Message saved.")
		bot.Send(msg)
	}
	return nil
}

func findCommand(update tgbotapi.Update, bot *tgbotapi.BotAPI) error {
	var msg tgbotapi.MessageConfig
	switch update.Message.Command() {
	case "start":
		start.StartMessage(update.Message.Chat.ID, bot)
	case "cache":
		err := eCache.ShowMessage(update.Message.Chat.ID, bot)
		if err != nil {
			log.Fatal(err)
		}
	case "db":
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "really pizdec")
		_, err := bot.Send(msg)
		if err != nil {
			log.Fatal(err)
		}
	default:
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "error: unknown command.")
		_, err := bot.Send(msg)
		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}
