package main

import (
	"activity-tracker/api/telegram"
	eventConsumer "activity-tracker/consumer/event_consumer"
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	bot      *telegram.Bot
	tokenbot = os.Getenv("BOT_TOKEN")
)

func init() {
	var err error
	bot, err = telegram.New(tokenbot)
	if err != nil {
		panic(err)
	}
}

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest(ctx context.Context, event interface{}) (events.APIGatewayProxyResponse, error) {
	b, err := json.Marshal(event)
	if err != nil {
		response := events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "json.Marshal method threw error",
		}
		return response, err
	}

	var update tgbotapi.Update
	if err := json.Unmarshal([]byte(b), &update); err != nil {
		response := events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "json.Unmarshal method threw error",
		}
		return response, err
	}

	if err := eventConsumer.Processor(bot, update); err != nil {
		response := events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "failed processing message",
		}
		return response, err
	}

	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "success",
	}
	return response, nil
}
