package trackers

import (
	"activity-tracker/shared"
	"context"
	"errors"
	"fmt"
)

var (
	ErrTrackerNotImplemented = errors.New("tracker not implemented")
	ErrInvalidContent        = errors.New("invalid content")

	mapTrackerByActivity = map[shared.Activity]func(shared.Activity, SourceType) (Tracker, error){
		shared.Cycling:    NewCyclingTracker,
		shared.Gomita:     NewGomitaTracker,
		shared.Gym:        NewGymTracker,
		shared.Keratine:   NewKeratineTracker,
		shared.Pipi:       NewPipiTracker,
		shared.Poop:       NewPoopTracker,
		shared.Read:       NewReadTracker,
		shared.Run:        NewRunTracker,
		shared.Shower:     NewShowerTracker,
		shared.Sleep:      NewSleepTracker,
		shared.Swimming:   NewSwimmingTracker,
		shared.ToothBrush: NewToothBrushTracker,
		shared.Water:      NewWaterTracker,
	}
)

type TrackerError struct {
	BaseError error
	Details   string
}

func (e *TrackerError) Error() string {
	return fmt.Sprintf("%s: %s", e.BaseError.Error(), e.Details)
}
func (e *TrackerError) Unwrap() error {
	return e.BaseError
}

type SourceType string
type ErrorMessages map[error]map[SourceType]string

const (
	APISource SourceType = "api"
	TGSource  SourceType = "telegram"
)

type Tracker interface {
	Track(ctx context.Context, username string, content string) error
	GetErrorMessage(err error) string
	GetSuccessMessage() string
}

func NewTracker(activityType shared.Activity, source SourceType) (Tracker, error) {
	newTracker, ok := mapTrackerByActivity[activityType]
	if !ok {
		return nil, &TrackerError{
			BaseError: ErrTrackerNotImplemented,
			Details:   fmt.Sprintf("[%s]", activityType),
		}
	}

	return newTracker(activityType, source)
}

func GetErrorMessageByTracker(err error, source SourceType, messages ErrorMessages) string {
	messageBySource, ok := messages[err]
	if !ok {
		return fmt.Sprintf(shared.ErrSendMessage, err.Error())
	}

	message, ok := messageBySource[source]
	if !ok {
		return fmt.Sprintf(shared.ErrSendMessage, err.Error())
	}

	return message
}
