package track

import (
	"activity-tracker/shared"
	"activity-tracker/trackers"
	"context"
)

// SendTrackActivity tracks the activities
func SendTrackActivity(ctx context.Context, activity shared.Activity, client *shared.Client, username, content string, chatID int64) error {
	tracker, err := trackers.NewTracker(activity, trackers.TGSource)
	if err != nil {
		return HandleTrackErrors(ctx, err, client, chatID, tracker)
	}

	err = tracker.Track(ctx, username, content)
	if err != nil {
		return HandleTrackErrors(ctx, err, client, chatID, tracker)
	}

	return client.SendMessage(chatID, tracker.GetSuccessMessage())
}

func HandleTrackErrors(_ context.Context, err error, client *shared.Client, chatID int64, tracker trackers.Tracker) error {
	if tracker == nil {
		errSendMessage := client.SendMessage(chatID, err.Error())
		if errSendMessage != nil {
			return errSendMessage
		}

		return err
	}

	return client.SendMessage(chatID, tracker.GetErrorMessage(err))
}
