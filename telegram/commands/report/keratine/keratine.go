package keratine

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"context"
	"fmt"
	"time"
)

var (
	reportMessage = `Pillá pues cómo son las vueltas precios@ %s 🍆
	
	Esta semana has tomado keratina así bb:

	Lunes: %s
	Martes: %s
	Miércoles: %s
	Jueves: %s
	Viernes: %s
	Sábado: %s
	Domingo: %s

	Si querés mejorar estos números ponete las pilas con la keratina 😘
	`

	keratineGif       = "https://media.giphy.com/media/v1.Y2lkPTc5MGI3NjExbWJjMHVzMXEwb2F0dGYzOWJlM3Njbnc5OXh1bnB5aDN2eHg4MmZxdyZlcD12MV9naWZzX3NlYXJjaCZjdD1n/D7z8JfNANqahW/giphy.gif"
	labelTookKeratine = "sisas"
)

// SendKeratineReport sends the keratine report
func SendKeratineReport(ctx context.Context, client *shared.Client, userName, content string, chatID int64) error {
	kr, err := GenerateKeratineReport(ctx, client, userName, chatID)
	if err != nil {
		return err
	}

	err = client.SendMessage(chatID, kr)
	if err != nil {
		return err
	}

	return client.SendAnimation(chatID, keratineGif)
}

// GenerateKeratineReport generates the keratine report
func GenerateKeratineReport(ctx context.Context, client *shared.Client, userName string, chatID int64) (string, error) {
	keratineActivities, err := storage.GetLastWeekUserHistoryPerActivity(ctx, userName, shared.Keratine)
	if err != nil {
		return "", client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	}

	labelBoolDefault := "nonas"

	tookKeratine := map[time.Weekday]string{
		time.Monday:    labelBoolDefault,
		time.Tuesday:   labelBoolDefault,
		time.Wednesday: labelBoolDefault,
		time.Thursday:  labelBoolDefault,
		time.Friday:    labelBoolDefault,
		time.Saturday:  labelBoolDefault,
		time.Sunday:    labelBoolDefault,
	}

	for _, activity := range keratineActivities {
		createdAt, err := time.Parse(time.RFC3339, activity.CreatedAt)
		if err != nil {
			return "", client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
		}

		day := createdAt.Weekday()

		tookKeratine[day] = labelTookKeratine
	}

	report := fmt.Sprintf(reportMessage, userName,
		tookKeratine[time.Monday],
		tookKeratine[time.Tuesday],
		tookKeratine[time.Wednesday],
		tookKeratine[time.Thursday],
		tookKeratine[time.Friday],
		tookKeratine[time.Saturday],
		tookKeratine[time.Sunday])

	return report, nil
}
