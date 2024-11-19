package track

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"fmt"
	"time"
)

// SendTrackRun tracks the run activity
func SendTrackRun(bot *shared.Bot, userName, content string, chatID int64) error {
	now, err := shared.GetNow()
	if err != nil {
		return shared.SendMessage(bot, chatID, err.Error())
	}

	nowStr := now.Format(time.RFC3339)

	userActivity := storage.UserActivity{
		ID:        storage.GenerateActivityItemID(now, userName, shared.ToothBrush),
		Name:      userName,
		Activity:  shared.ToothBrush,
		CreatedAt: nowStr,
		Content:   content,
	}

	err = storage.Create(userActivity)
	if err != nil {
		return shared.SendMessage(bot, chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	}

	message := "mi papacho el más usain vol 🏃‍♂️"
	if content != "" {
		message = fmt.Sprintf("uy mi papacho corrió %s? lo iba robar un negro o qué manito. anwy congrats", content)
	}

	return shared.SendMessage(bot, chatID, message)
}
