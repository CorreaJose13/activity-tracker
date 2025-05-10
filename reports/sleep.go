package reports

import (
	"activity-tracker/shared"
	"fmt"
)

var (
	reportSleepMessage = `Veamos cómo has dormido esta semana %s 😴

	Según tus registros, tus horas de sueño fueron:

	Lunes: %s
	Martes: %s
	Miércoles: %s
	Jueves: %s
	Viernes: %s
	Sábado: %s
	Domingo: %s

	Duerme más, que te ves cansado 😪
	`

	sleepGif = "https://media.giphy.com/media/3o7TKFvVQZvQZvQZvQ/giphy.gif"
)

type Sleep struct {
	activityType      shared.Activity
	username          string
	dailyReports      []DailyReport
	isReportGenerated bool
	period            Period
}

func NewSleepReport(username string) (Report, error) {
	return &Sleep{
		activityType: shared.Sleep,
		username:     username,
		period:       PeriodWeekly,
	}, nil
}

func (r *Sleep) DailyReports() []DailyReport {
	if !r.isReportGenerated {
		PanicReportNotGenerated()
	}

	return r.dailyReports
}

func (r *Sleep) GetGifURL() string {
	return sleepGif
}

func (r *Sleep) GetPeriod() Period {
	return r.period
}

func (r *Sleep) GetUsername() string {
	return r.username
}

func (r *Sleep) GetActivity() shared.Activity {
	return r.activityType
}

func (r *Sleep) SetDailyReports(dailyReport []DailyReport) {
	r.dailyReports = dailyReport
	r.isReportGenerated = true
}

func (r *Sleep) String() string {
	if !r.isReportGenerated {
		PanicReportNotGenerated()
	}

	return fmt.Sprintf(reportSleepMessage, r.username,
		getSleepString(r.dailyReports, 0),
		getSleepString(r.dailyReports, 1),
		getSleepString(r.dailyReports, 2),
		getSleepString(r.dailyReports, 3),
		getSleepString(r.dailyReports, 4),
		getSleepString(r.dailyReports, 5),
		getSleepString(r.dailyReports, 6),
	)
}

func getSleepString(reports []DailyReport, index int) string {
	count := float64(0)

	if len(reports) > index {
		count = reports[index].Result
	}

	if count == 0 {
		return "ni una hora, zombie 🧟‍♂️"
	}

	return fmt.Sprintf("%.1f", count) + " horas 😴"
}
