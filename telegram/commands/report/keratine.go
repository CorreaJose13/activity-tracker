package report

import (
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

// SendKeratineReport sends the keratine report
func SendKeratineReport(bot *shared.Bot, userName, content string, chatID int64) error {
	kr, err := generateKeratineReport(bot, userName, chatID)
	if err != nil {
		return err
	}

	return shared.SendMessage(bot, chatID, kr)
}

func generateKeratineReport(bot *shared.Bot, userName string, chatID int64) (string, error) {
	keratineActivities, err := storage.GetLastWeekUserHistoryPerActivity(userName, shared.Keratine)
	if err != nil {
		return "", shared.SendMessage(bot, chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
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
			return "", shared.SendMessage(bot, chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
		}

		day := createdAt.Weekday()

		tookKeratine[day] = labelTookKeratine
	}

	report := fmt.Sprintf(reportKeratineMessage, userName, tookKeratine[time.Monday], tookKeratine[time.Tuesday], tookKeratine[time.Wednesday], tookKeratine[time.Thursday], tookKeratine[time.Friday], tookKeratine[time.Saturday], tookKeratine[time.Sunday])

	return report, nil
}
