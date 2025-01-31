package track

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"fmt"
)

var (
	missingPoopContentMessage = "sos down???? mandame la info así:\n\t /track poop <times u pooped today>"
	invalidPoopContentMessage = "al parecer la kk la tienes en el cerebro, manda el número bien ps 😡"
	successPoopMessage        = "a ber de q color es? 👀👀"
)

// SendTrackPoop tracks how many times did u poop
func SendTrackPoop(client *shared.Client, userName, content string, chatID int64) error {
	if content == "" {
		return client.SendMessage(chatID, missingPoopContentMessage)
	}

	if !shared.IsValidNumber(content) {
		return client.SendMessage(chatID, invalidPoopContentMessage)
	}

	userActivity, err := shared.NewActivity(shared.Poop, userName, content)
	if err != nil {
		return client.SendMessage(chatID, err.Error())
	}

	err = storage.Create(userActivity)
	if err != nil {
		return client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	}

	return client.SendMessage(chatID, successPoopMessage)
}
