package goals

import "activity-tracker/shared"

// SendCreateGoal handles the create goal message
func SendCreateGoal(bot *shared.Bot, userName, content string, chatID int64) error {
	return shared.SendMessage(bot, chatID, "goal created fake")
}
