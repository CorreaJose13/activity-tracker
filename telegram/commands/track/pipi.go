package track

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"fmt"
	"time"
)

// SendTrackPipi tracks the pipi activity
func SendTrackPipi(bot *shared.Bot, userName, content string, chatID int64) error {
	now, err := shared.GetNow()
	if err != nil {
		return shared.SendMessage(bot, chatID, err.Error())
	}

	nowStr := now.Format(time.RFC3339)

	userActivity := storage.UserActivity{
		ID:        storage.GenerateActivityItemID(now, userName, shared.Pipi),
		Name:      userName,
		Activity:  shared.Pipi,
		CreatedAt: nowStr,
	}

	err = storage.Create(userActivity)
	if err != nil {
		return shared.SendMessage(bot, chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	}

	if userName == shared.Valery {
		return shared.SendMessage(bot, chatID, "Epa, buena esa mionsita 😎")
	}

	return shared.SendMessage(bot, chatID, "Epa, buena esa mionsito 😎")
}
