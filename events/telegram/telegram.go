package telegram

import (
	"activity-tracker/api/telegram"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
)

var (
	errMissingUser = errors.New("user object it's empty")
	errInvalidUser = errors.New("user not allowed to use JH Bot")

	allowedUsers = map[string]bool{
		"BrayanEscobar": true,
		"mcortazar":     true,
		"JohanFlorez":   true,
		"jCorreaM":      true,
	}
)

func Fetch(bot *telegram.Bot, update telegram.Update) (err error) {
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
	text := message.Text
	date := time.Now()

	if user == nil {
		return errMissingUser
	}

	if _, ok := allowedUsers[user.UserName]; !ok {
		return errInvalidUser
	}

	if strings.HasPrefix(text, "/") {
		// Print to console username,text and date
		log.Printf("got new command '%s' from '%s at %s", text, user.UserName, date)
		if err := doCommand(bot, message.Chat.ID, user.UserName, text); err != nil {
			return fmt.Errorf("can't do command: %w", err)
		}
	}
	return nil
}
