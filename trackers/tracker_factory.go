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
	switch activityType {
	case shared.Cycling:
		return NewCyclingTracker(activityType, source)
	case shared.Gomita:
		return NewGomitaTracker(activityType, source)
	case shared.Gym:
		return NewGymTracker(activityType, source)
	case shared.Keratine:
		return NewKeratineTracker(activityType, source)
	case shared.Pipi:
		return NewPipiTracker(activityType, source)
	case shared.Poop:
		return NewPoopTracker(activityType, source)
	case shared.Read:
		return NewReadTracker(activityType, source)
	case shared.Run:
		return NewRunTracker(activityType, source)
	case shared.Shower:
		return NewShowerTracker(activityType, source)
	case shared.Sleep:
		return NewSleepTracker(activityType, source)
	case shared.Swimming:
		return NewSwimmingTracker(activityType, source)
	case shared.ToothBrush:
		return NewToothBrushTracker(activityType, source)
	case shared.Water:
		return NewWaterTracker(activityType, source)
	}

	return nil, &TrackerError{
		BaseError: ErrTrackerNotImplemented,
		Details:   fmt.Sprintf("[%s]", activityType),
	}
}

func GetErrorMessageByTracker(err error, source SourceType, messages ErrorMessages) string {
	messageBySource, ok := mapGymMessagesBySource[err]
	if !ok {
		return fmt.Sprintf(shared.ErrSendMessage, err.Error())
	}

	message, ok := messageBySource[source]
	if !ok {
		return fmt.Sprintf(shared.ErrSendMessage, err.Error())
	}

	return message
}
