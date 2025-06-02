package report

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"activity-tracker/telegram/commands/report/keratine"
	"activity-tracker/telegram/commands/report/pipi"
	"activity-tracker/telegram/commands/report/poop"
	"activity-tracker/telegram/commands/report/read"
	"activity-tracker/telegram/commands/report/run"
	"activity-tracker/telegram/commands/report/shower"
	"activity-tracker/telegram/commands/report/sleep"
	"activity-tracker/telegram/commands/report/tooth"
	"activity-tracker/telegram/commands/report/water"
	"context"
	"fmt"
	"os"
	"strconv"
	"time"
)

var (
	reportsFunctions = []func(ctx context.Context, bot *shared.Client, userName string, chatID int64) (string, error){
		keratine.GenerateKeratineReport,
		pipi.GeneratePipiReport,
		poop.GeneratePoopReport,
		read.GenerateReadReport,
		run.GenerateRunReport,
		shower.GenerateShowerReport,
		sleep.GenerateSleepReport,
		tooth.GenerateToothReport,
		water.GenerateWaterReport,
	}

	generateReportErrorMessage = "Error generando reporte. Error: %s"
	currentMonthReportMsg      = "Lista de actividades del mes %s - %s - %s\n\n%s"
	invalidActivityMsg         = "kejesa mondá"
	noActivitiesMsg            = "ni una perra actividad este mes 😐"
)

// GenerateAllReports generates all reports and send it in txt file
func GenerateAllReports(ctx context.Context, client *shared.Client, userName, content string, chatID int64) error {
	reports := ""
	filePath := os.TempDir() + "/all_reports.txt"

	for _, fn := range reportsFunctions {
		report, err := fn(ctx, client, userName, chatID)
		if err != nil {
			_ = client.SendMessage(chatID, fmt.Sprintf(generateReportErrorMessage, err.Error()))

			continue
		}

		reports += report + "\n\n"
	}

	err := os.WriteFile(filePath, []byte(reports), 0644)
	if err != nil {
		return client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	}

	return client.SendFile(chatID, filePath)
}

// GenerateMonthlyReport generates a monthly report for a user
func GenerateMonthlyReport(ctx context.Context, client *shared.Client, userName, activity string, chatID int64) error {
	parsedActivity := shared.Activity(activity)

	if !shared.IsValidActivity(parsedActivity) {
		return client.SendMessage(chatID, invalidActivityMsg)
	}

	activities, err := storage.GetCurrentMonthUserHistoryPerActivity(ctx, userName, parsedActivity)
	if err != nil {
		return client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	}

	if len(activities) == 0 {
		return client.SendMessage(chatID, noActivitiesMsg)
	}

	activitiesMsg := "Cantidad de actividades: " + strconv.Itoa(len(activities)) + "\n\n"

	for _, a := range activities {
		createdAtTime, _ := time.Parse(time.RFC3339, a.CreatedAt)
		formattedDate := createdAtTime.Format("02 / January / 2006 15:04")

		activitiesMsg += fmt.Sprintf("--> %s - %s\n\n", formattedDate, a.Content)
	}

	now, err := shared.GetNow()
	if err != nil {
		return client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	}

	return client.SendMessage(chatID, fmt.Sprintf(currentMonthReportMsg, userName, activity, now.Month(), activitiesMsg))
}
