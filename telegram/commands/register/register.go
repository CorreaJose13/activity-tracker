package register

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"context"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	msgUserAlreadyRegistered = "User already registered"
	acceptMessageButton      = "Accept ✅"
	rejectMessageButton      = "Reject ❌"
	msgUserWantsToRegister   = `User %s (ID: %d) wants to register 👀

%s

%s`
)

// RegisterUser registers a user
func RegisterUser(ctx context.Context, client *shared.Client, userName string, chatID int64) error {
	_, err := storage.GetUser(ctx, userName)
	if err == nil {
		return client.SendMessage(chatID, msgUserAlreadyRegistered)
	}

	for _, adminChatID := range shared.AdminUsersChatIDs {
		msg := tgbotapi.NewMessage(adminChatID, fmt.Sprintf(msgUserWantsToRegister, userName, chatID))

		keyboard := shared.NewInlineKeyboardMarkup(
			shared.NewInlineKeyboardRow(
				shared.NewInlineKeyboardButtonData(acceptMessageButton, fmt.Sprintf("accept_%d_%s", chatID, userName)),
				shared.NewInlineKeyboardButtonData(rejectMessageButton, fmt.Sprintf("reject_%d_%s", chatID, userName)),
			),
		)
		msg.ReplyMarkup = keyboard

		_, err := client.Bot.Send(msg)
		if err != nil {
			return client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
		}
	}

	return nil
}
