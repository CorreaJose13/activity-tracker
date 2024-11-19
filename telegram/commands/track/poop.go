package track

import "activity-tracker/shared"

// SendTrackPoop tracks the poop activity
func SendTrackPoop(bot *shared.Bot, userName, content string, chatID int64) error {
	return shared.SendMessage(bot, chatID, "a ber?")
}
