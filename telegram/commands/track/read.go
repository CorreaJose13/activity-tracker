package track

import "activity-tracker/shared"

// SendTrackRead tracks the read activity
func SendTrackRead(client *shared.Client, userName, content string, chatID int64) error {
	return client.SendMessage(chatID, "lectura")
}
