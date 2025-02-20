package shared

import (
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	// ErrForcedFailure error when force mock to fail
	ErrForcedFailure = errors.New("forced mock to failure")
	// ForceMockFailure to make mock fail
	ForceMockFailure = false
)

type tgClientInterface interface {
	Send(c tgbotapi.Chattable) (tgbotapi.Message, error)
	Request(c tgbotapi.Chattable) (*tgbotapi.APIResponse, error)
}

func ActivateMockBot(c *Client) {
	ForceMockFailure = false
	c.Bot = &mockBot{}
}

func DeactivateMockBot(c *Client) {
	ForceMockFailure = false
	c.Bot = BotAPI
}

type mockBot struct{}

func NewMockBot(token string) (*Client, error) {
	ForceMockFailure = false
	return &Client{Bot: &mockBot{}}, nil
}

func (m *mockBot) Send(c tgbotapi.Chattable) (tgbotapi.Message, error) {
	if ForceMockFailure {
		return tgbotapi.Message{}, ErrForcedFailure
	}

	return tgbotapi.Message{}, nil
}

func (m *mockBot) Request(c tgbotapi.Chattable) (*tgbotapi.APIResponse, error) {
	if ForceMockFailure {
		return nil, ErrForcedFailure
	}

	return &tgbotapi.APIResponse{Ok: true}, nil
}
