package main

import (
	"activity-tracker/api/telegram"
	"activity-tracker/config"
	"context"
	"log"
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/lambda"
)

var (
	bot    *telegram.Bot
	chatID = os.Getenv("CHAT_ID")
)

func init() {
	cfg, err := config.MustLoad()
	if err != nil {
		panic(err)
	}

	bot, err = telegram.New(cfg.TgBotToken)
	if err != nil {
		panic(err)
	}
}

// Schedule represents a programmed schedule
type Schedule struct {
	Message string `json:"message"`
}

func handler(ctx context.Context, event Schedule) error {
	i, err := strconv.ParseInt(chatID, 10, 64)
	if err != nil {
		panic(err)
	}

	err = telegram.SendMessage(bot, i, "Tetas o culos mi fafá? no sé pero toma awita perro hpta")
	if err != nil {
		log.Println(err.Error())
	}

	return nil
}

func main() {
	lambda.Start(handler)
}
