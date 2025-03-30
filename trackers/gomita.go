package trackers

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrGomitaInvalidContent = errors.New("invalid number of gomitas")

	missingGomitaContentMessage = "tan trabao estás?? mandame la info así vicioso:\n\t /track gomita <fraction or float>"
	invalidGomitaContentMessage = "concentrate hijo de toda tu puta madre, escribe bien el número 😡"

	successMessageGomita = "buen biaje mi sielo 🌈"

	mapGomitaMessagesBySource = map[error]map[SourceType]string{
		ErrInvalidContent: {
			APISource: "tan trabao estás?? mandame la info así vicioso",
			TGSource:  "tan trabao estás?? mandame la info así vicioso:\n\t /track gomita <fraction or float>",
		},
		ErrGomitaInvalidContent: {
			APISource: invalidGomitaContentMessage,
			TGSource:  invalidGomitaContentMessage,
		},
	}
)

type Gomita struct {
	activityType shared.Activity
	sourceType   SourceType
	hasContent   bool
}

func NewGomitaTracker(activityType shared.Activity, source SourceType) (Tracker, error) {
	return &Gomita{
		activityType: activityType,
		sourceType:   source,
	}, nil
}

func (t *Gomita) Track(ctx context.Context, username string, content string) error {
	if content == "" {
		return ErrInvalidContent
	}

	if strings.Contains(content, "/") {
		floatString, err := fractionToFloatString(content)
		if err != nil {
			return ErrGomitaInvalidContent
		}

		content = floatString
	} else if !shared.IsValidFloat(content) {
		return ErrGomitaInvalidContent
	}

	userActivity, err := shared.NewActivity(shared.Gomita, username, content)
	if err != nil {
		return err
	}

	return storage.Create(ctx, userActivity)
}

func fractionToFloatString(content string) (string, error) {
	split := strings.Split(content, "/")
	num, err := strconv.ParseFloat(split[0], 64)
	if err != nil {
		return "", err
	}

	den, err := strconv.ParseFloat(split[1], 64)
	if err != nil {
		return "", err
	}

	result := num / den
	floatString := fmt.Sprintf("%f", result)

	return floatString, nil
}

func (t *Gomita) GetErrorMessage(err error) string {
	return GetErrorMessageByTracker(err, t.sourceType, mapGomitaMessagesBySource)
}

func (t *Gomita) GetSuccessMessage() string {
	return successMessageGomita
}
