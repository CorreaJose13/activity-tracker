package reports

import (
	"activity-tracker/shared"
	"fmt"
)

var (
	reportShowerMessage = `Miremos cuántas veces te has bañado esta semana %s 🚿

	Según la base de datos, tus duchas fueron así:

	Lunes: %s
	Martes: %s
	Miércoles: %s
	Jueves: %s
	Viernes: %s
	Sábado: %s
	Domingo: %s

	No seas cochino, báñate más seguido 🧼
	`

	showerGif = "https://media.giphy.com/media/3o7TKFvVQZvQZvQZvQ/giphy.gif"
)

type Shower struct {
	activityType      shared.Activity
	username          string
	dailyReports      []DailyReport
	isReportGenerated bool
	period            Period
}

func NewShowerReport(username string) (Report, error) {
	return &Shower{
		activityType: shared.Shower,
		username:     username,
		period:       PeriodWeekly,
	}, nil
}

func (r *Shower) DailyReports() []DailyReport {
	if !r.isReportGenerated {
		PanicReportNotGenerated()
	}

	return r.dailyReports
}

func (r *Shower) GetGifURL() string {
	return showerGif
}

func (r *Shower) GetPeriod() Period {
	return r.period
}

func (r *Shower) GetUsername() string {
	return r.username
}

func (r *Shower) GetActivity() shared.Activity {
	return r.activityType
}

func (r *Shower) SetDailyReports(dailyReport []DailyReport) {
	r.dailyReports = dailyReport
	r.isReportGenerated = true
}

func (r *Shower) String() string {
	if !r.isReportGenerated {
		PanicReportNotGenerated()
	}

	return fmt.Sprintf(reportShowerMessage, r.username,
		getShowerString(r.dailyReports, 0),
		getShowerString(r.dailyReports, 1),
		getShowerString(r.dailyReports, 2),
		getShowerString(r.dailyReports, 3),
		getShowerString(r.dailyReports, 4),
		getShowerString(r.dailyReports, 5),
		getShowerString(r.dailyReports, 6),
	)
}

func getShowerString(reports []DailyReport, index int) string {
	count := float64(0)

	if len(reports) > index {
		count = reports[index].Result
	}

	if count == 0 {
		return "ni una ducha, cochino 🤢"
	}

	if count == 1 {
		return "1 ducha 🚿"
	}

	return fmt.Sprintf("%.0f", count) + " duchas 🚿"
}
