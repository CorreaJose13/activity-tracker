package telegram

import (
	"activity-tracker/api/telegram"
	"activity-tracker/shared"
	"errors"
	"fmt"
	"log"
	"strings"
)

var (
	errMissingUser = errors.New("user object it's empty")
	errInvalidUser = errors.New("user not allowed to use JH Bot")

	allowedUsers = map[string]bool{
		shared.Brayan: true,
		shared.Mauro:  true,
		shared.Johan:  true,
		shared.Jose:   true,
		shared.Valery: true,
	}
)

func Fetch(bot *telegram.Bot, update telegram.Update) error {
	if err := Process(bot, update); err != nil {
		return fmt.Errorf("error while proccess: %w", err)
	}

	return nil
}

func Process(bot *telegram.Bot, update telegram.Update) error {
	if update.Message != nil {
		return processMessage(bot, update.Message)
	}

	return nil
}

func processMessage(bot *telegram.Bot, message *telegram.Message) error {
	user := message.From
	if user == nil {
		return errMissingUser
	}

	text := message.Text
	date := shared.GetNow()

	_, ok := allowedUsers[user.UserName]
	if !ok {
		return errInvalidUser
	}

	if strings.HasPrefix(text, "/") {
		// Print to console username,text and date
		log.Printf("got new command '%s' from '%s at %s", text, user.UserName, date)

		err := doCommand(bot, message.Chat.ID, user.UserName, text)
		if err != nil {
			return fmt.Errorf("can't do command: %w", err)
		}
	}

	return nil
}
