package track

import "activity-tracker/shared"

// SendTrackPoop tracks the poop activity
func SendTrackPoop(client *shared.Client, userName, content string, chatID int64) error {
	return client.SendMessage(chatID, "a ber?")
}
