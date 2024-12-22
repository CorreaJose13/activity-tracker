package track

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"fmt"
)

// SendTrackSleep tracks the sleep activity
func SendTrackSleep(client *shared.Client, userName, content string, chatID int64) error {
	userActivity, err := shared.NewActivity(shared.Sleep, userName)
	if err != nil {
		return client.SendMessage(chatID, err.Error())
	}

	err = storage.Create(userActivity)
	if err != nil {
		return client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	}

	return client.SendMessage(chatID, "Que descanses y sueñes conmigo 😏")
}
