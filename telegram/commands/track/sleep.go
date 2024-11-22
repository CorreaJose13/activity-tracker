package track

import "activity-tracker/shared"

// SendTrackSleep tracks the sleep activity
func SendTrackSleep(bot *shared.Bot, userName, content string, chatID int64) error {
	return shared.SendMessage(bot, chatID, "zzzzz")
}
