package track

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"strings"
	"time"
)

// SendTrackGym tracks gym activity
func SendTrackGym(bot *shared.Bot, userName, content string, chatID int64) error {
	duration, muscle, ok := strings.Cut(content, " ")
	if !ok {
		return shared.SendMessage(bot, chatID, "eh pero vos sos tonto o te haces? mandame el time y el muscle. asi:\n\t /track gym <duration> <muscles by comma>")
	}

	if duration == "" {
		return shared.SendMessage(bot, chatID, "eh pero cuánto tiempo te ejercitaste sapa inmunda")
	}

	if muscle == "" {
		return shared.SendMessage(bot, chatID, "eh pero hiciste chisme al fallo o q 🐸? mandame el musculo que ejercitaste sapa. Ej: bicep,pecho,jeta")
	}

	now, err := shared.GetNow()
	if err != nil {
		return shared.SendMessage(bot, chatID, err.Error())
	}

	nowStr := now.Format(time.RFC3339)

	userActivity := shared.UserActivity{
		ID:        shared.GenerateActivityItemID(now, userName, shared.Gym),
		Name:      userName,
		Activity:  shared.Gym,
		CreatedAt: nowStr,
		Content:   content,
	}

	err = storage.Create(userActivity)
	if err != nil {
		return shared.SendMessage(bot, chatID, "algo falló mi fafá: "+err.Error())
	}

	return shared.SendMessage(bot, chatID, "isss mi papacho el próximo cbum ve")
}
