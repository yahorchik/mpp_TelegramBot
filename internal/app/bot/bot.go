package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
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

func SendMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update, text string) error {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	_, err := bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}
