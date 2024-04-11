package telegram

import (
	"fmt"
	"log"

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
	log.Println("Message sent succesfully")
	return nil
}

func SendPhoto(bot *Bot, chatID int64) error {

	msg := tgbotapi.NewPhoto(chatID, tgbotapi.FileURL("https://external-preview.redd.it/jrtz49x5F1cjvDQoFzb0I4cv2dwhA5RDhqaEcBbiXIU.png?format=pjpg&auto=webp&s=3ef741c83f7927eca91cb8ac2d610fd6f010d5b0"))
	_, err := bot.Send(msg)
	if err != nil {
		return fmt.Errorf("can't send message: %w", err)
	}
	log.Println("Message sent succesfully")
	return nil
}
