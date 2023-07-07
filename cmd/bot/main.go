package main

import (
	"log"
	"mpp_TelegramBot/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
