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

var errMissingUser = errors.New("user object it's empty")

func Fetch(ctx context.Context, bot *telegram.Bot, updates telegram.Channel) error {
	// `for {` means the loop is infinite until we manually stop it
	for {
		select {
		// stop looping if ctx is cancelled
		case <-ctx.Done():
			return nil
		// receive update from channel and then handle it
		case update := <-updates:
			if err := Process(bot, update); err != nil {
				log.Println(err)
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

func processMessage(bot *telegram.Bot, message *telegram.Message) (err error) {
	user := message.From
	text := message.Text
	date := time.Now()

	if user == nil {
		return errMissingUser
	}

	if strings.HasPrefix(text, "/") {
		// Print to console username,text and date
		log.Printf("got new command '%s' from '%s at %s", text, user.UserName, date)
		err = doCommand(bot, message.Chat.ID, text)
	}
	if err != nil {
		return fmt.Errorf("an error ocurred: %w", err)
	}
	return nil
}
