package track

import "activity-tracker/shared"

// SendTrackShower tracks the shower activity
func SendTrackShower(bot *shared.Bot, userName, content string, chatID int64) error {
	return shared.SendMessage(bot, chatID, "ya era hora olias a obo")
}
