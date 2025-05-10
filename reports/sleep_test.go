package reports

import (
	"activity-tracker/database"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestSleepReport(t *testing.T) {
	c := require.New(t)

	database.InitMongoMock()

	report, err := NewSleepReport("BrayanEscobar")
	c.NoError(err)

	now := time.Now()
	monday := now.AddDate(0, 0, -int(now.Weekday())+1)

	report.SetDailyReports([]DailyReport{
		{
			Date:   monday,
			Result: float64(8),
		},
	})

	rp := report.String()
	c.Equal(rp, "Veamos cómo has dormido esta semana BrayanEscobar 😴\n\n\tSegún tus registros, tus horas de sueño fueron:\n\n\tLunes: 8.0 horas 😴\n\tMartes: ni una hora, zombie 🧟‍♂️\n\tMiércoles: ni una hora, zombie 🧟‍♂️\n\tJueves: ni una hora, zombie 🧟‍♂️\n\tViernes: ni una hora, zombie 🧟‍♂️\n\tSábado: ni una hora, zombie 🧟‍♂️\n\tDomingo: ni una hora, zombie 🧟‍♂️\n\n\tDuerme más, que te ves cansado 😪\n\t")
}
