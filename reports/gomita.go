package reports

import (
	"activity-tracker/shared"
	"fmt"
	"math/big"
)

var (
	reportGomitaMessage = `Aquí tienes un resumen de cuantas veces te trabaste con gomita esta semana  %s 🍬🌿

	De acuerdo a la cochinada de base de datos que tenemos, esto es lo que encontré:

	Lunes: %s
	Martes: %s
	Miércoles: %s
	Jueves: %s
	Viernes: %s
	Sábado: %s
	Domingo: %s

	Ponte las pilas pa consumir más 🥵
	`

	gomitaGif = "https://media.giphy.com/media/rWiEbamfqOHrq/giphy.gif?cid=790b76113to9r3hgdf5j4317otxivd5ldi4gg7pug36eam97&ep=v1_gifs_search&rid=giphy.gif&ct=g"
)

type Gomita struct {
	activityType      shared.Activity
	username          string
	dailyReports      []DailyReport
	isReportGenerated bool
	period            Period
}

func NewGomitaReport(username string) (Report, error) {
	return &Gomita{
		activityType: shared.Gomita,
		username:     username,
		period:       PeriodWeekly,
	}, nil
}

func (r *Gomita) DailyReports() []DailyReport {
	if !r.isReportGenerated {
		PanicReportNotGenerated()
	}

	return r.dailyReports
}

func (r *Gomita) GetGifURL() string {
	return gomitaGif
}

func (r *Gomita) GetPeriod() Period {
	return r.period
}

func (r *Gomita) GetUsername() string {
	return r.username
}

func (r *Gomita) GetActivity() shared.Activity {
	return r.activityType
}

func (r *Gomita) SetDailyReports(dailyReports []DailyReport) {
	r.dailyReports = dailyReports
	r.isReportGenerated = true
}

func (r *Gomita) String() string {
	if !r.isReportGenerated {
		PanicReportNotGenerated()
	}

	return fmt.Sprintf(reportGomitaMessage, r.username,
		getGomitaString(r.dailyReports, 0),
		getGomitaString(r.dailyReports, 1),
		getGomitaString(r.dailyReports, 2),
		getGomitaString(r.dailyReports, 3),
		getGomitaString(r.dailyReports, 4),
		getGomitaString(r.dailyReports, 5),
		getGomitaString(r.dailyReports, 6),
	)
}

func floatToFraction(f float64) (p, q int64) {
	bf := big.NewFloat(f)

	bf.SetPrec(64)

	rat := new(big.Rat)
	rat.SetFloat64(f)

	p = rat.Num().Int64()
	q = rat.Denom().Int64()

	return p, q
}

func getGomitaString(reports []DailyReport, index int) string {
	count := float64(0)

	if len(reports) > index {
		count = reports[index].Result
	}

	if count == 0 {
		return "ni una bb :("
	}

	if count < 1 {
		p, q := floatToFraction(count)
		return fmt.Sprintf("%d/%d", p, q) + " de gomita 🍁"
	}

	if count == 1 {
		return "1 gomita 🍁"
	}

	return fmt.Sprintf("%.2f", count) + " gomitas 🍁"
}
