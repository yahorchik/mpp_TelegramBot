package database

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/yahorchik/mpp_TelegramBot/internal/pkg/config"
	"log"
)

var DB *sql.DB

func ConnectDB(ctx context.Context) {
	conn, err := sql.Open("postgres", config.Cfg.DB.GetUrl())
	if err != nil {
		log.Fatal(err)
	}
	DB = conn
}
