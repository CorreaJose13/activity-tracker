package reports

import (
	"activity-tracker/api/telegram"
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

// GenerateWaterReport generates a weekly water report
func GenerateWaterReport(bot *telegram.Bot, userName string, chatID int64) (string, error) {
	waterActivities, err := storage.GetLastWeekUserHistoryPerActivity(userName, "water")
	if err != nil {
		telegram.SendMessage(bot, chatID, "algo falló mi fafá: "+err.Error())

		return "", err
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
		day := activity.CreatedAt.Weekday()
		waterPerDay[day]++
	}

	report := fmt.Sprintf(reportMessage, userName, waterPerDay[time.Monday], waterPerDay[time.Tuesday], waterPerDay[time.Wednesday], waterPerDay[time.Thursday], waterPerDay[time.Friday], waterPerDay[time.Saturday], waterPerDay[time.Sunday])

	return report, nil
}
