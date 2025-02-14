package report

import (
	"activity-tracker/shared"
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
	"os"
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
	generateReportErrorMessage = "Error generando reporte"
)

// GenerateAllReports generates all reports and send it in txt file
func GenerateAllReports(ctx context.Context, client *shared.Client, userName, content string, chatID int64) error {
	reports := ""
	filePath := os.TempDir() + "/all_reports.txt"

	for _, fn := range reportsFunctions {
		report, err := fn(ctx, client, userName, chatID)
		if err != nil {
			_ = client.SendMessage(chatID, generateReportErrorMessage)

			continue
		}

		reports += report + "\n\n"
	}

	err := os.WriteFile(filePath, []byte(reports), 0644)
	if err != nil {
		return err
	}

	return client.SendFile(chatID, filePath)
}
