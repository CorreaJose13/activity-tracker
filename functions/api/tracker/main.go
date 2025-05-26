package main

import (
	"activity-tracker/shared"
	"activity-tracker/trackers"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type TrackerRequest struct {
	ActivityType shared.Activity `json:"type"`
	Username     string          `json:"username"`
	Content      string          `json:"content"`
	tracker      trackers.Tracker
}

var (
	errMissingActivityType = errors.New("estupid@ el type")
	errMissingUsername     = errors.New("no hay username, a quien putas trackeo?")

	invalidActivityTypeMessage = "nunca en mi vida e visto ese type [%s]"
)

func newApiGatewayResponse(status int, body map[string]any) (events.APIGatewayProxyResponse, error) {
	body["status"] = status

	responseBody, err := json.Marshal(body)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       string(responseBody),
	}, nil
}

func (req *TrackerRequest) handleErrorMessage(err error) string {
	if errors.Is(err, trackers.ErrTrackerNotImplemented) {
		return fmt.Sprintf(invalidActivityTypeMessage, req.ActivityType)
	}

	return req.tracker.GetErrorMessage(err)
}

func parseRequest(request events.APIGatewayProxyRequest) (*TrackerRequest, error) {
	var trackerRequest TrackerRequest

	err := json.Unmarshal([]byte(request.Body), &trackerRequest)
	if err != nil {
		return nil, err
	}

	if trackerRequest.ActivityType == "" {
		return nil, errMissingActivityType
	}

	if trackerRequest.Username == "" {
		return nil, errMissingUsername
	}

	return &trackerRequest, nil
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	req, err := parseRequest(request)
	if err != nil {
		return newApiGatewayResponse(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
		})
	}

	req.tracker, err = trackers.NewTracker(req.ActivityType, trackers.APISource)
	if err != nil {
		return newApiGatewayResponse(http.StatusBadRequest, map[string]any{
			"message": req.handleErrorMessage(err),
		})
	}

	err = req.tracker.Track(ctx, req.Username, req.Content)
	if err != nil {
		return newApiGatewayResponse(http.StatusBadRequest, map[string]any{
			"message": req.handleErrorMessage(err),
		})
	}

	return newApiGatewayResponse(http.StatusCreated, map[string]any{
		"message": req.tracker.GetSuccessMessage(),
	})
}

func main() {
	lambda.Start(handler)
}
