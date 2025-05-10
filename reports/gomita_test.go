package reports

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGomitaReport(t *testing.T) {
	c := require.New(t)

	report, err := NewGomitaReport("johanFlorez")
	c.NoError(err)

	report.SetDailyReports([]DailyReport{
		{
			Date:   time.Now(),
			Result: float64(5),
		},
	})

	c.Len(report.DailyReports(), 1)
	c.Equal(report.DailyReports()[0].Result, float64(5))
}

func TestPanicGettingDailyReports(t *testing.T) {
	c := require.New(t)

	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("Expected a panic")
		}

		c.Equal(r, "reports were not generated")
	}()

	report, err := NewGomitaReport("johanFlorez")
	c.NoError(err)

	report.DailyReports()
}

func TestGomitaReportString(t *testing.T) {
	c := require.New(t)

	report, err := NewGomitaReport("johanFlorez")
	c.NoError(err)

	now := time.Now()
	monday := now.AddDate(0, 0, -int(now.Weekday())+1)

	report.SetDailyReports([]DailyReport{
		{
			Date:   monday,
			Result: float64(5),
		},
	})

	rp := report.String()
	c.Equal(rp, "Aquí tienes un resumen de cuantas veces te trabaste con gomita esta semana  johanFlorez 🍬🌿\n\n\tDe acuerdo a la cochinada de base de datos que tenemos, esto es lo que encontré:\n\n\tLunes: 5.00 gomitas 🍁\n\tMartes: ni una bb :(\n\tMiércoles: ni una bb :(\n\tJueves: ni una bb :(\n\tViernes: ni una bb :(\n\tSábado: ni una bb :(\n\tDomingo: ni una bb :(\n\n\tPonte las pilas pa consumir más 🥵\n\t")
}
