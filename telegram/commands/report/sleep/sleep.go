package sleep

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"context"
	"fmt"
	"time"
)

var (
	reportSleepMessage = `Pillá pues cómo son las vueltas precios@ %s 🍆

	Esta semana te has dormido esta cantidad de horas bb:

	Lunes: %s
	Martes: %s
	Miércoles: %s
	Jueves: %s
	Viernes: %s
	Sábado: %s
	Domingo: %s

	Si no queres tener ojeras andate a dormir 🐼
	`

	sleepGif = "https://media.giphy.com/media/v1.Y2lkPTc5MGI3NjExNHQ5aDZqZ3Q5cTRscGZodmVjMXdjNHM4bHV4dDA2dms4Z3Q4a3V6diZlcD12MV9naWZzX3NlYXJjaCZjdD1n/mkhMTALnrYRLnuoe5P/giphy.gif"
)

// SendSleepReport sends the report of sleep tracker
func SendSleepReport(ctx context.Context, client *shared.Client, userName string, _ string, chatID int64) error {
	report, err := GenerateSleepReport(ctx, client, userName, chatID)
	if err != nil {
		return client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	}

	err = client.SendMessage(chatID, report)
	if err != nil {
		return err
	}

	return client.SendAnimation(chatID, sleepGif)
}

func GenerateSleepReport(ctx context.Context, client *shared.Client, userName string, chatID int64) (string, error) {
	sleepActivities, err := storage.GetLastWeekUserHistoryPerActivity(ctx, userName, shared.Sleep)
	if err != nil {
		return "", client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	}

	timeDefault := "---"

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
			return "", client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
		}

		day := createdAt.Weekday()

		timeSleep[day] = activity.Content
	}

	report := fmt.Sprintf(reportSleepMessage, userName,
		timeSleep[time.Monday],
		timeSleep[time.Tuesday],
		timeSleep[time.Wednesday],
		timeSleep[time.Thursday],
		timeSleep[time.Friday],
		timeSleep[time.Saturday],
		timeSleep[time.Sunday])

	return report, nil
}
