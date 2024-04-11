package telegram

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	timeOut = 60
	offset  = 0
)

func New(token string) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("getting bot object failed: %w", err)
	}
	// Set this to true to log all interactions with telegram servers
	bot.Debug = false
	return bot, nil
}

func Updates(bot *Bot) Channel {
	// Set up updates configuration.
	u := tgbotapi.NewUpdate(offset)
	u.Timeout = timeOut
	// `updates` is a golang channel which receives telegram updates
	updates := bot.GetUpdatesChan(u)
	return updates
}

func SendMessage(bot *Bot, chatID int64, text string) error {
	msg := tgbotapi.NewMessage(chatID, text)
	_, err := bot.Send(msg)
	if err != nil {
		return fmt.Errorf("can't send message: %w", err)
	}
	return nil
}

func SendPhoto(bot *Bot, chatID int64, url string) error {
	msg := tgbotapi.NewPhoto(chatID, tgbotapi.FileURL(url))
	_, err := bot.Send(msg)
	if err != nil {
		return fmt.Errorf("can't send message: %w", err)
	}
	return nil
}
