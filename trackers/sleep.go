package trackers

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"context"
	"regexp"
	"time"
)

var (
	messageInvalidHour = "pone bien el gran hpta formato, por ejemplo 8h o 10m o 2h30m"

	successMessageSleep = "Que sueñes con los angelitos bb 😴😴😴"

	secondsRegex = `\d+s$`

	mapSleepMessagesBySource = ErrorMessages{
		ErrInvalidContent: {
			APISource: messageInvalidHour,
			TGSource:  messageInvalidHour,
		},
	}
)

type SleepTracker struct {
	activityType shared.Activity
	sourceType   SourceType
}

func NewSleepTracker(activityType shared.Activity, source SourceType) (Tracker, error) {
	return &SleepTracker{
		activityType: activityType,
		sourceType:   source,
	}, nil
}

func (t *SleepTracker) Track(ctx context.Context, username string, content string) error {
	if content == "" {
		return ErrInvalidContent
	}

	duration, err := time.ParseDuration(content)
	if err != nil {
		return ErrInvalidContent
	}

	isNewActivity, userActivity, err := getUserActivity(ctx, username, duration)
	if err != nil {
		return err
	}

	if isNewActivity {
		err = storage.Create(ctx, userActivity)
	} else {
		err = storage.UpdateContent(ctx, userActivity)
	}

	return err
}

func getUserActivity(ctx context.Context, userName string, duration time.Duration) (bool, shared.UserActivity, error) {
	activities, err := storage.GetCurrentDayActivities(ctx, userName, shared.Sleep)
	if err != nil {
		return false, shared.UserActivity{}, err
	}

	if len(activities) > 0 {
		activity := activities[0]

		err := addSleepTime(activity, duration)
		if err != nil {
			return false, shared.UserActivity{}, err
		}

		return false, *activity, nil
	}

	activity, err := shared.NewActivity(shared.Sleep, userName, durationToStringWithoutSeconds(duration))

	return true, activity, err
}

func addSleepTime(activity *shared.UserActivity, duration time.Duration) error {
	currentTime, err := time.ParseDuration(activity.Content)
	if err != nil {
		return err
	}

	newTime := currentTime + duration

	activity.Content = durationToStringWithoutSeconds(newTime)

	now, err := shared.GetNow()
	if err != nil {
		return err
	}

	nowStr := now.Format(time.RFC3339)
	activity.UpdatedAt = nowStr

	return nil
}

func durationToStringWithoutSeconds(duration time.Duration) string {
	re := regexp.MustCompile(secondsRegex)
	return re.ReplaceAllString(duration.String(), "")
}

func (t *SleepTracker) GetErrorMessage(err error) string {
	return GetErrorMessageByTracker(err, t.sourceType, mapSleepMessagesBySource)
}

func (t *SleepTracker) GetSuccessMessage() string {
	return successMessageSleep
}
