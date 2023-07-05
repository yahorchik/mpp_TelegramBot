package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	. "mpp_TelegramBot/internal/app/setup"
	"mpp_TelegramBot/internal/pkg/cache"
)

func Run() error {
	token, err := SetupToken()
	if err != nil {
		log.Fatal(err)
	}
	initBot(token)
	return nil
}

func initBot(token string) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = true
	updatesBot(bot)
}

func updatesBot(bot *tgbotapi.BotAPI) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			messageBot(update.Message, bot)
			//		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			//		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			//		msg.ReplyToMessageID = update.Message.MessageID

			//		bot.Send(msg)
		}
	}
}

func messageBot(update *tgbotapi.Message, bot *tgbotapi.BotAPI) {
	var msg tgbotapi.MessageConfig
	cache.MsgToCache(update.Text)
	switch update {
	case "start":
		msg = tgbotapi.NewMessage(update.Chat.ID, "Hello! I'm Message_cache_bot. Privet Sanya")
	case "/cache":
		msg = tgbotapi.NewMessage(update.Chat.ID, "Бля, попроси что-нибудь попроще")
	default:
		msg = tgbotapi.NewMessage(update.Chat.ID, update.Text)
	}
	//	msg.ReplyToMessageID = update.MessageID
	bot.Send(msg)
}
