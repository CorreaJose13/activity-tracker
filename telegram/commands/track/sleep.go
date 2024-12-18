package track

import "activity-tracker/shared"

// SendTrackSleep tracks the sleep activity
func SendTrackSleep(client *shared.Client, userName, content string, chatID int64) error {
	return client.SendMessage(chatID, "zzzzz")
}
