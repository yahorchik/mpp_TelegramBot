package repositories

import (
	"github.com/go-jet/jet/v2/postgres"
	"github.com/yahorchik/mpp_TelegramBot/internal/database"
	"github.com/yahorchik/mpp_TelegramBot/internal/pkg/repositories/gen/postgres/public/model"
	"github.com/yahorchik/mpp_TelegramBot/internal/pkg/repositories/gen/postgres/public/table"
	"log"
)

func SendToDB(userinfo *model.UserInfo, msginfo []*model.MessageInfo) error {
	stmt1 := table.UserInfo.INSERT(table.UserInfo.AllColumns).MODEL(userinfo).ON_CONFLICT(table.UserInfo.UserID).
		DO_UPDATE(
			postgres.SET(
				table.UserInfo.UserFirstname.SET(postgres.String(*userinfo.UserFirstname)),
				table.UserInfo.UserLastname.SET(postgres.String(*userinfo.UserLastname)),
				table.UserInfo.UserNickname.SET(postgres.String(*userinfo.UserNickname)),
			).WHERE(table.UserInfo.UserID.EQ(postgres.String(userinfo.UserID))),
		)
	stmt := table.MessageInfo.INSERT(table.MessageInfo.AllColumns).MODELS(msginfo)
	_, err := stmt1.Exec(database.DB)
	if err != nil {
		log.Println(err)
	}
	_, err = stmt.Exec(database.DB)
	if err != nil {
		log.Println(err)
	}
	return nil
}
