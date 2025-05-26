package report

import (
	"activity-tracker/reports"
	"activity-tracker/shared"
	"context"
	"fmt"
)

func SendReportActicvity(ctx context.Context, client *shared.Client, activityType shared.Activity, username, content string, chatID int64) error {
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

	return client.SendAnimation(chatID, report.GetGifURL())
}
