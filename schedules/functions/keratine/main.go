package main

import (
	"activity-tracker/shared"
	"context"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

var (
	client   *shared.Client
	botToken = os.Getenv("BOT_TOKEN")
	message  = "acordate de tomar la creatina 💪 sapa asquerosa"
)

func init() {
	var err error
	client, err = shared.New(botToken)
	if err != nil {
		panic(err)
	}
}

// Schedule represents a programmed schedule
type Schedule struct {
	Message string `json:"message"`
}

func handler(ctx context.Context, event Schedule) error {
	for _, chatID := range shared.KeratineChatIDs {
		err := client.SendMessage(chatID, message)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	lambda.Start(handler)
}
