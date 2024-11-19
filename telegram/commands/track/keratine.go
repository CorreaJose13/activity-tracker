package track

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"fmt"
	"time"
)

var goalKeratineConsume = 1

// SendTrackKeratine tracks the keratine activity
func SendTrackKeratine(bot *shared.Bot, userName, content string, chatID int64) error {
	isGoalCompleted := isKeratineGoalCompleted(bot, userName, chatID)
	if isGoalCompleted {
		return shared.SendMessage(bot, chatID, "ya te tomaste la keratina de hoy, aprende a tener límites xfi")
	}

	now, err := shared.GetNow()
	if err != nil {
		return shared.SendMessage(bot, chatID, err.Error())
	}

	nowStr := now.Format(time.RFC3339)

	userActivity := storage.UserActivity{
		ID:        storage.GenerateActivityItemID(now, userName, shared.Keratine),
		Name:      userName,
		Activity:  shared.Keratine,
		CreatedAt: nowStr,
		Content:   content,
	}

	err = storage.Create(userActivity)
	if err != nil {
		return shared.SendMessage(bot, chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	}

	return shared.SendMessage(bot, chatID, "se wardó tu tomadita de keratina >:)")
}

func isKeratineGoalCompleted(bot *shared.Bot, userName string, chatID int64) bool {
	currentDayKeratineActivities, err := storage.GetCurrentDayActivities(userName, shared.Keratine)
	if err != nil {
		_ = shared.SendMessage(bot, chatID, "tenemos problemas papi"+err.Error())

		return false
	}

	return len(currentDayKeratineActivities) == goalKeratineConsume
}
