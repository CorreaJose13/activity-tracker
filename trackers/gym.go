package trackers

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"context"
	"errors"
	"strings"
	"time"
)

var (
	ErrGymMissingDuration = errors.New("missing duration")
	ErrGymMissingMuscle   = errors.New("missing muscle")

	missingDurationMessage = "eh pero cuánto tiempo te ejercitaste sapa inmunda"
	missingMuscleMessage   = "eh pero hiciste chisme al fallo o q 🐸? mandame el musculo que ejercitaste sapa. Ej: bicep,pecho,jeta"

	mapGymMessagesBySource = ErrorMessages{
		ErrInvalidContent: {
			APISource: "eh pero vos sos tonto o te haces? mandame el time y el muscle",
			TGSource:  "eh pero vos sos tonto o te haces? mandame el time y el muscle. asi:\n\t /track gym <duration> <muscles by comma>",
		},
		ErrGymMissingDuration: {
			APISource: missingDurationMessage,
			TGSource:  missingDurationMessage,
		},
		ErrGymMissingMuscle: {
			APISource: missingMuscleMessage,
			TGSource:  missingMuscleMessage,
		},
	}
)

type GymTracker struct {
	activityType shared.Activity
	sourceType   SourceType
}

func NewGymTracker(activityType shared.Activity, source SourceType) (Tracker, error) {
	return &GymTracker{
		activityType: activityType,
		sourceType:   source,
	}, nil
}

func (t *GymTracker) Track(ctx context.Context, username string, content string) error {
	duration, muscle, ok := strings.Cut(content, " ")
	if !ok {
		return ErrInvalidContent
	}

	if duration == "" {
		return ErrGymMissingDuration
	}

	if muscle == "" {
		return ErrGymMissingMuscle
	}

	now, err := shared.GetNow()
	if err != nil {
		return err
	}

	nowStr := now.Format(time.RFC3339)

	userActivity := shared.UserActivity{
		ID:        shared.GenerateActivityItemID(now, username, shared.Gym),
		Name:      username,
		Activity:  shared.Gym,
		CreatedAt: nowStr,
		Content:   content,
	}

	return storage.Create(ctx, userActivity)
}

func (t *GymTracker) GetErrorMessage(err error) string {
	return GetErrorMessageByTracker(err, t.sourceType, mapGymMessagesBySource)
}

func (t *GymTracker) GetSuccessMessage() string {
	return "isss mi papacho el próximo cbum ve"
}
