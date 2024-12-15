package report

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"fmt"
	"time"
)

var (
	reportToothMessage = `
	%s - Esta semana te has cepillado por día esta cantidad 🪥🪥:

	Lunes: %d
	Martes: %d 
	Miércoles: %d 
	Jueves: %d
	Viernes: %d
	Sábado: %d
	Domingo: %d

	Con estos números se avecina una caries ni la hp, ponete juicios@ pues 🤢
	`
)

// SendToothReport sends the tooth report
func SendToothReport(bot *shared.Bot, userName, content string, chatID int64) error {
	wr, err := GenerateToothReport(bot, userName, chatID)
	if err != nil {
		return err
	}

	return shared.SendMessage(bot, chatID, wr)
}

func GenerateToothReport(bot *shared.Bot, userName string, chatID int64) (string, error) {
	toothActivities, err := storage.GetLastWeekUserHistoryPerActivity(userName, "tooth")
	if err != nil {
		return "", shared.SendMessage(bot, chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	}

	toothPerDay := map[time.Weekday]int{
		time.Monday:    0,
		time.Tuesday:   0,
		time.Wednesday: 0,
		time.Thursday:  0,
		time.Friday:    0,
		time.Saturday:  0,
		time.Sunday:    0,
	}

	for _, activity := range toothActivities {
		createdAt, err := time.Parse(time.RFC3339, activity.CreatedAt)
		if err != nil {
			return "", shared.SendMessage(bot, chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
		}

		day := createdAt.Weekday()

		toothPerDay[day]++
	}

	report := fmt.Sprintf(reportToothMessage, userName, toothPerDay[time.Monday], toothPerDay[time.Tuesday], toothPerDay[time.Wednesday], toothPerDay[time.Thursday], toothPerDay[time.Friday], toothPerDay[time.Saturday], toothPerDay[time.Sunday])

	return report, nil
}
