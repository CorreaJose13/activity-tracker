package trackers

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"context"
	"errors"
	"fmt"
)

var (
	missingRunContentMessage = "mandame el hpta numero de kilometros que corriste bobo hpta"
	invalidRunContentMessage = " eso ni siquiera es un número válido, perro hpta"

	templateSuccessMessageRun = "uy manito corriste %s  kilometros? te benias cagando o qué? 🏃🏾‍♂️ te felicito bb"

	ErrRunInvalidNumber = errors.New("invalid run number")

	mapRunMessagesBySource = ErrorMessages{
		ErrInvalidContent: {
			APISource: missingRunContentMessage,
			TGSource:  missingRunContentMessage,
		},
		ErrRunInvalidNumber: {
			APISource: invalidRunContentMessage,
			TGSource:  invalidRunContentMessage,
		},
	}
)

type RunTracker struct {
	activityType shared.Activity
	sourceType   SourceType
	content      string
}

func NewRunTracker(activityType shared.Activity, source SourceType) (Tracker, error) {
	return &RunTracker{
		activityType: activityType,
		sourceType:   source,
	}, nil
}

func (t *RunTracker) Track(ctx context.Context, username string, content string) error {
	if content == "" {
		return ErrInvalidContent
	}

	if !shared.IsValidFloat(content) {
		return ErrRunInvalidNumber
	}

	userActivity, err := shared.NewActivity(shared.Run, username, content)
	if err != nil {
		return err
	}

	t.content = content

	return storage.Create(ctx, userActivity)
}

func (t *RunTracker) GetErrorMessage(err error) string {
	return GetErrorMessageByTracker(err, t.sourceType, mapRunMessagesBySource)
}

func (t *RunTracker) GetSuccessMessage() string {
	return fmt.Sprintf(templateSuccessMessageRun, t.content)
}
