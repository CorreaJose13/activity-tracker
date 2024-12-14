package track

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"fmt"
	"time"
)

// SendTrackRun tracks the run activity
func SendTrackRun(client *shared.Client, userName, content string, chatID int64) error {
	now, err := shared.GetNow()
	if err != nil {
		return client.SendMessage(chatID, err.Error())
	}

	nowStr := now.Format(time.RFC3339)

	userActivity := shared.UserActivity{
		ID:        shared.GenerateActivityItemID(now, userName, shared.Run),
		Name:      userName,
		Activity:  shared.Run,
		CreatedAt: nowStr,
		Content:   content,
	}

	err = storage.Create(userActivity)
	if err != nil {
		return client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	}

	message := "mi papacho el más usain vol 🏃‍♂️"
	if content != "" {
		message = fmt.Sprintf("uy mi papacho corrió %s? lo iba robar un negro o qué manito. anwy congrats", content)
	}

	return client.SendMessage(chatID, message)
}
