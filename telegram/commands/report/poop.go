package report

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"fmt"
	"strconv"
	"time"
)

var (
	reportPoopMessage = `Mirá todas las veces que has cagado esta semana %s 🧐
	
	según ese intento de base de datos que tenemos, cagaste de la siguiente manera:

	Lunes: %d %s 💩
	Martes: %d %s 💩
	Miércoles: %d %s 💩
	Jueves: %d %s 💩
	Viernes: %d %s 💩
	Sábado: %d %s 💩
	Domingo: %d %s 💩

	`

	poopinGif = "https://media1.tenor.com/m/fUHxQ89S4uAAAAAC/kitten-cat.gif"
)

func getCagadaString(count int) string {
	if count == 1 {
		return "cagada"
	}
	return "cagadas"
}

func SendPoopReport(client *shared.Client, userName, content string, chatID int64) error {
	report, err := generatePoopReport(client, userName, chatID)
	if err != nil {
		return err
	}

	_ = client.SendMessage(chatID, report)

	return client.SendAnimation(chatID, poopinGif)
}

func generatePoopReport(client *shared.Client, userName string, chatID int64) (string, error) {
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

	report := fmt.Sprintf(reportPoopMessage, userName,
		poopsPerDay[time.Monday], getCagadaString(poopsPerDay[time.Monday]),
		poopsPerDay[time.Tuesday], getCagadaString(poopsPerDay[time.Tuesday]),
		poopsPerDay[time.Wednesday], getCagadaString(poopsPerDay[time.Wednesday]),
		poopsPerDay[time.Thursday], getCagadaString(poopsPerDay[time.Thursday]),
		poopsPerDay[time.Friday], getCagadaString(poopsPerDay[time.Friday]),
		poopsPerDay[time.Saturday], getCagadaString(poopsPerDay[time.Saturday]),
		poopsPerDay[time.Sunday], getCagadaString(poopsPerDay[time.Sunday]))

	return report, nil
}
