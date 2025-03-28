package trackers

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"context"
	"errors"
	"time"
)

var (
	goalKeratineConsume = 1

	limitMessage = "ya te tomaste la keratina de hoy, aprende a tener límites xfi"

	ErrKeratineLimitReached = errors.New("keratine limit reached")

	mapKeratineMessagesBySource = ErrorMessages{
		ErrKeratineLimitReached: {
			TGSource:  limitMessage,
			APISource: limitMessage,
		},
	}
)

type KeratineTracker struct {
	activityType shared.Activity
	sourceType   SourceType
}

func NewKeratineTracker(activityType shared.Activity, source SourceType) (Tracker, error) {
	return &KeratineTracker{
		activityType: activityType,
		sourceType:   source,
	}, nil
}

func (t *KeratineTracker) Track(ctx context.Context, username string, content string) error {
	isGoalCompleted, err := isKeratineGoalCompleted(ctx, username)
	if err != nil {
		return err
	}

	if isGoalCompleted {
		return ErrKeratineLimitReached
	}

	now, err := shared.GetNow()
	if err != nil {
		return err
	}

	nowStr := now.Format(time.RFC3339)

	userActivity := shared.UserActivity{
		ID:        shared.GenerateActivityItemID(now, username, shared.Keratine),
		Name:      username,
		Activity:  shared.Keratine,
		CreatedAt: nowStr,
		Content:   content,
	}

	return storage.Create(ctx, userActivity)
}

func isKeratineGoalCompleted(ctx context.Context, username string) (bool, error) {
	currentDayKeratineActivities, err := storage.GetCurrentDayActivities(ctx, username, shared.Keratine)
	if err != nil {
		return false, err
	}

	return len(currentDayKeratineActivities) == goalKeratineConsume, nil
}

func (t *KeratineTracker) GetErrorMessage(err error) string {
	return GetErrorMessageByTracker(err, t.sourceType, mapKeratineMessagesBySource)
}

func (t *KeratineTracker) GetSuccessMessage() string {
	return "se wardó tu tomadita de keratina >:)"
}
