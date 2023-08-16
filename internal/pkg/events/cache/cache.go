package cache

import (
	"context"
	"encoding/hex"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/patrickmn/go-cache"
	"github.com/yahorchik/mpp_TelegramBot/internal/database"
	lc "github.com/yahorchik/mpp_TelegramBot/internal/pkg/cache"
	rep "github.com/yahorchik/mpp_TelegramBot/internal/pkg/repositories"
	"github.com/yahorchik/mpp_TelegramBot/internal/pkg/repositories/gen/postgres/public/model"
	"google.golang.org/protobuf/proto"
	"log"
	"os/exec"
	"strconv"
	"time"
)

type Message struct {
	Data int
	Text string
	User int64
}

func ShowMessage(ctx context.Context, id int64, update tgbotapi.Update, bot *tgbotapi.BotAPI, c *lc.Cache, dbconn *database.DBconn) error {
	var msg tgbotapi.MessageConfig
	modelUser := &model.UserInfo{
		UserID:        strconv.FormatInt(id, 10),
		UserNickname:  &update.Message.Chat.UserName,
		UserFirstname: &update.Message.Chat.FirstName,
		UserLastname:  &update.Message.Chat.LastName,
	}
	if c.C.ItemCount() == 0 {
		msg = tgbotapi.NewMessage(id, "Сохраненных сообщений нет.")
		_, err := bot.Send(msg)
		if err != nil {
			log.Println(err)
		}
		return nil
	}
	msg = tgbotapi.NewMessage(id, "История сохраненных сообщений:")
	_, err := bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
	var msgtext string
	modelsMessages := make([]*model.MessageInfo, 0, len(c.C.Items()))

	items := c.C.Items()
	for key, item := range items {
		minfo, ok := item.Object.(Message)
		if !ok {
			log.Println(err)
		}

		//x, _ := c.C.Get(key)
		//minfo, ok := x.(Message)
		//if !ok {
		//	log.Println("aboba")
		//}

		if minfo.User != id {
			continue
		}
		tm := time.Unix(int64(minfo.Data), 0)
		timeStr := tm.Format("2006-01-02T15:04:05")
		msgtext = fmt.Sprintf("User: %v. \nMessage: %v. \nDate: %v.",
			strconv.FormatInt(minfo.User, 10), minfo.Text, timeStr)
		msg = tgbotapi.NewMessage(id, msgtext)
		_, err := bot.Send(msg)
		if err != nil {
			log.Println(err)
		}
		md := &model.MessageInfo{
			UserID:      proto.String(strconv.FormatInt(minfo.User, 10)),
			MessageText: proto.String(minfo.Text),
			MessageDate: &tm,
		}
		c.C.Delete(key)
		modelsMessages = append(modelsMessages, md)
	}

	err = rep.SendToDB(ctx, modelUser, modelsMessages, dbconn)
	if err != nil {
		log.Println(err)
	}
	return nil
}

func MsgToCache(msg *tgbotapi.Message, c *lc.Cache) error {
	var message Message
	key, err := exec.Command("uuidgen").Output()
	if err != nil {
		//	return nil, err
		fmt.Println("JOPA AHAHAHAHAH")
	}
	message.Data = msg.Date
	message.Text = msg.Text
	message.User = msg.Chat.ID
	fmt.Println(&c)
	c.C.Set(hex.EncodeToString(key), message, cache.DefaultExpiration)
	for _, item := range c.C.Items() {
		minfo, ok := item.Object.(Message)
		if !ok {
			log.Println(err)
		}
		fmt.Println(minfo)
	}
	return nil
}
