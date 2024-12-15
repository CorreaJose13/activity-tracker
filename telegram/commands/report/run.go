package report

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"fmt"
	"time"
)

var (
	reportRunMessage = `Llegaron las métricas atleta de mierda %s 🏃🏾‍♂️

	Lunes: %s
	Martes: %s
	Miércoles: %s
	Jueves: %s
	Viernes: %s
	Sábado: %s
	Domingo: %s

	Si no querés las pantorrillas lesionadas de johan ponete a correr un poquito más 🤢
	`

	labelTookRun = "sisas"
)

// SendRunReport sends the run report
func SendRunReport(bot *shared.Bot, userName, content string, chatID int64) error {
	kr, err := GenerateRunReport(bot, userName, chatID)
	if err != nil {
		return err
	}

	return shared.SendMessage(bot, chatID, kr)
}

// GenerateRunReport generates the run report
func GenerateRunReport(bot *shared.Bot, userName string, chatID int64) (string, error) {
	runActivities, err := storage.GetLastWeekUserHistoryPerActivity(userName, shared.Run)
	if err != nil {
		return "", shared.SendMessage(bot, chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	}

	labelBoolDefault := "nonas"

	tookRun := map[time.Weekday]string{
		time.Monday:    labelBoolDefault,
		time.Tuesday:   labelBoolDefault,
		time.Wednesday: labelBoolDefault,
		time.Thursday:  labelBoolDefault,
		time.Friday:    labelBoolDefault,
		time.Saturday:  labelBoolDefault,
		time.Sunday:    labelBoolDefault,
	}

	for _, activity := range runActivities {
		createdAt, err := time.Parse(time.RFC3339, activity.CreatedAt)
		if err != nil {
			return "", shared.SendMessage(bot, chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
		}

		day := createdAt.Weekday()

		tookRun[day] = labelTookRun
	}

	report := fmt.Sprintf(reportRunMessage, userName, tookRun[time.Monday], tookRun[time.Tuesday], tookRun[time.Wednesday], tookRun[time.Thursday], tookRun[time.Friday], tookRun[time.Saturday], tookRun[time.Sunday])

	return report, nil
}
