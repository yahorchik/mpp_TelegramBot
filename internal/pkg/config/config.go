package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

var Cfg *Config

func SetupConfig() error {
	buf := &Config{}
	var data, err = ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(data, buf)
	if err != nil {
		log.Fatal(err)
	}
	Cfg = buf
	return nil
}

func (c *Database) GetUrl() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", c.User, c.Password, c.Host, c.Port, c.DB, c.SSLmod)
}
