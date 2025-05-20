package reports

import (
	"activity-tracker/shared"
	"fmt"
)

var (
	reportToothMessage = `Miremos cuántas veces te has cepillado los dientes esta semana %s 🪥

	Según la base de datos, tus cepilladas fueron:

	Lunes: %s
	Martes: %s
	Miércoles: %s
	Jueves: %s
	Viernes: %s
	Sábado: %s
	Domingo: %s

	No querés que se te caigan los dientes, ¿o sí? 🦷
	`

	toothGif = "https://media.giphy.com/media/3o7TKFvVQZvQZvQZvQ/giphy.gif"
)

type Tooth struct {
	activityType      shared.Activity
	username          string
	dailyReports      []DailyReport
	isReportGenerated bool
	period            Period
}

func NewToothReport(username string) (Report, error) {
	return &Tooth{
		activityType: shared.ToothBrush,
		username:     username,
		period:       PeriodWeekly,
	}, nil
}

func (r *Tooth) DailyReports() []DailyReport {
	if !r.isReportGenerated {
		PanicReportNotGenerated()
	}

	return r.dailyReports
}

func (r *Tooth) GetGifURL() string {
	return toothGif
}

func (r *Tooth) GetPeriod() Period {
	return r.period
}

func (r *Tooth) GetUsername() string {
	return r.username
}

func (r *Tooth) GetActivity() shared.Activity {
	return r.activityType
}

func (r *Tooth) SetDailyReports(dailyReport []DailyReport) {
	r.dailyReports = dailyReport
	r.isReportGenerated = true
}

func (r *Tooth) String() string {
	if !r.isReportGenerated {
		PanicReportNotGenerated()
	}

	return fmt.Sprintf(reportToothMessage, r.username,
		getToothString(r.dailyReports, 0),
		getToothString(r.dailyReports, 1),
		getToothString(r.dailyReports, 2),
		getToothString(r.dailyReports, 3),
		getToothString(r.dailyReports, 4),
		getToothString(r.dailyReports, 5),
		getToothString(r.dailyReports, 6),
	)
}

func getToothString(reports []DailyReport, index int) string {
	count := float64(0)

	if len(reports) > index {
		count = reports[index].Result
	}

	if count == 0 {
		return "ni una cepillada, cochino 🦷"
	}

	if count == 1 {
		return "1 cepillada 🪥"
	}

	return fmt.Sprintf("%.0f", count) + " cepilladas 🪥"
}
