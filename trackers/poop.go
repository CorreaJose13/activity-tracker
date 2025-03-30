package trackers

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"context"
	"errors"
)

var (
	invalidPoopContentMessage = "al parecer la kk la tienes en el cerebro, manda el número bien ps 😡"

	successMessagePoop = "a ber de q color es? 👀👀"

	ErrPoopInvalidNumber = errors.New("invalid poop number")

	mapPoopMessagesBySource = ErrorMessages{
		ErrInvalidContent: {
			APISource: "sos down???? mandame la info así",
			TGSource:  "sos down???? mandame la info así:\n\t /track poop <times u pooped today>",
		},
		ErrPoopInvalidNumber: {
			APISource: invalidGomitaContentMessage,
			TGSource:  invalidGomitaContentMessage,
		},
	}
)

type PoopTracker struct {
	activityType shared.Activity
	sourceType   SourceType
}

func NewPoopTracker(activityType shared.Activity, source SourceType) (Tracker, error) {
	return &PoopTracker{
		activityType: activityType,
		sourceType:   source,
	}, nil
}

func (t *PoopTracker) Track(ctx context.Context, username string, content string) error {
	if content == "" {
		return ErrInvalidContent
	}

	if !shared.IsValidNumber(content) {
		return ErrPoopInvalidNumber
	}

	userActivity, err := shared.NewActivity(shared.Poop, username, content)
	if err != nil {
		return err
	}

	return storage.Create(ctx, userActivity)
}

func (t *PoopTracker) GetErrorMessage(err error) string {
	return GetErrorMessageByTracker(err, t.sourceType, mapPoopMessagesBySource)
}

func (t *PoopTracker) GetSuccessMessage() string {
	return successMessagePoop
}
