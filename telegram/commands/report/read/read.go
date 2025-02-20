package read

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"context"
	"fmt"
	"strconv"
	"time"
)

var (
	reportReadMessage = `Miremos cómo te ha ido con la lectura esta semana %s 🥸
	
	según la esquizofrenia de la base de datos, por día leíste la siguiente cantidad de páginas:

	Lunes: %s
	Martes: %s
	Miércoles: %s
	Jueves: %s
	Viernes: %s
	Sábado: %s
	Domingo: %s

	Pa' que dejés de decir 'haiga' toca leer un toque más 🙄
	`

	readingGif = "https://media.giphy.com/media/v1.Y2lkPTc5MGI3NjExNTI4bWdsNGZhcWNxNWpyam05czVkMjd1OGt2YjNjd2w0YzR4dHM4biZlcD12MV9naWZzX3NlYXJjaCZjdD1n/WoWm8YzFQJg5i/giphy.gif"
)

func getReadString(count int) string {
	if count == 0 {
		return "sos un analfabeta"
	}

	if count == 1 {
		return "1 perra página a lo bien??? 🤨"
	}

	return strconv.Itoa(count) + " páginas 📖"
}

// SendReadReport sends the report of read tracker
func SendReadReport(ctx context.Context, client *shared.Client, userName, content string, chatID int64) error {
	report, err := GenerateReadReport(ctx, client, userName, chatID)
	if err != nil {
		return err
	}

	err = client.SendMessage(chatID, report)
	if err != nil {
		return err
	}

	return client.SendAnimation(chatID, readingGif)
}

func GenerateReadReport(ctx context.Context, client *shared.Client, userName string, chatID int64) (string, error) {
	readActivities, err := storage.GetLastWeekUserHistoryPerActivity(ctx, userName, shared.Read)
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

	report := fmt.Sprintf(reportReadMessage, userName,
		getReadString(pagesPerDay[time.Monday]),
		getReadString(pagesPerDay[time.Tuesday]),
		getReadString(pagesPerDay[time.Wednesday]),
		getReadString(pagesPerDay[time.Thursday]),
		getReadString(pagesPerDay[time.Friday]),
		getReadString(pagesPerDay[time.Saturday]),
		getReadString(pagesPerDay[time.Sunday]))

	return report, nil
}
