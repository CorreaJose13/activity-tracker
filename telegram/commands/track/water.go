package track

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"fmt"
	"time"
)

var goalWaterConsume = 3

// SendTrackWater tracks the water activity
func SendTrackWater(client *shared.Client, userName, content string, chatID int64) error {
	isGoalCompleted := isWaterGoalCompleted(client, userName, chatID)
	if isGoalCompleted {
		return client.SendMessage(chatID, "ya te tomaste los 3L de awa mi papacho, aprende a tener límites")
	}

	now, err := shared.GetNow()
	if err != nil {
		return client.SendMessage(chatID, err.Error())
	}

	nowStr := now.Format(time.RFC3339)

	userActivity := shared.UserActivity{
		ID:        shared.GenerateActivityItemID(now, userName, shared.Water),
		Name:      userName,
		Activity:  shared.Water,
		CreatedAt: nowStr,
		Content:   content, // TODO: add logic to validate the content and use it in isGoalCompleted function
	}

	err = storage.Create(userActivity)
	if err != nil {
		return client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	}

	return client.SendMessage(chatID, "se wardó tu tomadita de awa golosito")
}

func isWaterGoalCompleted(client *shared.Client, userName string, chatID int64) bool {
	currentDayWaterActivities, err := storage.GetCurrentDayActivities(userName, shared.Water)
	if err != nil {
		_ = client.SendMessage(chatID, "tenemos problemas papi"+err.Error())

		return false
	}

	return len(currentDayWaterActivities) >= goalWaterConsume
}
