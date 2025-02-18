package main

import (
	"activity-tracker/shared"
	"activity-tracker/telegram/commands/report"
	"context"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

var (
	client   *shared.Client
	botToken = os.Getenv("BOT_TOKEN")
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

// TODO: This lambda will be used in another ticket when the TF files are created
func handler(ctx context.Context, event Schedule) error {
	for userName, chatID := range shared.AllReportsSchedulerChatIDs {
		err := report.GenerateAllReports(client, userName, "", chatID)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	lambda.Start(handler)
}
