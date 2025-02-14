package swimming

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"context"
	"fmt"
	"time"
)

// SendTrackSwimming tracks swimming activity
func SendTrackSwimming(ctx context.Context, client *shared.Client, userName, content string, chatID int64) error {
	now, err := shared.GetNow()
	if err != nil {
		return client.SendMessage(chatID, err.Error())
	}

	nowStr := now.Format(time.RFC3339)

	userActivity := shared.UserActivity{
		ID:        shared.GenerateActivityItemID(now, userName, shared.Swimming),
		Name:      userName,
		Activity:  shared.Swimming,
		CreatedAt: nowStr,
		Content:   content,
	}

	err = storage.Create(ctx, userActivity)
	if err != nil {
		return client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	}

	message := "mi papacho el más bagrecito 🐟️"
	if content != "" {
		message = fmt.Sprintf("uy mi papacho nadó %s? lo iba robar un bagre negro o q? anwy congrats", content)
	}

	return client.SendMessage(chatID, message)
}
