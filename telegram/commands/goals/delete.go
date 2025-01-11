package goals

import "activity-tracker/shared"

// SendDeleteGoal handles the delete goal message
func SendDeleteGoal(client *shared.Client, userName, content string, chatID int64) error {
	return client.SendMessage(chatID, "goal deleted fake")
}
