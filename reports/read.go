package reports

import (
	"activity-tracker/shared"
	"fmt"
	"strconv"
)

var (
	reportReadMessage = `Miremos cómo te ha ido con la lectura esta semana %s 🥸

	según la esquizofrenia de la base de datos, por día leíste la siguiente cantidad de páginas:

	Lunes: %s
	Martes: %s
	Miércoles: %s
	Jueves: %s
	Viernes: %s
	Sábado: %s
	Domingo: %s

	Pa' que dejés de decir 'haiga' toca leer un toque más 🙄
	`

	readingGif = "https://media.giphy.com/media/v1.Y2lkPTc5MGI3NjExNTI4bWdsNGZhcWNxNWpyam05czVkMjd1OGt2YjNjd2w0YzR4dHM4biZlcD12MV9naWZzX3NlYXJjaCZjdD1n/WoWm8YzFQJg5i/giphy.gif"
)

type Read struct {
	activityType      shared.Activity
	username          string
	dailyReports      []DailyReport
	isReportGenerated bool
	period            Period
}

func NewReadReport(username string) (Report, error) {
	return &Read{
		activityType: shared.Read,
		username:     username,
		period:       PeriodWeekly,
	}, nil
}

func (r *Read) DailyReports() []DailyReport {
	if !r.isReportGenerated {
		PanicReportNotGenerated()
	}

	return r.dailyReports
}

func (r *Read) GetGifURL() string {
	return readingGif
}

func (r *Read) GetPeriod() Period {
	return r.period
}

func (r *Read) GetUsername() string {
	return r.username
}

func (r *Read) GetActivity() shared.Activity {
	return r.activityType
}

func (r *Read) SetDailyReports(dailyReport []DailyReport) {
	r.dailyReports = dailyReport
	r.isReportGenerated = true
}

func (r *Read) String() string {
	if !r.isReportGenerated {
		PanicReportNotGenerated()
	}

	return fmt.Sprintf(reportReadMessage, r.username,
		getReadString(r.dailyReports, 0),
		getReadString(r.dailyReports, 1),
		getReadString(r.dailyReports, 2),
		getReadString(r.dailyReports, 3),
		getReadString(r.dailyReports, 4),
		getReadString(r.dailyReports, 5),
		getReadString(r.dailyReports, 6),
	)
}

func getReadString(reports []DailyReport, index int) string {
	count := float64(0)

	if len(reports) > index {
		count = reports[index].Result
	}

	if count == 0 {
		return "sos un analfabeta"
	}

	if count == 1 {
		return "1 perra página a lo bien??? 🤨"
	}

	return strconv.Itoa(int(count)) + " páginas 📖"
}
