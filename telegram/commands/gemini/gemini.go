package gemini

import (
	"activity-tracker/gemini"
	"activity-tracker/shared"
	"errors"
	"fmt"
)

var (
	ErrNoPrompt = errors.New("En serio metiste un prompt vacío bobo hijueputa? Haz algo bien sub-humano 🙄")
)

func HandleGemini(bot *shared.Bot, chatID int64, userName, content string) error {
	if content == "" {
		return ErrNoPrompt
	}
	response, err := gemini.Gemini(content + "responde solo texto")

	if err != nil {
		return shared.SendMessage(bot, chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	}

	return shared.SendMessage(bot, chatID, response)
}
