package tooth

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"context"
	"fmt"
	"strconv"
	"time"
)

var (
	reportToothMessage = `%s : Esta semana te has cepillado por día esta cantidad 🪥🪥:

	Lunes: %s
	Martes: %s 
	Miércoles: %s 
	Jueves: %s
	Viernes: %s
	Sábado: %s
	Domingo: %s

	Con estos números se avecina una caries ni la hp, ponete juicios@ pues 🤢
	`

	toothGif = "https://media.giphy.com/media/3xz2BNVxo4D9mO1qLu/giphy.gif?cid=790b7611dg6d3e5ymbqjq2vtpjk204ijinr6cbm25vy4ansh&ep=v1_gifs_search&rid=giphy.gif&ct=g"
)

func getToothString(count int) any {
	if count == 0 {
		return "ni una vez 🤢"
	}

	if count == 1 {
		return "1 vez 🪥"
	}

	return strconv.Itoa(count) + " veces 🪥"
}

// SendToothReport sends the tooth report
func SendToothReport(ctx context.Context, client *shared.Client, userName, content string, chatID int64) error {
	tr, err := GenerateToothReport(ctx, client, userName, chatID)
	if err != nil {
		return err
	}

	err = client.SendMessage(chatID, tr)
	if err != nil {
		return err
	}

	return client.SendAnimation(chatID, toothGif)
}

func GenerateToothReport(ctx context.Context, client *shared.Client, userName string, chatID int64) (string, error) {
	toothActivities, err := storage.GetLastWeekUserHistoryPerActivity(ctx, userName, shared.ToothBrush)
	if err != nil {
		return "", client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	}

	toothPerDay := map[time.Weekday]int{
		time.Monday:    0,
		time.Tuesday:   0,
		time.Wednesday: 0,
		time.Thursday:  0,
		time.Friday:    0,
		time.Saturday:  0,
		time.Sunday:    0,
	}

	for _, activity := range toothActivities {
		createdAt, err := time.Parse(time.RFC3339, activity.CreatedAt)
		if err != nil {
			return "", client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
		}

		day := createdAt.Weekday()

		toothPerDay[day]++
	}

	report := fmt.Sprintf(reportToothMessage, userName,
		getToothString(toothPerDay[time.Monday]),
		getToothString(toothPerDay[time.Tuesday]),
		getToothString(toothPerDay[time.Wednesday]),
		getToothString(toothPerDay[time.Thursday]),
		getToothString(toothPerDay[time.Friday]),
		getToothString(toothPerDay[time.Saturday]),
		getToothString(toothPerDay[time.Sunday]))

	return report, nil
}
