package shower

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"context"
	"fmt"
	"time"
)

var (
	reportShowerMessage = `Después de ir a echarle un ojo al baño me encontré este resumen precios@ %s 🛀🏻

	Reporte semanal:
	Lunes: %s
	Martes: %s
	Miércoles: %s
	Jueves: %s
	Viernes: %s
	Sábado: %s
	Domingo: %s

	Última racha:
	Bañándose: %d %s 🚿
	Sin bañarse: %d %s 🤢

	Andá bañate pa' que estos números mejoren, cochin@ de 💩💩`

	oneDayLabel        = "día"
	textSisasStreak    = "días"
	textNonasStreak    = "días"
	labelTookShower    = "sisas"
	labelNotTookShower = "nonas"
	oneDayStreak       = 1
)

func SendShowerReport(ctx context.Context, client *shared.Client, userName, content string, chatID int64) error {
	reportMessage, err := GenerateShowerReport(ctx, client, userName, chatID)
	if err != nil {
		return err
	}

	return client.SendMessage(chatID, reportMessage)
}

func GenerateShowerReport(ctx context.Context, client *shared.Client, userName string, chatID int64) (string, error) {
	showerActivities, err := storage.GetLastWeekUserHistoryPerActivity(ctx, userName, shared.Shower)
	if err != nil {
		return "", client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	}

	tookShower := map[time.Weekday]string{
		time.Monday:    labelNotTookShower,
		time.Tuesday:   labelNotTookShower,
		time.Wednesday: labelNotTookShower,
		time.Thursday:  labelNotTookShower,
		time.Friday:    labelNotTookShower,
		time.Saturday:  labelNotTookShower,
		time.Sunday:    labelNotTookShower,
	}

	for _, activity := range showerActivities {
		createdAt, err := time.Parse(time.RFC3339, activity.CreatedAt)
		if err != nil {
			return "", client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
		}

		tookShower[createdAt.Weekday()] = labelTookShower
	}
	return calculateShowerStreaks(tookShower, userName)
}

func calculateShowerStreaks(tookShower map[time.Weekday]string, userName string) (string, error) {
	var currentSisasStreak, currentNonasStreak, lastNonasStreak, lastSisasStreak int

	days := []string{
		tookShower[time.Monday],
		tookShower[time.Tuesday],
		tookShower[time.Wednesday],
		tookShower[time.Thursday],
		tookShower[time.Friday],
		tookShower[time.Saturday],
		tookShower[time.Sunday],
	}

	for _, stringValue := range days {
		if stringValue == labelTookShower {
			currentSisasStreak++

			if currentNonasStreak != 0 {
				lastNonasStreak = currentNonasStreak
			}

			currentNonasStreak = 0

			continue
		}

		currentNonasStreak++

		if currentSisasStreak != 0 {
			lastSisasStreak = currentSisasStreak
		}

		currentSisasStreak = 0
	}

	if currentNonasStreak != 0 {
		lastNonasStreak = currentNonasStreak
	}

	if currentSisasStreak != 0 {
		lastSisasStreak = currentSisasStreak
	}

	if lastSisasStreak == oneDayStreak {
		textSisasStreak = oneDayLabel
	}

	if lastNonasStreak == oneDayStreak {
		textNonasStreak = oneDayLabel
	}

	return fmt.Sprintf(reportShowerMessage,
		userName,
		tookShower[time.Monday],
		tookShower[time.Tuesday],
		tookShower[time.Wednesday],
		tookShower[time.Thursday],
		tookShower[time.Friday],
		tookShower[time.Saturday],
		tookShower[time.Sunday],
		lastSisasStreak, textSisasStreak,
		lastNonasStreak, textNonasStreak,
	), nil
}
