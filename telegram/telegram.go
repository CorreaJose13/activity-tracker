package telegram

import (
	"activity-tracker/shared"
	"activity-tracker/telegram/commands"
	"errors"
	"strings"
)

var (
	errMissingUser = errors.New("user object is empty")
	errInvalidUser = errors.New("user is not allowed to use Samantha Bot")

	allowedUsers = map[string]bool{
		shared.Brayan: true,
		shared.Mauro:  true,
		shared.Johan:  true,
		shared.Jose:   true,
		shared.Valery: true,
	}
)

// Fetch is the main function to process telegram updates
func Fetch(client *shared.Client, update shared.Update) error {
	err := process(client, update)
	if err != nil {
		return client.SendMessage(update.Message.From.ID, "error while proccess: "+err.Error())
	}

	return nil
}

func process(client *shared.Client, update shared.Update) error {
	if update.Message != nil {
		return processMessage(client, update.Message)
	}

	return nil
}

func processMessage(client *shared.Client, message *shared.Message) error {
	user := message.From
	if user == nil {
		return errMissingUser
	}

	text := message.Text

	_, ok := allowedUsers[user.UserName]
	if !ok {
		return errInvalidUser
	}

	if !strings.HasPrefix(text, "/") {
		return nil
	}

	err := commands.DoCommand(client, message.Chat.ID, user.UserName, text)
	if err != nil {
		return client.SendMessage(message.Chat.ID, "can't do command: "+err.Error())
	}

	return nil
}
