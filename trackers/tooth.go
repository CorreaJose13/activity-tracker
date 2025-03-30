package trackers

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"context"
	"time"
)

var (
	successMessageTooth = "menos mal, ya te olia a qlo la boca mi papacho 💩"

	mapToothBrushMessagesBySource = ErrorMessages{}
)

type ToothBrushTracker struct {
	activityType shared.Activity
	sourceType   SourceType
}

func NewToothBrushTracker(activityType shared.Activity, source SourceType) (Tracker, error) {
	return &ToothBrushTracker{
		activityType: activityType,
		sourceType:   source,
	}, nil
}

func (t *ToothBrushTracker) Track(ctx context.Context, username string, content string) error {
	now, err := shared.GetNow()
	if err != nil {
		return err
	}

	nowStr := now.Format(time.RFC3339)

	userActivity := shared.UserActivity{
		ID:        shared.GenerateActivityItemID(now, username, shared.ToothBrush),
		Name:      username,
		Activity:  shared.ToothBrush,
		CreatedAt: nowStr,
		Content:   content,
	}

	return storage.Create(ctx, userActivity)
}

func (t *ToothBrushTracker) GetErrorMessage(err error) string {
	return GetErrorMessageByTracker(err, t.sourceType, mapToothBrushMessagesBySource)
}

func (t *ToothBrushTracker) GetSuccessMessage() string {
	return successMessageTooth
}
