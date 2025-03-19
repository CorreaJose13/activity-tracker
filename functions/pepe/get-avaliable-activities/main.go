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

func buildJsonResponse(activities []*shared.Activity, headers map[string]string) (events.APIGatewayProxyResponse, error) {
	if len(activities) == 0 {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusNotFound,
			Body:       "No activities found",
			Headers:    headers,
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
			Headers:    headers,
		}, nil // do not retry
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(responseBytes),
		Headers:    headers,
	}, nil
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	headers := map[string]string{
		"Content-Type":                "application/json",
		"Access-Control-Allow-Origin": "*",
		"Cache-Control":               "no-store",
		"Pragma":                      "no-cache",
		"Strict-Transport-Security":   "max-age=63072000",
	}

	activities, err := storage.GetAvailableActivities(ctx)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Internal server error",
			Headers:    headers,
		}, err
	}

	return buildJsonResponse(activities, headers)
}

func main() {
	lambda.Start(handler)
}
