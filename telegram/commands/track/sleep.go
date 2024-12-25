package track

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"errors"
	"fmt"
	"regexp"
	"time"
)

const (
	sleepMessage = "Espero estar en esos ricos sueños 😏"
)

var (
	errInvalidTimeFormat = errors.New("invalid time format")
	messageInvalidHour   = "Ve pone bien el formato de horas, por ejemplo 8h o 10m o 2h30m"
	secondsRegex         = `\d+s$`
)

// SendTrackSleep tracks the sleep activity
func SendTrackSleep(client *shared.Client, userName, content string, chatID int64) error {
	if content == "" {
		return client.SendMessage(chatID, messageInvalidHour)
	}

	duration, err := time.ParseDuration(content)
	if err != nil {
		return client.SendMessage(chatID, messageInvalidHour)
	}

	isNewActivity, userActivity, err := getUserActivity(userName, duration)
	if err != nil {
		return client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	}

	if isNewActivity {
		err = storage.Create(userActivity)
	} else {
		err = storage.UpdateContent(userActivity)
	}

	if err != nil {
		return client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	}

	return client.SendMessage(chatID, sleepMessage)
}

func getUserActivity(userName string, duration time.Duration) (bool, shared.UserActivity, error) {
	activities, err := storage.GetCurrentDayActivities(userName, shared.Sleep)
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
