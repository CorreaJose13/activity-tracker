package water

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"context"
	"fmt"
	"strconv"
	"time"
)

var (
	reportMessage = `Pillá pues cómo son las vueltas precioso %s 🍆
	
	Esta semana has tomado awita así bb:

	Lunes: %s
	Martes: %s
	Miércoles: %s
	Jueves: %s
	Viernes: %s
	Sábado: %s
	Domingo: %s

	Si querés mejorar estos números ponete a tomar awa en vez de pensar en tu ex 😘
	`

	waterGif = "https://media.giphy.com/media/Tu1p1x4QwnZAc/giphy.gif?cid=790b7611pq6cg14nw5waf7t19fv2p6y988jgv3hhgaoaawm9&ep=v1_gifs_search&rid=giphy.gif&ct=g"
)

func getWaterString(count int) string {
	if count == 0 {
		return "no tomaste awita bb 😱"
	}

	if count == 1 {
		return "1 litro 💧"
	}

	return strconv.Itoa(count) + " litros 💧"
}

// SendWaterReport sends the water report
func SendWaterReport(ctx context.Context, client *shared.Client, userName, content string, chatID int64) error {
	wr, err := GenerateWaterReport(ctx, client, userName, chatID)
	if err != nil {
		return err
	}

	err = client.SendMessage(chatID, wr)
	if err != nil {
		return err
	}

	return client.SendAnimation(chatID, waterGif)
}

func GenerateWaterReport(ctx context.Context, client *shared.Client, userName string, chatID int64) (string, error) {
	waterActivities, err := storage.GetLastWeekUserHistoryPerActivity(ctx, userName, shared.Water)
	if err != nil {
		return "", client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	}

	waterPerDay := map[time.Weekday]int{
		time.Monday:    0,
		time.Tuesday:   0,
		time.Wednesday: 0,
		time.Thursday:  0,
		time.Friday:    0,
		time.Saturday:  0,
		time.Sunday:    0,
	}

	for _, activity := range waterActivities {
		createdAt, err := time.Parse(time.RFC3339, activity.CreatedAt)
		if err != nil {
			return "", client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
		}

		day := createdAt.Weekday()

		waterPerDay[day]++
	}

	report := fmt.Sprintf(reportMessage, userName,
		getWaterString(waterPerDay[time.Monday]),
		getWaterString(waterPerDay[time.Tuesday]),
		getWaterString(waterPerDay[time.Wednesday]),
		getWaterString(waterPerDay[time.Thursday]),
		getWaterString(waterPerDay[time.Friday]),
		getWaterString(waterPerDay[time.Saturday]),
		getWaterString(waterPerDay[time.Sunday]))

	return report, nil
}
