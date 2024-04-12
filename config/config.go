package config

import (
	"errors"
)

type Config struct {
	TgBotToken           string
	MongoConnectionToken string
}

var (
	errMissingBotToken   = errors.New("missing bot token")
	errMissingMongoToken = errors.New("missing mongo token")
)

func MustLoad() (*Config, error) {
	botToken := getValue("BOT_TOKEN", "")
	mongoToken := getValue("MONGO_TOKEN", "")

	if botToken == "" {
		return nil, errMissingBotToken
	}
	if mongoToken == "" {
		return nil, errMissingMongoToken
	}

	return &Config{
		TgBotToken:           botToken,
		MongoConnectionToken: mongoToken,
	}, nil
}
