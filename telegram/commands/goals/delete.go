package goals

import "activity-tracker/shared"

// SendDeleteGoal handles the delete goal message
func SendDeleteGoal(bot *shared.Bot, userName, content string, chatID int64) error {
	return shared.SendMessage(bot, chatID, "goal deleted fake")
}
