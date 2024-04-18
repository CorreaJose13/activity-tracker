package telegram

import (
	"activity-tracker/api/telegram"
	"context"
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

func Fetch(ctx context.Context, bot *telegram.Bot, updates telegram.Channel) (err error) {
	// `for {` means the loop is infinite until we manually stop it
	for {
		select {
		// stop looping if ctx is cancelled
		case <-ctx.Done():
			//   resolver este inconveniente
			return
		// receive update from channel and then handle it
		case update := <-updates:
			if err := Process(bot, update); err != nil {
				return fmt.Errorf("error while proccess: %w", err)
			}
		}
	}
}

func Process(bot *telegram.Bot, update telegram.Update) error {
	//update.UpdateID can be handy when using webhooks
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
