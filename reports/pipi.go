package reports

import (
	"activity-tracker/shared"
	"fmt"
	"strconv"
)

var (
	reportPipiMessage = `Pillá pues cómo son las vueltas precios@ %s 🍆

	Esta semana has miado así bb:

	Lunes: %s 🍆
	Martes: %s 🍆
	Miércoles: %s 🍆
	Jueves: %s 🍆
	Viernes: %s 🍆
	Sábado: %s 🍆
	Domingo: %s 🍆

	Si querés miar más ponete a tomar awa en vez de pensar en tu ex 😘
	`

	pipiGif = "https://media.giphy.com/media/z0b9YVvaAQZe8/giphy.gif?cid=790b76112exkvfjoxs001tnxfa0pgac7vj1m27mcjhiyeizf&ep=v1_gifs_search&rid=giphy.gif&ct=g"
)

type Pipi struct {
	activityType      shared.Activity
	username          string
	dailyReports      []DailyReport
	isReportGenerated bool
	period            Period
}

func NewPipiReport(username string) (Report, error) {
	return &Pipi{
		activityType: shared.Pipi,
		username:     username,
		period:       PeriodWeekly,
	}, nil
}

func (r *Pipi) DailyReports() []DailyReport {
	if !r.isReportGenerated {
		PanicReportNotGenerated()
	}

	return r.dailyReports
}

func (r *Pipi) GetGifURL() string {
	return pipiGif
}

func (r *Pipi) GetPeriod() Period {
	return r.period
}

func (r *Pipi) GetUsername() string {
	return r.username
}

func (r *Pipi) GetActivity() shared.Activity {
	return r.activityType
}

func (r *Pipi) SetDailyReports(dailyReport []DailyReport) {
	r.dailyReports = dailyReport
	r.isReportGenerated = true
}

func (r *Pipi) String() string {
	if !r.isReportGenerated {
		PanicReportNotGenerated()
	}

	return fmt.Sprintf(reportPipiMessage, r.username,
		getPipiString(r.dailyReports, 0),
		getPipiString(r.dailyReports, 1),
		getPipiString(r.dailyReports, 2),
		getPipiString(r.dailyReports, 3),
		getPipiString(r.dailyReports, 4),
		getPipiString(r.dailyReports, 5),
		getPipiString(r.dailyReports, 6),
	)
}

func getPipiString(reports []DailyReport, index int) string {
	count := float64(0)

	if len(reports) > index {
		count = reports[index].Result
	}

	if count == 0 {
		return "no miaste bb"
	}

	if count == 1 {
		return "1 miada"
	}

	return strconv.Itoa(int(count)) + " miadas"
}
