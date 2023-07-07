package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/patrickmn/go-cache"
	"log"
	lc "mpp_TelegramBot/internal/pkg/cache"
)

type Message struct {
	data int
	text string
	user int64
}

func InitBot(token string) (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = false
	return bot, nil
}

func FindMessage(update tgbotapi.Update, bot *tgbotapi.BotAPI, c *cache.Cache) error {
	var msg tgbotapi.MessageConfig
	//	cache.MsgToCache(update.Text)
	if update.Message.IsCommand() == true {
		err := findCommand(update, bot, c)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := lc.MsgToCache(update.Message, c)
		if err != nil {
			log.Fatal(err)
		}
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Message saved.")
		bot.Send(msg)
	}
	return nil
}

func findCommand(update tgbotapi.Update, bot *tgbotapi.BotAPI, c *cache.Cache) error {
	var msg tgbotapi.MessageConfig
	switch update.Message.Command() {
	case "start":
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Hello! I'm message_cache_bot. Send me a message!")
	case "cache":
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "pizdec")
	case "db":
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "really pizdec")
	case "show":
		lc.ShowMessage(c, update.Message.Chat.ID, bot)
	default:
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "error: unknown command.")
	}
	_, err := bot.Send(msg)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

/*func sshowMessage(c *cache.Cache, id int64, bot *tgbotapi.BotAPI) error {
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
		msgtext = "User: " + strconv.FormatInt(minfo.user, 10) + ". /nMessage: " + minfo.text + ". /nData: " + timeStr + "."
		msg = tgbotapi.NewMessage(id, msgtext)
		bot.Send(msg)
	}
	return nil
}
*/
