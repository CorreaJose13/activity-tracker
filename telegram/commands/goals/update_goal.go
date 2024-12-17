package goals

import "activity-tracker/shared"

// SendUpdateGoal handles the update goal message
func SendUpdateGoal(client *shared.Client, userName, content string, chatID int64) error {
	return client.SendMessage(chatID, "goal updated fake")
}
