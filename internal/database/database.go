package database

import (
	"database/sql"
	"fmt"
	"github.com/yahorchik/mpp_TelegramBot/internal/pkg/config"
	"log"
)

type DBconn struct {
	DB *sql.DB
}

func ConnectDB(cfg *config.Config) (*DBconn, error) {
	//	conn, err := sql.Open("postgres", config.Cfg.DB.GetUrl())
	conn, err := sql.Open("postgres", cfg.DB.GetURL())
	if err != nil {
		err = fmt.Errorf("failed on sql.Open: %w", err)
		log.Printf("database.ConnectDB: %v", err)
		return nil, err
	}

	return &DBconn{DB: conn}, nil
}
