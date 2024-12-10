package report

import (
	"activity-tracker/shared"
	"fmt"
	"time"
)

var (
	reportShowerMessage = `Pillá pues cómo son las vueltas precios@ %s 🍆

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
Sin bañarse: %d %s 🤢`

	textSisasStreak = "días"
	textNonasStreak = "días"
)

func SendShowerReport(bot *shared.Bot, userName, content string, chatID int64) error {
	a, b, c, d, err := generateShowerReport(bot, userName, chatID)
	if err != nil {
		return err
	}
	fmt.Println(a, b, c, d)

	return shared.SendMessage(bot, chatID, "reportMessage")
}

func generateShowerReport(bot *shared.Bot, userName string, chatID int64) (int, int, int, int, error) {
	// showerActivities, err := storage.GetLastWeekUserHistoryPerActivity(userName, shared.Shower)
	// if err != nil {
	// 	return "", shared.SendMessage(bot, chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	// }
	labelTookShower := "sisas"
	labelNotTookShower := "nonas"

	tookShower := map[time.Weekday]string{
		time.Monday:    labelNotTookShower,
		time.Tuesday:   labelTookShower,
		time.Wednesday: labelTookShower,
		time.Thursday:  labelNotTookShower,
		time.Friday:    labelNotTookShower,
		time.Saturday:  labelTookShower,
		time.Sunday:    labelNotTookShower,
	}

	// for _, activity := range showerActivities {
	// 	createdAt, err := time.Parse(time.RFC3339, activity.CreatedAt)
	// 	if err != nil {
	// 		return "", shared.SendMessage(bot, chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	// 	}
	// 	tookShower[createdAt.Weekday()] = labelTookShower
	// }

	currentSisasStreak := 0
	currentNonasStreak := 0
	lastNonasStreak := 0
	lastSisasStreak := 0

	for _, stringValue := range tookShower {
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

	if lastSisasStreak == 1 {
		textSisasStreak = "día"
	} else {
		textSisasStreak = "días"
	}
	if lastNonasStreak == 1 {
		textNonasStreak = "día"
	} else {
		textNonasStreak = "días"
	}
	// return fmt.Sprintf(reportShowerMessage,
	// 	userName,
	// 	tookShower[time.Monday],
	// 	tookShower[time.Tuesday],
	// 	tookShower[time.Wednesday],
	// 	tookShower[time.Thursday],
	// 	tookShower[time.Friday],
	// 	tookShower[time.Saturday],
	// 	tookShower[time.Sunday],
	// 	lastSisasStreak, textSisasStreak,
	// 	lastNonasStreak, textNonasStreak,
	// ), nil
	fmt.Println("lastSisas", lastSisasStreak, "lastNonas", lastNonasStreak, "currentNonas", currentNonasStreak, "currentSisas", currentSisasStreak)
	return lastSisasStreak, lastNonasStreak, currentNonasStreak, currentSisasStreak, nil
}
