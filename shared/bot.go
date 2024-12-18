package shared

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot = tgbotapi.BotAPI

type Update = tgbotapi.Update

type Message = tgbotapi.Message

type Client struct {
	Bot tgClientInterface
}

var (
	BotAPI tgClientInterface = &tgbotapi.BotAPI{}
)

// New creates a new telegram bot
func New(token string) (*Client, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("getting bot object failed: %w", err)
	}

	// Set this to true to log all interactions with telegram servers
	bot.Debug = false

	BotAPI = bot

	return &Client{Bot: BotAPI}, nil
}

// SendMessage sends a message to the chat
func (c *Client) SendMessage(chatID int64, text string) error {
	msg := tgbotapi.NewMessage(chatID, text)

	_, err := c.Bot.Send(msg)
	if err != nil {
		return fmt.Errorf("can't send message: %w", err)
	}

	return nil
}

// SendPhoto sends a photo to the chat
func (c *Client) SendPhoto(chatID int64, url string) error {
	msg := tgbotapi.NewPhoto(chatID, tgbotapi.FileURL(url))

	_, err := c.Bot.Send(msg)
	if err != nil {
		return fmt.Errorf("can't send photo: %w", err)
	}

	return nil
}

// SendFile sends a file to the chat
func (c *Client) SendFile(chatID int64, filePath string) error {
	msg := tgbotapi.NewDocument(chatID, tgbotapi.FilePath(filePath))

	_, err := c.Bot.Send(msg)
	if err != nil {
		return fmt.Errorf("can't send file: %w", err)
	}

	return nil
}
