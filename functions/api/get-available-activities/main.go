package main

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type response struct {
	Activities []*shared.Activity `json:"activities"`
}

func buildJsonResponse(activities []*shared.Activity) (events.APIGatewayProxyResponse, error) {
	if len(activities) == 0 {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusNotFound,
			Body:       "No activities found",
		}, nil
	}

	response := response{
		Activities: activities,
	}

	responseBytes, err := json.Marshal(response)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Internal server error",
		}, nil // do not retry
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(responseBytes),
	}, nil
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	activities, err := storage.GetAvailableActivities(ctx)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Internal server error",
		}, err
	}

	return buildJsonResponse(activities)
}

func main() {
	lambda.Start(handler)
}
