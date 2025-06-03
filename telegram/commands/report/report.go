package report

import (
	"activity-tracker/reports"
	"activity-tracker/shared"
	"context"
	"fmt"
)

func SendReportActivity(ctx context.Context, client *shared.Client, activityType shared.Activity, username, content string, chatID int64, sendGif bool) error {
	report, err := reports.NewReport(activityType, username)
	if err != nil {
		return client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	}

	err = reports.AddDailyReports(ctx, report)
	if err != nil {
		return client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	}

	err = client.SendMessage(chatID, report.String())
	if err != nil {
		return client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	}

	if sendGif {
		return client.SendAnimation(chatID, report.GetGifURL())
	}

	return nil
}

func SendReportAllActivities(ctx context.Context, client *shared.Client, username string, activities []shared.Activity, content string, chatID int64) error {
	for _, activity := range activities {
		err := SendReportActivity(ctx, client, activity, username, content, chatID, false)
		if err != nil {
			client.SendMessage(chatID, fmt.Sprintf("Error generating report of activity: %s. Error: %s", activity, err.Error()))
		}
	}

	return nil
}
