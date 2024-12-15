package report

import (
	"activity-tracker/shared"
	"os"
)

var (
	reportsFunctions = []func(bot *shared.Bot, userName string, chatID int64) (string, error){
		GenerateKeratineReport,
		GeneratePipiReport,
		GenerateShowerReport,
		GenerateToothReport,
		GenerateWaterReport,
	}
)

// GenerateAllReports generates all reports and send it in txt file
func GenerateAllReports(bot *shared.Bot, userName, content string, chatID int64) error {
	reports := ""
	filePath := "all_reports.txt"

	for _, fn := range reportsFunctions {
		report, err := fn(bot, userName, chatID)
		if err != nil {
			_ = shared.SendMessage(bot, chatID, "f")

			continue
		}

		reports += report + "\n\n"
	}

	err := os.WriteFile(filePath, []byte(reports), 0644)
	if err != nil {
		return err
	}

	return shared.SendFile(bot, chatID, filePath)
}
