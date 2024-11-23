package track

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"time"

	flag "github.com/spf13/pflag"
)

func sendTrackGym(bot *shared.Bot, userName, content string, chatID int64) error {
	now, err := shared.GetNow()
	if err != nil {
		return shared.SendMessage(bot, chatID, err.Error())
	}

	nowStr := now.Format(time.RFC3339)

	userActivity := shared.UserActivity{
		ID:        shared.GenerateActivityItemID(now, userName, shared.Gym),
		Name:      userName,
		Activity:  shared.Gym,
		CreatedAt: nowStr,
		Content:   content,
	}

	var time string
	var muscle string

	flag.StringVarP(&time, "time", "t", "", "time you were exercising")
	flag.StringVarP(&muscle, "muscle", "p", "", "muscles you exercised splitted by comma [bicep,back,shoulder]")
	flag.Parse()

	if time == "" {
		return shared.SendMessage(bot, chatID, "eh pero cuánto tiempo te ejercitaste sapa inmunda, mandame el -time")
	}

	if muscle == "" {
		return shared.SendMessage(bot, chatID, "eh pero hiciste chisme al fallo o q 🐸? mandame el -muscle sapa. Ej: -muscle bicep,pecho,jeta")
	}

	err = storage.Create(userActivity)
	if err != nil {
		return shared.SendMessage(bot, chatID, "algo falló mi fafá: "+err.Error())
	}

	return shared.SendMessage(bot, chatID, "isss mi papacho el próximo cbum ve")
}
