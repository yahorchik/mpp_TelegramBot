package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type Config struct {
	BotToken string    `yaml:"botToken"`
	DB       *Database `yaml:"dataBase"`
}

type Database struct {
	User       string `yaml:"user"`
	Password   string `yaml:"password"`
	DB         string `yaml:"db"`
	Port       string `yaml:"port"`
	Host       string `yaml:"host"`
	Scheme     string `yaml:"scheme"`
	MaxConnect int    `yaml:"maxConnect"`
	Timeout    int    `yaml:"timeout"`
	SSLmod     string `yaml:"sslmod"`
}

func SetupConfig() (*Config, error) {
	configPath := "config.yaml"
	//configPath := os.Getenv("CONFIG_PATH")
	//if configPath == "" {
	//	log.Fatal("CONFIG_PATH is not set")
	//}

	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
		return nil, err
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
		return nil, err
	}

	return &cfg, nil
}

func (c *Database) GetURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", c.User, c.Password, c.Host, c.Port, c.DB, c.SSLmod)
}
