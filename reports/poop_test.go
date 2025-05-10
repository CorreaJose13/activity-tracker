package reports

import (
	"activity-tracker/database"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPoopReport(t *testing.T) {
	c := require.New(t)

	database.InitMongoMock()

	report, err := NewPoopReport("BrayanEscobar")
	c.NoError(err)

	now := time.Now()
	monday := now.AddDate(0, 0, -int(now.Weekday())+1)

	report.SetDailyReports([]DailyReport{
		{
			Date:   monday,
			Result: float64(1),
		},
	})

	rp := report.String()
	c.NoError(err)
	c.Equal(rp, "Mirá todas las veces que has cagado esta semana BrayanEscobar 🧐\n\n\tsegún ese intento de base de datos que tenemos, cagaste de la siguiente manera:\n\n\tLunes: 1 cagada 💩\n\tMartes: no cagaste bb 💩\n\tMiércoles: no cagaste bb 💩\n\tJueves: no cagaste bb 💩\n\tViernes: no cagaste bb 💩\n\tSábado: no cagaste bb 💩\n\tDomingo: no cagaste bb 💩\n\t")
}
