package main

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/require"
)

func newApiGatewayRequest(body map[string]any) (events.APIGatewayProxyRequest, error) {
	responseBody, err := json.Marshal(body)
	if err != nil {
		return events.APIGatewayProxyRequest{}, err
	}

	return events.APIGatewayProxyRequest{
		Body: string(responseBody),
	}, nil
}

func TestHandlerInvalidTrackerType(t *testing.T) {
	c := require.New(t)

	request, err := newApiGatewayRequest(map[string]any{})

	c.NoError(err)

	response, err := handler(context.Background(), request)
	c.NoError(err)

	c.Equal(response.StatusCode, http.StatusBadRequest)
	c.Contains(response.Body, "estupid@ el type")
}

func TestHandlerMissingUsername(t *testing.T) {
	c := require.New(t)

	request, err := newApiGatewayRequest(map[string]any{
		"type": "dummy",
	})
	c.NoError(err)

	response, err := handler(context.Background(), request)
	c.NoError(err)

	c.Equal(response.StatusCode, http.StatusBadRequest)
	c.Contains(response.Body, "no hay username, a quien putas trackeo?")
}

func TestHandlerTrackerTypeNotImplemented(t *testing.T) {
	c := require.New(t)

	request, err := newApiGatewayRequest(map[string]any{
		"type":     "dummy",
		"username": "dummy",
		"content":  "dummy",
	})
	c.NoError(err)

	response, err := handler(context.Background(), request)
	c.NoError(err)

	c.Equal(response.StatusCode, http.StatusBadRequest)
	c.Contains(response.Body, "nunca en mi vida e visto ese type")
}

func TestHandlerTrackFailed(t *testing.T) {
	c := require.New(t)

	request, err := newApiGatewayRequest(map[string]any{
		"type":     "gym",
		"username": "dummy",
		"content":  "dummy",
	})
	c.NoError(err)

	response, err := handler(context.Background(), request)
	c.NoError(err)

	c.Equal(response.StatusCode, http.StatusBadRequest)
	c.Contains(response.Body, "eh pero vos sos tonto o te haces? mandame el time y el muscle")
}
