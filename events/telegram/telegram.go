package telegram

import (
	"activity-tracker/api/telegram"
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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
			return Process(bot, update)
		}
	}
}

func Process(bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
	//update.UpdateID can be handy when using webhooks
	if update.Message != nil {
		return processMessage(bot, update.Message)
	}

	return nil
}

func processMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) (err error) {
	user := message.From
	text := message.Text
	date := time.Now()

	if user == nil {
		return errMissingUser
	}

	// Print to console username,text and date
	log.Printf("%s wrote %s at %s", user.UserName, text, date)

	if strings.HasPrefix(text, "/") {
		err = doCommand(bot, message.Chat.ID, text)
	}
	if err != nil {
		return fmt.Errorf("an error ocurred: %w", err)
	}
	return nil
}
