package cache

import (
	"encoding/hex"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/patrickmn/go-cache"
	"github.com/yahorchik/mpp_TelegramBot/internal/database"
	lc "github.com/yahorchik/mpp_TelegramBot/internal/pkg/cache"
	"github.com/yahorchik/mpp_TelegramBot/internal/pkg/repositories/gen/postgres/public/model"
	"github.com/yahorchik/mpp_TelegramBot/internal/pkg/repositories/gen/postgres/public/table"
	"google.golang.org/protobuf/proto"
	"log"
	"os/exec"
	"strconv"
	"time"
)

func ShowMessage(id int64, bot *tgbotapi.BotAPI) error {
	var msg tgbotapi.MessageConfig
	msg = tgbotapi.NewMessage(id, "История сохраненных сообщений:")
	bot.Send(msg)
	var msgtext string
	for _, item := range lc.Cache.Items() {
		minfo, ok := item.Object.(lc.Message)
		if !ok {
			log.Fatal()
		}
		tm := time.Unix(int64(minfo.Data), 0)
		timeStr := tm.Format("2006-01-02T15:04:05")
		msgtext = fmt.Sprintf("User: %v. \nMessage: %v. \nDate: %v.",
			strconv.FormatInt(minfo.User, 10), minfo.Text, timeStr)
		msg = tgbotapi.NewMessage(id, msgtext)
		_, err := bot.Send(msg)
		if err != nil {
			log.Fatal(err)
		}
		md := &model.MessageInfo{
			UserID:      proto.String(strconv.FormatInt(minfo.User, 10)),
			MessageText: proto.String(minfo.Text),
			MessageDate: &tm,
		}
		md1 := &model.UserInfo{
			UserID:        strconv.FormatInt(minfo.User, 10),
			UserNickname:  nil,
			UserFirstname: nil,
			UserLastname:  nil,
		}
		stmt1 := table.UserInfo.INSERT(table.UserInfo.AllColumns).MODEL(md1)
		stmt := table.MessageInfo.INSERT(table.MessageInfo.AllColumns).MODEL(md)
		_, err = stmt1.Exec(database.DB)
		if err != nil {
			log.Println(err)
		}
		_, err = stmt.Exec(database.DB)
		if err != nil {
			log.Println(err)
		}
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
