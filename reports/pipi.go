package reports

import (
	"activity-tracker/api/telegram"
	"activity-tracker/shared"
	"activity-tracker/storage"
	"fmt"
	"time"
)

var (
	reportPipiMessage = `Pillá pues cómo son las vueltas precioso %s 🍆
	
	Esta semana has miado así bb:

	Lunes: %d miadas 🍆
	Martes: %d miadas 🍆
	Miércoles: %d miadas 🍆
	Jueves: %d miadas 🍆
	Viernes: %d miadas 🍆
	Sábado: %d miadas 🍆
	Domingo: %d miadas 🍆

	Si querés miar más ponete a tomar awa en vez de pensar en tu ex 😘
	`
)

// GeneratePipiReport generates a weekly pipi report
func GeneratePipiReport(bot *telegram.Bot, userName string, chatID int64) (string, error) {
	pipiActivities, err := storage.GetLastWeekUserHistoryPerActivity(userName, shared.Pipi)
	if err != nil {
		return "", telegram.SendMessage(bot, chatID, "algo falló mi fafá: "+err.Error())
	}

	pipiPerDay := map[time.Weekday]int{
		time.Monday:    0,
		time.Tuesday:   0,
		time.Wednesday: 0,
		time.Thursday:  0,
		time.Friday:    0,
		time.Saturday:  0,
		time.Sunday:    0,
	}

	for _, activity := range pipiActivities {
		day := activity.CreatedAt.Weekday()
		pipiPerDay[day]++
	}

	report := fmt.Sprintf(reportPipiMessage, userName, pipiPerDay[time.Monday], pipiPerDay[time.Tuesday], pipiPerDay[time.Wednesday], pipiPerDay[time.Thursday], pipiPerDay[time.Friday], pipiPerDay[time.Saturday], pipiPerDay[time.Sunday])

	return report, nil
}
