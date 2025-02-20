package pipi

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"context"
	"fmt"
	"time"
)

// SendTrackPipi tracks the pipi activity
func SendTrackPipi(ctx context.Context, client *shared.Client, userName, content string, chatID int64) error {
	now, err := shared.GetNow()
	if err != nil {
		return client.SendMessage(chatID, err.Error())
	}

	nowStr := now.Format(time.RFC3339)

	userActivity := shared.UserActivity{
		ID:        shared.GenerateActivityItemID(now, userName, shared.Pipi),
		Name:      userName,
		Activity:  shared.Pipi,
		CreatedAt: nowStr,
	}

	err = storage.Create(ctx, userActivity)
	if err != nil {
		return client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	}

	if userName == shared.Valery {
		return client.SendMessage(chatID, "Epa, buena esa mionsita 😎")
	}

	return client.SendMessage(chatID, "Epa, buena esa mionsito 😎")
}
