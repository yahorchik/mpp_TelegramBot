package setup

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

func SetupToken() (string, error) {
	var authBot AuthBot
	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(data, &authBot)
	if err != nil {
		log.Fatal(err)
	}
	return authBot.BotToken, nil
}
