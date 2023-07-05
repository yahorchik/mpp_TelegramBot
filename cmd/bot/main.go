package main

import (
	"log"
	"mpp_TelegramBot/internal/app/bot"
)

func main() {
	if err := bot.Run(); err != nil {
		log.Fatal(err)
	}
}
