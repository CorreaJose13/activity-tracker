package tooth

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"context"
	"fmt"
	"time"
)

// SendTrackTooth tracks the tooth activity
func SendTrackTooth(ctx context.Context, client *shared.Client, userName, _ string, chatID int64) error {
	now, err := shared.GetNow()
	if err != nil {
		return client.SendMessage(chatID, err.Error())
	}

	nowStr := now.Format(time.RFC3339)

	userActivity := shared.UserActivity{
		ID:        shared.GenerateActivityItemID(now, userName, shared.ToothBrush),
		Name:      userName,
		Activity:  shared.ToothBrush,
		CreatedAt: nowStr,
	}

	err = storage.Create(ctx, userActivity)
	if err != nil {
		return client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	}

	return client.SendMessage(chatID, "menos mal, ya te olia a qlo la boca mi papacho 💩")
}
