package gemini

import (
	"activity-tracker/gemini"
	"activity-tracker/shared"
	"context"
	"errors"
	"fmt"
)

var (
	errNoPrompt = errors.New("en serio metiste un prompt vacío bobo hijueputa? haz algo bien sub-humano 🙄")
	forceText   = ", responde solo texto."
)

func HandleGemini(ctx context.Context, client *shared.Client, chatID int64, userName, content string) error {
	if content == "" {
		return errNoPrompt
	}

	//content concatenated with ForceText string forces Gemini to return a response using text format only
	response, err := gemini.QueryGemini(content + forceText)

	if err != nil {
		return client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	}

	return client.SendMessage(chatID, response)
}
