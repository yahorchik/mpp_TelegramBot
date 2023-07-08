package main

import (
	"context"
	"github.com/yahorchik/mpp_TelegramBot/internal/app"
	"log"
)

func main() {
	if err := app.Run(context.Background()); err != nil {
		log.Fatal(err)
	}
}
