package cache

import (
	"encoding/hex"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/patrickmn/go-cache"
	"os/exec"
	"time"
)

type Message struct {
	data int
	text string
	user int64
}

func InitCache() (*cache.Cache, error) {
	c := cache.New(5*time.Minute, 10*time.Minute)
	return c, nil
}

func MsgToCache(msg *tgbotapi.Message, c *cache.Cache) error {
	var message Message
	key, err := exec.Command("uuidgen").Output()
	if err != nil {
		return err
	}
	message.data = msg.Date
	message.text = msg.Text
	message.user = msg.Chat.ID
	c.Set(hex.EncodeToString(key), message, cache.DefaultExpiration)
	return nil
}

/*c.Items()
	for dick, item := range c.Items() {
		log.Println(dick, item)
		switch item.Object.(type) {
		case int:

		}
	}
}
*/
