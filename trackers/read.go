package trackers

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"context"
	"errors"
)

var (
	missingReadContentMessage = "y vos qué creés que voy a trackear si no me mandas el número de páginas que te leíste 🐸"
	invalidReadContentMessage = "es muy difícil mandar un número bien? conectá solo 2 neuronas y mandá un número de páginas 🐸"

	successMessageRead = "congrats por decidir dejar la ignorancia un poquito atrás, seguí leyendo a ver si cambiamos este cochino país 👃🏾◻️"

	ErrReadInvalidNumber = errors.New("read invalid page number")

	mapReadMessagesBySource = ErrorMessages{
		ErrInvalidContent: {
			APISource: missingGomitaContentMessage,
			TGSource:  missingGomitaContentMessage,
		},
		ErrReadInvalidNumber: {
			APISource: invalidReadContentMessage,
			TGSource:  invalidReadContentMessage,
		},
	}
)

type ReadTracker struct {
	activityType shared.Activity
	sourceType   SourceType
}

func NewReadTracker(activityType shared.Activity, source SourceType) (Tracker, error) {
	return &ReadTracker{
		activityType: activityType,
		sourceType:   source,
	}, nil
}

func (t *ReadTracker) Track(ctx context.Context, username string, content string) error {
	if content == "" {
		return ErrInvalidContent
	}

	if !shared.IsValidNumber(content) {
		return ErrReadInvalidNumber
	}

	userActivity, err := shared.NewActivity(shared.Read, username, content)
	if err != nil {
		return err
	}

	return storage.Create(ctx, userActivity)
}

func (t *ReadTracker) GetErrorMessage(err error) string {
	return GetErrorMessageByTracker(err, t.sourceType, mapReadMessagesBySource)
}

func (t *ReadTracker) GetSuccessMessage() string {
	return successMessageRead
}
