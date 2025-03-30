package trackers

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"context"
	"time"
)

var (
	successMessageCyclingWithContent = "ve pero y ese poco de kilometros? te perseguía un veneco o q? anwy congrats"
	successMessageCycling            = "ataca pogachaaaaaa 🚴"

	mapCyclingMessagesBySource = map[error]map[SourceType]string{}
)

type Cycling struct {
	activityType shared.Activity
	sourceType   SourceType
	hasContent   bool
}

func NewCyclingTracker(activityType shared.Activity, source SourceType) (Tracker, error) {
	return &Cycling{
		activityType: activityType,
		sourceType:   source,
	}, nil
}

func (t *Cycling) Track(ctx context.Context, username string, content string) error {
	now, err := shared.GetNow()
	if err != nil {
		return err
	}

	nowStr := now.Format(time.RFC3339)

	userActivity := shared.UserActivity{
		ID:        shared.GenerateActivityItemID(now, username, shared.Cycling),
		Name:      username,
		Activity:  shared.Cycling,
		CreatedAt: nowStr,
		Content:   content,
	}

	return storage.Create(ctx, userActivity)
}

func (t *Cycling) GetErrorMessage(err error) string {
	return GetErrorMessageByTracker(err, t.sourceType, mapCyclingMessagesBySource)
}

func (t *Cycling) GetSuccessMessage() string {
	if t.hasContent {
		return successMessageCyclingWithContent
	}

	return successMessageCycling
}
