package report

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"fmt"
	"time"
)

var (
	timeFormat   = "03:04 pm"
	locationTime = "America/Bogota"

	reportSleepMessage = `Pillá pues cómo son las vueltas precios@ %s 🍆

	Esta semana te has acostado a estas horas bb:

	Lunes: %s
	Martes: %s
	Miércoles: %s
	Jueves: %s
	Viernes: %s
	Sábado: %s
	Domingo: %s

	Si no queres tener ojeras andate a dormir 🐼
	`
)

func SendSleepReport(client *shared.Client, userName string, _ string, chatID int64) error {
	kr, err := generateSleepReport(userName)
	if err != nil {
		return client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	}

	return client.SendMessage(chatID, kr)
}

func generateSleepReport(userName string) (string, error) {
	sleepActivities, err := storage.GetLastWeekUserHistoryPerActivity(userName, shared.Sleep)
	if err != nil {
		return "", err
	}

	timeDefault := "--:-- --"

	timeSleep := map[time.Weekday]string{
		time.Monday:    timeDefault,
		time.Tuesday:   timeDefault,
		time.Wednesday: timeDefault,
		time.Thursday:  timeDefault,
		time.Friday:    timeDefault,
		time.Saturday:  timeDefault,
		time.Sunday:    timeDefault,
	}

	for _, activity := range sleepActivities {
		createdAt, err := time.Parse(time.RFC3339, activity.CreatedAt)
		if err != nil {
			return "", err
		}

		location, err := time.LoadLocation(locationTime) // o cualquier ubicación en -05:00
		if err != nil {
			return "", err
		}

		day := createdAt.Weekday()

		timeSleep[day] = createdAt.In(location).Format(timeFormat)
	}

	return fmt.Sprintf(reportSleepMessage, userName, timeSleep[time.Monday], timeSleep[time.Tuesday], timeSleep[time.Wednesday], timeSleep[time.Thursday], timeSleep[time.Friday], timeSleep[time.Saturday], timeSleep[time.Sunday]), nil
}
