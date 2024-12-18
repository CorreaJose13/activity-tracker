package goals

import "activity-tracker/shared"

// SendCreateGoal handles the create goal message
func SendCreateGoal(client *shared.Client, userName, content string, chatID int64) error {
	return client.SendMessage(chatID, "goal created fake")
}
