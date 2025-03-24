package trackers

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"context"
	"time"
)

var (
	mapShowerMessagesBySource = ErrorMessages{}
)

type ShowerTracker struct {
	activityType shared.Activity
	sourceType   SourceType
}

func NewShowerTracker(activityType shared.Activity, source SourceType) (Tracker, error) {
	return &ShowerTracker{
		activityType: activityType,
		sourceType:   source,
	}, nil
}

func (t *ShowerTracker) Track(ctx context.Context, username string, content string) error {
	now, err := shared.GetNow()
	if err != nil {
		return err
	}

	nowStr := now.Format(time.RFC3339)

	userActivity := shared.UserActivity{
		ID:        shared.GenerateActivityItemID(now, username, shared.Shower),
		Name:      username,
		Activity:  shared.Shower,
		CreatedAt: nowStr,
	}

	return storage.Create(ctx, userActivity)
}

func (t *ShowerTracker) GetErrorMessage(err error) string {
	return GetErrorMessageByTracker(err, t.sourceType, mapShowerMessagesBySource)
}

func (t *ShowerTracker) GetSuccessMessage() string {
	return "ya era hora cochino hijueputa 🤢🤢🤢"
}
