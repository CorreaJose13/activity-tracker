package track

import "activity-tracker/shared"

// SendTrackRead tracks the read activity
func SendTrackRead(bot *shared.Bot, userName, content string, chatID int64) error {
	return shared.SendMessage(bot, chatID, "lectura")
}
