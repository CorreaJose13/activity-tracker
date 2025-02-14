package poop

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"fmt"
	"strconv"
	"time"
)

var (
	reportMessage = `Mirá todas las veces que has cagado esta semana %s 🧐
	
	según ese intento de base de datos que tenemos, cagaste de la siguiente manera:

	Lunes: %s 💩
	Martes: %s 💩
	Miércoles: %s 💩
	Jueves: %s 💩
	Viernes: %s 💩
	Sábado: %s 💩
	Domingo: %s 💩

	`

	poopinGif = "https://media1.tenor.com/m/fUHxQ89S4uAAAAAC/kitten-cat.gif"
)

func getCagadaString(count int) string {
	if count == 0 {
		return "no cagaste bb"
	}

	if count == 1 {
		return "1 cagada"
	}

	return strconv.Itoa(count) + " cagadas"
}

func SendPoopReport(client *shared.Client, userName, content string, chatID int64) error {
	report, err := GeneratePoopReport(client, userName, chatID)
	if err != nil {
		return err
	}

	err = client.SendMessage(chatID, report)
	if err != nil {
		return err
	}

	return client.SendAnimation(chatID, poopinGif)
}

func GeneratePoopReport(client *shared.Client, userName string, chatID int64) (string, error) {
	poopActivities, err := storage.GetLastWeekUserHistoryPerActivity(userName, shared.Poop)
	if err != nil {
		return "", client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	}

	poopsPerDay := map[time.Weekday]int{
		time.Monday:    0,
		time.Tuesday:   0,
		time.Wednesday: 0,
		time.Thursday:  0,
		time.Friday:    0,
		time.Saturday:  0,
		time.Sunday:    0,
	}

	for _, activity := range poopActivities {
		createdAt, err := time.Parse(time.RFC3339, activity.CreatedAt)
		if err != nil {
			return "", client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
		}

		day := createdAt.Weekday()

		num, err := strconv.Atoi(activity.Content)
		if err != nil {
			return "", client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
		}

		poopsPerDay[day] += num
	}

	report := fmt.Sprintf(reportMessage, userName,
		getCagadaString(poopsPerDay[time.Monday]),
		getCagadaString(poopsPerDay[time.Tuesday]),
		getCagadaString(poopsPerDay[time.Wednesday]),
		getCagadaString(poopsPerDay[time.Thursday]),
		getCagadaString(poopsPerDay[time.Friday]),
		getCagadaString(poopsPerDay[time.Saturday]),
		getCagadaString(poopsPerDay[time.Sunday]))

	return report, nil
}
