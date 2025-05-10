package reports

import (
	"activity-tracker/shared"
	"fmt"
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

	keratineGif       = "https://media.giphy.com/media/v1.Y2lkPTc5MGI3NjExbWJjMHVzMXEwb2F0dGYzOWJlM3Njbnc5OXh1bnB5aDN2eHg4MmZxdyZlcD12MV9naWZzX3NlYXJjaCZjdD1n/D7z8JfNANqahW/giphy.gif"
	labelTookKeratine = "sisas"
)

type Keratine struct {
	activityType      shared.Activity
	username          string
	dailyReports      []DailyReport
	isReportGenerated bool
	period            Period
}

func NewKeratineReport(username string) (Report, error) {
	return &Keratine{
		activityType: shared.Keratine,
		username:     username,
		period:       PeriodWeekly,
	}, nil
}

func (r *Keratine) DailyReports() []DailyReport {
	if !r.isReportGenerated {
		PanicReportNotGenerated()
	}

	return r.dailyReports
}

func (r *Keratine) GetGifURL() string {
	return keratineGif
}

func (r *Keratine) GetPeriod() Period {
	return r.period
}

func (r *Keratine) GetUsername() string {
	return r.username
}

func (r *Keratine) GetActivity() shared.Activity {
	return r.activityType
}

func (r *Keratine) SetDailyReports(dailyReport []DailyReport) {
	r.dailyReports = dailyReport
	r.isReportGenerated = true
}

func (r *Keratine) String() string {
	if !r.isReportGenerated {
		PanicReportNotGenerated()
	}

	return fmt.Sprintf(reportKeratineMessage, r.username,
		getKeatineString(r.dailyReports, 0),
		getKeatineString(r.dailyReports, 1),
		getKeatineString(r.dailyReports, 2),
		getKeatineString(r.dailyReports, 3),
		getKeatineString(r.dailyReports, 4),
		getKeatineString(r.dailyReports, 5),
		getKeatineString(r.dailyReports, 6),
	)
}

func getKeatineString(reports []DailyReport, index int) string {
	count := float64(0)

	if len(reports) > index {
		count = reports[index].Result
	}

	if count > 0 {
		return labelTookKeratine
	}

	return ""
}
