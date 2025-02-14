package run

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"fmt"
)

var (
	missingContentMessage = "mandame el hpta numero de kilometros que corriste bobo hpta"
	invalidContentMessage = " eso ni siquiera es un número válido, perro hpta"
	successMessage        = ` uy manito corriste %s  kilometros? te benias cagando o qué? 🏃🏾‍♂️ te felicito bb`
)

// SendTrackRun tracks the run activity
func SendTrackRun(client *shared.Client, userName, content string, chatID int64) error {
	if content == "" {
		return client.SendMessage(chatID, missingContentMessage)
	}

	if !shared.IsValidFloat(content) {
		return client.SendMessage(chatID, invalidContentMessage)
	}

	userActivity, err := shared.NewActivity(shared.Run, userName, content)
	if err != nil {
		return client.SendMessage(chatID, err.Error())
	}

	err = storage.Create(userActivity)
	if err != nil {
		return client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	}

	formatedMessage := fmt.Sprintf(successMessage, content)

	return client.SendMessage(chatID, formatedMessage)
}
