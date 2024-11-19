package goals

import "activity-tracker/shared"

// SendUpdateGoal handles the update goal message
func SendUpdateGoal(bot *shared.Bot, userName, content string, chatID int64) error {
	return shared.SendMessage(bot, chatID, "goal updated fake")
}
