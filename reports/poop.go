package reports

import (
	"activity-tracker/shared"
	"fmt"
	"strconv"
)

var (
	reportPoopMessage = `Mirá todas las veces que has cagado esta semana %s 🧐

	según ese intento de base de datos que tenemos, cagaste de la siguiente manera:

	Lunes: %s 💩
	Martes: %s 💩
	Miércoles: %s 💩
	Jueves: %s 💩
	Viernes: %s 💩
	Sábado: %s 💩
	Domingo: %s 💩
	`

	poopingGif = "https://media1.tenor.com/m/fUHxQ89S4uAAAAAC/kitten-cat.gif"
)

type Poop struct {
	activityType      shared.Activity
	username          string
	dailyReports      []DailyReport
	isReportGenerated bool
	period            Period
}

func NewPoopReport(username string) (Report, error) {
	return &Poop{
		activityType: shared.Poop,
		username:     username,
		period:       PeriodWeekly,
	}, nil
}

func (r *Poop) DailyReports() []DailyReport {
	if !r.isReportGenerated {
		PanicReportNotGenerated()
	}

	return r.dailyReports
}

func (r *Poop) GetGifURL() string {
	return poopingGif
}

func (r *Poop) GetPeriod() Period {
	return r.period
}

func (r *Poop) GetUsername() string {
	return r.username
}

func (r *Poop) GetActivity() shared.Activity {
	return r.activityType
}

func (r *Poop) SetDailyReports(dailyReport []DailyReport) {
	r.dailyReports = dailyReport
	r.isReportGenerated = true
}

func (r *Poop) String() string {
	if !r.isReportGenerated {
		PanicReportNotGenerated()
	}

	return fmt.Sprintf(reportPoopMessage, r.username,
		getPoopString(r.dailyReports, 0),
		getPoopString(r.dailyReports, 1),
		getPoopString(r.dailyReports, 2),
		getPoopString(r.dailyReports, 3),
		getPoopString(r.dailyReports, 4),
		getPoopString(r.dailyReports, 5),
		getPoopString(r.dailyReports, 6),
	)
}

func getPoopString(reports []DailyReport, index int) string {
	count := float64(0)

	if len(reports) > index {
		count = reports[index].Result
	}

	if count == 0 {
		return "no cagaste bb"
	}

	if count == 1 {
		return "1 cagada"
	}

	return strconv.Itoa(int(count)) + " cagadas"
}
