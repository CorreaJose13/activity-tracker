package shower

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"context"
	"fmt"
	"time"
)

// SendTrackShower tracks the shower activity
func SendTrackShower(ctx context.Context, client *shared.Client, userName, content string, chatID int64) error {
	now, err := shared.GetNow()
	if err != nil {
		return client.SendMessage(chatID, err.Error())
	}

	nowStr := now.Format(time.RFC3339)

	userActivity := shared.UserActivity{
		ID:        shared.GenerateActivityItemID(now, userName, shared.Shower),
		Name:      userName,
		Activity:  shared.Shower,
		CreatedAt: nowStr,
	}

	err = storage.Create(ctx, userActivity)
	if err != nil {
		return client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	}

	return client.SendMessage(chatID, "ya era hora cochino hijueputa 🤢🤢🤢")
}
