package track

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"fmt"
	"time"
)

var goalWaterConsume = 3

// SendTrackWater tracks the water activity
func SendTrackWater(bot *shared.Bot, userName, content string, chatID int64) error {
	isGoalCompleted := isWaterGoalCompleted(bot, userName, chatID)
	if isGoalCompleted {
		return shared.SendMessage(bot, chatID, "ya te tomaste los 3L de awa mi papacho, aprende a tener límites")
	}

	now, err := shared.GetNow()
	if err != nil {
		return shared.SendMessage(bot, chatID, err.Error())
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
		return shared.SendMessage(bot, chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	}

	return shared.SendMessage(bot, chatID, "se wardó tu tomadita de awa golosito")
}

func isWaterGoalCompleted(bot *shared.Bot, userName string, chatID int64) bool {
	currentDayWaterActivities, err := storage.GetCurrentDayActivities(userName, shared.Water)
	if err != nil {
		_ = shared.SendMessage(bot, chatID, "tenemos problemas papi"+err.Error())

		return false
	}

	return len(currentDayWaterActivities) >= goalWaterConsume
}
