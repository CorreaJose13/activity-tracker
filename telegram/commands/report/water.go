package report

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"fmt"
	"time"
)

var (
	reportMessage = `Pillá pues cómo son las vueltas precioso %s 🍆
	
	Esta semana has tomado awita así bb:

	Lunes: %dL
	Martes: %dL
	Miércoles: %dL
	Jueves: %dL
	Viernes: %dL
	Sábado: %dL
	Domingo: %dL

	Si querés mejorar estos números ponete a tomar awa en vez de pensar en tu ex 😘
	`
)

// SendWaterReport sends the water report
func SendWaterReport(bot *shared.Bot, userName, content string, chatID int64) error {
	wr, err := GenerateWaterReport(bot, userName, chatID)
	if err != nil {
		return err
	}

	return shared.SendMessage(bot, chatID, wr)
}

func GenerateWaterReport(bot *shared.Bot, userName string, chatID int64) (string, error) {
	waterActivities, err := storage.GetLastWeekUserHistoryPerActivity(userName, "water")
	if err != nil {
		return "", shared.SendMessage(bot, chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	}

	waterPerDay := map[time.Weekday]int{
		time.Monday:    0,
		time.Tuesday:   0,
		time.Wednesday: 0,
		time.Thursday:  0,
		time.Friday:    0,
		time.Saturday:  0,
		time.Sunday:    0,
	}

	for _, activity := range waterActivities {
		createdAt, err := time.Parse(time.RFC3339, activity.CreatedAt)
		if err != nil {
			return "", shared.SendMessage(bot, chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
		}

		day := createdAt.Weekday()

		waterPerDay[day]++
	}

	report := fmt.Sprintf(reportMessage, userName, waterPerDay[time.Monday], waterPerDay[time.Tuesday], waterPerDay[time.Wednesday], waterPerDay[time.Thursday], waterPerDay[time.Friday], waterPerDay[time.Saturday], waterPerDay[time.Sunday])

	return report, nil
}
