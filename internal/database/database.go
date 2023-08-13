package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/yahorchik/mpp_TelegramBot/internal/pkg/config"
)

var DB *sql.DB

func ConnectDB() error {
	conn, err := sql.Open("postgres", config.Cfg.DB.GetUrl())
	if err != nil {
		return err
	}
	DB = conn
	return nil
}
