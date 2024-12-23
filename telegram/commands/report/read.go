package report

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"fmt"
	"strconv"
	"time"
)

var (
	reportReadMessage = `Miremos cómo te ha ido con la lectura esta semana %s 🥸
	
	según la esquizofrenia de la base de datos, por día leíste la siguiente cantidad de páginas:

	Lunes: %d páginas 📖
	Martes: %d páginas 📖
	Miércoles: %d páginas 📖
	Jueves: %d páginas 📖
	Viernes: %d páginas 📖
	Sábado: %d páginas 📖
	Domingo: %d páginas 📖

	Pa' que dejés de decir 'haiga' toca leer un toque más 🙄
	`
)

// SendReadReport sends the read report
func SendReadReport(client *shared.Client, userName, content string, chatID int64) error {
	report, err := generateReadReport(client, userName, chatID)
	if err != nil {
		return err
	}

	return client.SendMessage(chatID, report)
}

func generateReadReport(client *shared.Client, userName string, chatID int64) (string, error) {
	readActivities, err := storage.GetLastWeekUserHistoryPerActivity(userName, shared.Read)
	if err != nil {
		return "", client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	}

	pagesPerDay := map[time.Weekday]int{
		time.Monday:    0,
		time.Tuesday:   0,
		time.Wednesday: 0,
		time.Thursday:  0,
		time.Friday:    0,
		time.Saturday:  0,
		time.Sunday:    0,
	}

	for _, activity := range readActivities {
		createdAt, err := time.Parse(time.RFC3339, activity.CreatedAt)
		if err != nil {
			return "", client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
		}

		day := createdAt.Weekday()

		num, err := strconv.Atoi(activity.Content)
		if err != nil {
			return "", client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
		}

		pagesPerDay[day] += num
	}

	report := fmt.Sprintf(reportReadMessage, userName, pagesPerDay[time.Monday], pagesPerDay[time.Tuesday], pagesPerDay[time.Wednesday], pagesPerDay[time.Thursday], pagesPerDay[time.Friday], pagesPerDay[time.Saturday], pagesPerDay[time.Sunday])

	return report, nil
}
