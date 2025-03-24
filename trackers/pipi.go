package trackers

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"context"
	"time"
)

var (
	mapPipiMessagesBySource = ErrorMessages{}
)

type PipiTracker struct {
	activityType shared.Activity
	sourceType   SourceType
	username     string
}

func NewPipiTracker(activityType shared.Activity, source SourceType) (Tracker, error) {
	return &PipiTracker{
		activityType: activityType,
		sourceType:   source,
	}, nil
}

func (t *PipiTracker) Track(ctx context.Context, username string, content string) error {
	now, err := shared.GetNow()
	if err != nil {
		return err
	}

	nowStr := now.Format(time.RFC3339)

	userActivity := shared.UserActivity{
		ID:        shared.GenerateActivityItemID(now, username, shared.Pipi),
		Name:      username,
		Activity:  shared.Pipi,
		CreatedAt: nowStr,
	}

	t.username = username

	return storage.Create(ctx, userActivity)
}

func (t *PipiTracker) GetErrorMessage(err error) string {
	return GetErrorMessageByTracker(err, t.sourceType, mapPipiMessagesBySource)
}

func (t *PipiTracker) GetSuccessMessage() string {
	if t.username == shared.Valery {
		return "Epa, buena esa mionsita 😎"
	}

	return "Epa, buena esa mionsito 😎"
}
