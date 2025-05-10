package reports

import (
	"activity-tracker/shared"
	"fmt"
)

var (
	reportWaterMessage = `Veamos cuánta agua has tomado esta semana %s 💧

	Según tus registros, tus vasos de agua fueron:

	Lunes: %s
	Martes: %s
	Miércoles: %s
	Jueves: %s
	Viernes: %s
	Sábado: %s
	Domingo: %s

	No te deshidrates, toma más agua 💦
	`

	waterGif = "https://media.giphy.com/media/3o7TKFvVQZvQZvQZvQ/giphy.gif"
)

type Water struct {
	activityType      shared.Activity
	username          string
	dailyReports      []DailyReport
	isReportGenerated bool
	period            Period
}

func NewWaterReport(username string) (Report, error) {
	return &Water{
		activityType: shared.Water,
		username:     username,
		period:       PeriodWeekly,
	}, nil
}

func (r *Water) DailyReports() []DailyReport {
	if !r.isReportGenerated {
		PanicReportNotGenerated()
	}

	return r.dailyReports
}

func (r *Water) GetGifURL() string {
	return waterGif
}

func (r *Water) GetPeriod() Period {
	return r.period
}

func (r *Water) GetUsername() string {
	return r.username
}

func (r *Water) GetActivity() shared.Activity {
	return r.activityType
}

func (r *Water) SetDailyReports(dailyReport []DailyReport) {
	r.dailyReports = dailyReport
	r.isReportGenerated = true
}

func (r *Water) String() string {
	if !r.isReportGenerated {
		PanicReportNotGenerated()
	}

	return fmt.Sprintf(reportWaterMessage, r.username,
		getWaterString(r.dailyReports, 0),
		getWaterString(r.dailyReports, 1),
		getWaterString(r.dailyReports, 2),
		getWaterString(r.dailyReports, 3),
		getWaterString(r.dailyReports, 4),
		getWaterString(r.dailyReports, 5),
		getWaterString(r.dailyReports, 6),
	)
}

func getWaterString(reports []DailyReport, index int) string {
	count := float64(0)

	if len(reports) > index {
		count = reports[index].Result
	}

	if count == 0 {
		return "ni un vaso, deshidratado 💀"
	}

	if count == 1 {
		return "1 vaso 💧"
	}

	return fmt.Sprintf("%.0f", count) + " vasos 💧"
}
