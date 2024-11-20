package main

import (
	"activity-tracker/shared"
	"context"
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/lambda"
)

var (
	bot      *shared.Bot
	botToken = os.Getenv("BOT_TOKEN")
	chatID   = os.Getenv("CHAT_ID")
)

func init() {
	var err error
	bot, err = shared.New(botToken)
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

	return shared.SendMessage(bot, i, "ya viene siendo como hora de tomar awita perr@ hpta 🙂")
}

func main() {
	lambda.Start(handler)
}
