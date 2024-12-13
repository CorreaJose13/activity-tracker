package gemini

import (
	"activity-tracker/gemini"
	"activity-tracker/shared"
	"errors"
	"fmt"
)

var (
	errNoPrompt = errors.New("En serio metiste un prompt vacío bobo hijueputa? Haz algo bien sub-humano 🙄")
	forceText   = ",responde solo texto."
)

func HandleGemini(bot *shared.Bot, chatID int64, userName, content string) error {
	if content == "" {
		return errNoPrompt
	}

	//content concatenated with ForceText string forces Gemini to return a response using text format only
	response, err := gemini.QueryGemini(content + forceText)

	if err != nil {
		return shared.SendMessage(bot, chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	}

	return shared.SendMessage(bot, chatID, response)
}
