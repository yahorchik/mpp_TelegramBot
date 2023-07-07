package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/patrickmn/go-cache"
	"log"
	lc "mpp_TelegramBot/internal/pkg/cache"
)

func InitBot(token string) (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = true
	return bot, nil
}

func FindMessage(update tgbotapi.Update, bot *tgbotapi.BotAPI, c *cache.Cache) {
	var msg tgbotapi.MessageConfig
	//	cache.MsgToCache(update.Text)
	if tgbotapi.Message.IsCommand(*update.Message) == true {
		err := findCommand(update, bot)
		if err != nil {
			log.Fatal(err)
		}
	}
	err := lc.MsgToCache(update.Message, c)
	if err != nil {
		log.Fatal(err)
	}
	msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Message saved.")
	bot.Send(msg)
}

func findCommand(update tgbotapi.Update, bot *tgbotapi.BotAPI) error {
	var msg tgbotapi.MessageConfig
	switch update.Message.Command() {
	case "start":
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Hello! I'm message_cache_bot. Send me a message!")
	case "cache":
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "pizdec")
	case "db":
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "really pizdec")
	default:
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "error: unknown command.")
	}
	_, err := bot.Send(msg)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
