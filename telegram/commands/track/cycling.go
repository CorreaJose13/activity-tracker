package track

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"fmt"
	"time"
)

// SendTrackCycling tracks cycling activity
func SendTrackCycling(client *shared.Client, userName, content string, chatID int64) error {
	now, err := shared.GetNow()
	if err != nil {
		return client.SendMessage(chatID, err.Error())
	}

	nowStr := now.Format(time.RFC3339)

	userActivity := shared.UserActivity{
		ID:        shared.GenerateActivityItemID(now, userName, shared.Cycling),
		Name:      userName,
		Activity:  shared.Cycling,
		CreatedAt: nowStr,
		Content:   content,
	}

	err = storage.Create(userActivity)
	if err != nil {
		return client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	}

	message := "ataca pogachaaaaaa 🚴"
	if content != "" {
		message = "ve pero y ese poco de kilometros? te perseguía un veneco o q? anwy congrats"
	}

	return client.SendMessage(chatID, message)
}
