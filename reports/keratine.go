package reports

import (
	"activity-tracker/api/telegram"
	"activity-tracker/shared"
	"activity-tracker/storage"
	"fmt"
	"time"
)

var (
	reportKeratineMessage = `Pillá pues cómo son las vueltas precios@ %s 🍆
	
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

	labelTookKeratine = "sisas"
)

// GenerateKeratineReport generates a weekly keratine report
func GenerateKeratineReport(bot *telegram.Bot, userName string, chatID int64) (string, error) {
	keratineActivities, err := storage.GetLastWeekUserHistoryPerActivity(userName, shared.Keratine)
	if err != nil {
		telegram.SendMessage(bot, chatID, "algo falló mi fafá: "+err.Error())

		return "", err
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
		day := activity.CreatedAt.Weekday()

		tookKeratine[day] = labelTookKeratine
	}

	report := fmt.Sprintf(reportKeratineMessage, userName, tookKeratine[time.Monday], tookKeratine[time.Tuesday], tookKeratine[time.Wednesday], tookKeratine[time.Thursday], tookKeratine[time.Friday], tookKeratine[time.Saturday], tookKeratine[time.Sunday])

	return report, nil
}
