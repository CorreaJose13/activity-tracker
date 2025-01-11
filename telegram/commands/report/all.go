package report

import (
	"activity-tracker/shared"
	"os"
)

var (
	reportsFunctions = []func(bot *shared.Client, userName string, chatID int64) (string, error){
		GenerateKeratineReport,
		GeneratePipiReport,
		GenerateShowerReport,
		GenerateToothReport,
		GenerateWaterReport,
	}
	generateReportErrorMessage = "Error generando reporte"
)

// GenerateAllReports generates all reports and send it in txt file
func GenerateAllReports(client *shared.Client, userName, content string, chatID int64) error {
	reports := ""
	filePath := "all_reports.txt"

	for _, fn := range reportsFunctions {
		report, err := fn(client, userName, chatID)
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
