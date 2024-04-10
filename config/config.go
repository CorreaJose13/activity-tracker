package config

import (
	"log"
)

type Config struct {
	TgBotToken           string
	MongoConnectionToken string
}

func MustLoad() Config {
	botToken := Get("BOT_TOKEN", "")
	mongoToken := Get("MONGO_TOKEN", "")

	if botToken == "" {
		log.Fatal("tg bot token is not specified")
	}
	if mongoToken == "" {
		log.Fatal("mongo connection string is not specified")
	}

	return Config{
		TgBotToken:           botToken,
		MongoConnectionToken: mongoToken,
	}
}
