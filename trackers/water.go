package trackers

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"context"
	"errors"
	"time"
)

var (
	goalWaterConsume = 3

	messageWaterLimit = "ya te tomaste los 3L de awa mi papacho, aprende a tener límites"

	successMessageWater = "se wardó tu tomadita de awa golosito"

	ErrWaterLimitReached = errors.New("water limit reached")

	mapWaterMessagesBySource = ErrorMessages{
		ErrWaterLimitReached: {
			APISource: messageWaterLimit,
			TGSource:  messageWaterLimit,
		},
	}
)

type WaterTracker struct {
	activityType shared.Activity
	sourceType   SourceType
}

func NewWaterTracker(activityType shared.Activity, source SourceType) (Tracker, error) {
	return &WaterTracker{
		activityType: activityType,
		sourceType:   source,
	}, nil
}

func (t *WaterTracker) Track(ctx context.Context, username string, content string) error {
	isGoalCompleted, err := isWaterGoalCompleted(ctx, username)
	if err != nil {
		return err
	}

	if isGoalCompleted {
		return ErrWaterLimitReached
	}

	now, err := shared.GetNow()
	if err != nil {
		return err
	}

	nowStr := now.Format(time.RFC3339)

	userActivity := shared.UserActivity{
		ID:        shared.GenerateActivityItemID(now, username, shared.Water),
		Name:      username,
		Activity:  shared.Water,
		CreatedAt: nowStr,
		Content:   content, // TODO: add logic to validate the content and use it in isGoalCompleted function
	}

	return storage.Create(ctx, userActivity)
}

func isWaterGoalCompleted(ctx context.Context, username string) (bool, error) {
	currentDayWaterActivities, err := storage.GetCurrentDayActivities(ctx, username, shared.Water)
	if err != nil {
		return false, err
	}

	return len(currentDayWaterActivities) >= goalWaterConsume, nil
}

func (t *WaterTracker) GetErrorMessage(err error) string {
	return GetErrorMessageByTracker(err, t.sourceType, mapWaterMessagesBySource)
}

func (t *WaterTracker) GetSuccessMessage() string {
	return successMessageWater
}
