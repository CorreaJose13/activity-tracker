package reports

import (
	"activity-tracker/shared"
	"fmt"
)

var (
	reportRunMessage = `Llegaron tus métricas de mierda subatleta %s 🏃🏾‍♂️

	Lunes: %s
	Martes: %s
	Miércoles: %s
	Jueves: %s
	Viernes: %s
	Sábado: %s
	Domingo: %s

	Si no querés que te robe alguien de sucre comenzá a correr más 🤢
	`

	runninGif = "https://media.giphy.com/media/XyNMDllviurb3bpfDE/giphy.gif?cid=ecf05e47xnoxtwqok8a12t2uqy5bsqr1z0iwo057vl78ik3b&ep=v1_gifs_search&rid=giphy.gif&ct=g"
)

type Run struct {
	activityType      shared.Activity
	username          string
	dailyReports      []DailyReport
	isReportGenerated bool
	period            Period
}

func NewRunReport(username string) (Report, error) {
	return &Run{
		activityType: shared.Run,
		username:     username,
		period:       PeriodWeekly,
	}, nil
}

func (r *Run) DailyReports() []DailyReport {
	if !r.isReportGenerated {
		PanicReportNotGenerated()
	}

	return r.dailyReports
}

func (r *Run) GetGifURL() string {
	return runninGif
}

func (r *Run) GetPeriod() Period {
	return r.period
}

func (r *Run) GetUsername() string {
	return r.username
}

func (r *Run) GetActivity() shared.Activity {
	return r.activityType
}

func (r *Run) SetDailyReports(dailyReport []DailyReport) {
	r.dailyReports = dailyReport
	r.isReportGenerated = true
}

func (r *Run) String() string {
	if !r.isReportGenerated {
		PanicReportNotGenerated()
	}

	return fmt.Sprintf(reportRunMessage, r.username,
		getRunString(r.dailyReports, 0),
		getRunString(r.dailyReports, 1),
		getRunString(r.dailyReports, 2),
		getRunString(r.dailyReports, 3),
		getRunString(r.dailyReports, 4),
		getRunString(r.dailyReports, 5),
		getRunString(r.dailyReports, 6),
	)
}

func getRunString(reports []DailyReport, index int) string {
	count := float64(0)

	if len(reports) > index {
		count = reports[index].Result
	}

	if count == 0 {
		return "ni un metro, perezoso 🤢"
	}

	return fmt.Sprintf("%.2f", count) + " kms 🏃🏾‍♂️"
}
