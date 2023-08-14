package cache

import (
	"encoding/hex"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/patrickmn/go-cache"
	lc "github.com/yahorchik/mpp_TelegramBot/internal/pkg/cache"
	rep "github.com/yahorchik/mpp_TelegramBot/internal/pkg/repositories"
	"github.com/yahorchik/mpp_TelegramBot/internal/pkg/repositories/gen/postgres/public/model"
	"google.golang.org/protobuf/proto"
	"log"
	"os/exec"
	"strconv"
	"time"
)

func ShowMessage(id int64, update tgbotapi.Update, bot *tgbotapi.BotAPI) error {
	var modelsMessages []*model.MessageInfo
	var msg tgbotapi.MessageConfig
	modelUser := &model.UserInfo{
		UserID:        strconv.FormatInt(id, 10),
		UserNickname:  &update.Message.Chat.UserName,
		UserFirstname: &update.Message.Chat.FirstName,
		UserLastname:  &update.Message.Chat.LastName,
	}
	if lc.Cache.ItemCount() == 0 {
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
	for key, item := range lc.Cache.Items() {
		minfo, ok := item.Object.(lc.Message)
		if !ok {
			log.Println(err)
		}
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
		lc.Cache.Delete(key)
		modelsMessages = append(modelsMessages, md)
	}
	err = rep.SendToDB(modelUser, modelsMessages)
	if err != nil {
		log.Println(err)
	}
	return nil
}
func MsgToCache(msg *tgbotapi.Message) error {
	var message lc.Message
	key, err := exec.Command("uuidgen").Output()
	if err != nil {
		return err
	}
	message.Data = msg.Date
	message.Text = msg.Text
	message.User = msg.Chat.ID
	lc.Cache.Set(hex.EncodeToString(key), message, cache.DefaultExpiration)

	return nil
}
