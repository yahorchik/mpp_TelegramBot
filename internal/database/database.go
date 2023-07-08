package database

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/yahorchik/mpp_TelegramBot/internal/pkg/config"
	"log"
)

var DB *pgx.Conn

func ConnectDB(ctx context.Context) {
	conn, err := pgx.Connect(ctx, config.Cfg.DB.GetUrl())
	if err != nil {
		log.Fatal(err)
	}
	DB = conn
}
