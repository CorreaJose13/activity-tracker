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
	botToken := Get("BOT_TOKEN", "")
	mongoToken := Get("MONGO_TOKEN", "")

	if botToken == "" {
		return &Config{
			TgBotToken:           "",
			MongoConnectionToken: mongoToken,
		}, errMissingBotToken
	}
	if mongoToken == "" {
		return &Config{
			TgBotToken:           botToken,
			MongoConnectionToken: "",
		}, errMissingMongoToken
	}

	return &Config{
		TgBotToken:           botToken,
		MongoConnectionToken: mongoToken,
	}, nil
}
