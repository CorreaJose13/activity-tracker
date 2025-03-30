package trackers

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"context"
	"fmt"
	"time"
)

var (
	templateSuccessMessageSwimmingWithContent = "uy mi papacho nadó %s? lo iba robar un bagre negro o q? anwy congrats"
	successMessageSwimming                    = "mi papacho el más bagrecito 🐟️"

	mapSwimmingMessagesBySource = ErrorMessages{}
)

type SwimmingTracker struct {
	activityType shared.Activity
	sourceType   SourceType
	content      string
}

func NewSwimmingTracker(activityType shared.Activity, source SourceType) (Tracker, error) {
	return &SwimmingTracker{
		activityType: activityType,
		sourceType:   source,
	}, nil
}

func (t *SwimmingTracker) Track(ctx context.Context, username string, content string) error {
	now, err := shared.GetNow()
	if err != nil {
		return err
	}

	nowStr := now.Format(time.RFC3339)

	userActivity := shared.UserActivity{
		ID:        shared.GenerateActivityItemID(now, username, shared.Swimming),
		Name:      username,
		Activity:  shared.Swimming,
		CreatedAt: nowStr,
		Content:   content,
	}

	t.content = content

	return storage.Create(ctx, userActivity)
}

func (t *SwimmingTracker) GetErrorMessage(err error) string {
	return GetErrorMessageByTracker(err, t.sourceType, mapSwimmingMessagesBySource)
}

func (t *SwimmingTracker) GetSuccessMessage() string {
	if t.content != "" {
		return fmt.Sprintf(templateSuccessMessageSwimmingWithContent, t.content)
	}

	return successMessageSwimming
}
