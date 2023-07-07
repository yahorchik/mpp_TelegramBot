package cache

import (
	"encoding/hex"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/patrickmn/go-cache"
	"log"
	"os/exec"
	"strconv"
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

/*func ShowMessage(c *cache.Cache, id int64, bot *tgbotapi.BotAPI) error {
	var msg tgbotapi.MessageConfig
	var msgtext string
	for _, item := range c.Items() {
		minfo, ok := item.Object.(Message)
		if !ok {
			log.Fatal()
		}
		msgtext = "User: " + strconv.FormatInt(minfo.user, 10) + ". Message: " + minfo.text + ". Data: " + strconv.FormatInt(int64(minfo.data), 10) + "."
		msg = tgbotapi.NewMessage(id, msgtext)
		bot.Send(msg)
	}
	return nil
}
*/
func ShowMessage(c *cache.Cache, id int64, bot *tgbotapi.BotAPI) {
	var msg tgbotapi.MessageConfig
	var msgtext string
	for _, item := range c.Items() {
		minfo, ok := item.Object.(Message)
		if !ok {
			log.Fatal()
		}
		fmt.Println(minfo.data)
		tm := time.Unix(int64(minfo.data), 0)
		fmt.Println(tm.Date())
		timeStr := tm.Format("2006-01-02T15:04:05")
		fmt.Println(timeStr)
		msgtext = "User: " + strconv.FormatInt(minfo.user, 10) + "." + " Message: " + minfo.text + ". Date: " + timeStr + "."
		msg = tgbotapi.NewMessage(id, msgtext)
		bot.Send(msg)
	}
}
