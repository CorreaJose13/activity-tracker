package keratine

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"context"
	"fmt"
	"time"
)

var goalKeratineConsume = 1

// SendTrackKeratine tracks the keratine activity
func SendTrackKeratine(ctx context.Context, client *shared.Client, userName, content string, chatID int64) error {
	isGoalCompleted := isKeratineGoalCompleted(ctx, client, userName, chatID)
	if isGoalCompleted {
		return client.SendMessage(chatID, "ya te tomaste la keratina de hoy, aprende a tener límites xfi")
	}

	now, err := shared.GetNow()
	if err != nil {
		return client.SendMessage(chatID, err.Error())
	}

	nowStr := now.Format(time.RFC3339)

	userActivity := shared.UserActivity{
		ID:        shared.GenerateActivityItemID(now, userName, shared.Keratine),
		Name:      userName,
		Activity:  shared.Keratine,
		CreatedAt: nowStr,
		Content:   content,
	}

	err = storage.Create(ctx, userActivity)
	if err != nil {
		return client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	}

	return client.SendMessage(chatID, "se wardó tu tomadita de keratina >:)")
}

func isKeratineGoalCompleted(ctx context.Context, client *shared.Client, userName string, chatID int64) bool {
	currentDayKeratineActivities, err := storage.GetCurrentDayActivities(ctx, userName, shared.Keratine)
	if err != nil {
		_ = client.SendMessage(chatID, "tenemos problemas papi"+err.Error())

		return false
	}

	return len(currentDayKeratineActivities) == goalKeratineConsume
}
