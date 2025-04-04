package main

import (
	"activity-tracker/shared"
	tg "activity-tracker/telegram"
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	client   *shared.Client
	tokenbot = os.Getenv("BOT_TOKEN")
)

func init() {
	var err error
	client, err = shared.New(tokenbot)
	if err != nil {
		panic(err)
	}
}

func processor(ctx context.Context, client *shared.Client, update shared.Update) error {
	err := tg.Fetch(ctx,client, update)
	if err != nil {
		return err
	}

	return nil
}

func HandleRequest(ctx context.Context, event interface{}) (events.APIGatewayProxyResponse, error) {
	b, err := json.Marshal(event)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "json.Marshal method threw error",
		}, err
	}

	var update tgbotapi.Update

	err = json.Unmarshal([]byte(b), &update)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "json.Unmarshal method threw error",
		}, err
	}

	err = processor(ctx, client, update)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "failed processing message",
		}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "success",
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
