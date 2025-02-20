package telegram

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"activity-tracker/telegram/commands"
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	register "activity-tracker/telegram/commands/register"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	acceptAction       = "accept"
	rejectAction       = "reject"
	registerCommand    = "/register"
	approvedMessage    = "✅ Your registration has been approved! You now can use the bot"
	rejectedMessage    = "❌ Sorry, your registration request has been rejected."
	userApprovedStatus = "User has been approved"
	userRejectedStatus = "User has been rejected"
	invalidChatID      = "invalid chat ID in callback: %w"
	failedCreateUser   = "failed to create user: %w"
	failedNotify       = "failed to notify user: %w"
	unknownAction      = "unknown action: %s"
	failedCallback     = "failed to answer callback: %w"
	failedUpdateMsg    = "failed to update message: %w"
	denyRegisterMsg    = "Registering in private chat is not allowed 🤡🤡"

	callbackDataParts = 3
)

var (
	errMissingUser        = errors.New("user object is empty")
	errInvalidUser        = errors.New("user is not allowed to use Samantha Bot")
	errInvalidCallbackFmt = errors.New("invalid callback data format")
)

// Fetch is the main function to process telegram updates
func Fetch(ctx context.Context, client *shared.Client, update shared.Update) error {
	err := process(ctx, client, update)
	if err != nil {
		return client.SendMessage(update.Message.From.ID, "error while proccess: "+err.Error())
	}

	return nil
}

func process(ctx context.Context, client *shared.Client, update shared.Update) error {
	if update.Message != nil {
		return processMessage(ctx, client, update.Message)
	}

	if update.CallbackQuery != nil {
		return processCallbackQuery(ctx, client, update.CallbackQuery)
	}

	return nil
}

func processMessage(ctx context.Context, client *shared.Client, message *shared.Message) error {
	user := message.From
	if user == nil {
		return errMissingUser
	}

	text := message.Text

	if !strings.HasPrefix(text, "/") {
		return nil
	}

	if text == registerCommand {
		if message.Chat.ID != shared.GroupChatID {
			return client.SendMessage(message.Chat.ID, denyRegisterMsg)
		}

		return register.RegisterUser(ctx, client, user.UserName, message.Chat.ID)
	}

	_, err := storage.GetUser(ctx, user.UserName)
	if err != nil {
		return errInvalidUser
	}

	err = commands.DoCommand(ctx, client, message.Chat.ID, user.UserName, text)
	if err != nil {
		return client.SendMessage(message.Chat.ID, "can't do command: "+err.Error())
	}

	return nil
}

// processCallbackQuery is the function to process callback queries, modify this function to handle other callback queries
func processCallbackQuery(ctx context.Context, client *shared.Client, callback *shared.CallbackQuery) error {
	parts := strings.Split(callback.Data, "_")
	if len(parts) != callbackDataParts {
		return errInvalidCallbackFmt
	}

	action := parts[0]

	userChatID, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return fmt.Errorf(invalidChatID, err)
	}

	userName := parts[2]
	var responseMsg string

	switch action {
	case acceptAction:
		user := shared.NewUser(userName, userChatID, []shared.Activity{})

		err = storage.CreateUser(ctx, *user)
		if err != nil {
			return fmt.Errorf(failedCreateUser, err)
		}

		err = client.SendMessage(userChatID, approvedMessage)
		if err != nil {
			return fmt.Errorf(failedNotify, err)
		}

		responseMsg = userApprovedStatus

	case rejectAction:
		err = client.SendMessage(userChatID, rejectedMessage)
		if err != nil {
			return fmt.Errorf(failedNotify, err)
		}

		responseMsg = userRejectedStatus

	default:
		return fmt.Errorf(unknownAction, action)
	}

	_, err = client.Bot.Request(tgbotapi.NewCallback(callback.ID, responseMsg))
	if err != nil {
		return fmt.Errorf(failedCallback, err)
	}

	editMsg := tgbotapi.NewEditMessageText(
		callback.Message.Chat.ID,
		callback.Message.MessageID,
		fmt.Sprintf("%s\n\nStatus: %s", callback.Message.Text, responseMsg),
	)

	_, err = client.Bot.Send(editMsg)
	if err != nil {
		return fmt.Errorf(failedUpdateMsg, err)
	}

	return nil
}
