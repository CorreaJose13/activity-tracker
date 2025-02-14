package pipi

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"context"
	"fmt"
	"strconv"
	"time"
)

var (
	reportMessage = `Pillá pues cómo son las vueltas precios@ %s 🍆
	
	Esta semana has miado así bb:

	Lunes: %s 🍆
	Martes: %s 🍆
	Miércoles: %s 🍆
	Jueves: %s 🍆
	Viernes: %s 🍆
	Sábado: %s 🍆
	Domingo: %s 🍆

	Si querés miar más ponete a tomar awa en vez de pensar en tu ex 😘
	`

	pipiGif = "https://media.giphy.com/media/z0b9YVvaAQZe8/giphy.gif?cid=790b76112exkvfjoxs001tnxfa0pgac7vj1m27mcjhiyeizf&ep=v1_gifs_search&rid=giphy.gif&ct=g"
)

func getPipiString(count int) string {
	if count == 0 {
		return "no miaste bb"
	}

	if count == 1 {
		return "1 miada"
	}

	return strconv.Itoa(count) + " miadas"
}

// SendPipiReport sends the pipi report
func SendPipiReport(ctx context.Context, client *shared.Client, userName, content string, chatID int64) error {
	pr, err := GeneratePipiReport(ctx, client, userName, chatID)
	if err != nil {
		return err
	}

	err = client.SendMessage(chatID, pr)
	if err != nil {
		return err
	}

	return client.SendAnimation(chatID, pipiGif)
}

func GeneratePipiReport(ctx context.Context, client *shared.Client, userName string, chatID int64) (string, error) {
	pipiActivities, err := storage.GetLastWeekUserHistoryPerActivity(ctx, userName, shared.Pipi)
	if err != nil {
		return "", client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
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
		createdAt, err := time.Parse(time.RFC3339, activity.CreatedAt)
		if err != nil {
			return "", client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
		}

		day := createdAt.Weekday()

		pipiPerDay[day]++
	}

	report := fmt.Sprintf(reportMessage, userName,
		getPipiString(pipiPerDay[time.Monday]),
		getPipiString(pipiPerDay[time.Tuesday]),
		getPipiString(pipiPerDay[time.Wednesday]),
		getPipiString(pipiPerDay[time.Thursday]),
		getPipiString(pipiPerDay[time.Friday]),
		getPipiString(pipiPerDay[time.Saturday]),
		getPipiString(pipiPerDay[time.Sunday]))

	return report, nil
}
