package track

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"fmt"
	"strconv"
)

var (
	missingContentMessage = "y vos qué creés que voy a trackear si no me mandas el número de páginas que te leíste 🐸"
	invalidContentMessage = "es muy difícil mandar un número bien? conectá solo 2 neuronas y mandá un número de páginas 🐸"
	successMessage        = "congrats por decidir dejar la ignorancia un poquito atrás, seguí leyendo a ver si cambiamos este cochino país 👃🏾◻️"
)

// SendTrackRead tracks the read activity
func SendTrackRead(client *shared.Client, userName, content string, chatID int64) error {
	if content == "" {
		return client.SendMessage(chatID, missingContentMessage)
	}

	if !isValidContent(content) {
		return client.SendMessage(chatID, invalidContentMessage)
	}

	userActivity, err := shared.NewActivity(shared.Read, userName, content)
	if err != nil {
		return client.SendMessage(chatID, err.Error())
	}

	err = storage.Create(userActivity)
	if err != nil {
		return client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	}

	return client.SendMessage(chatID, successMessage)
}

func isValidContent(content string) bool {
	value, err := strconv.Atoi(content)

	return err == nil && value > 0
}
