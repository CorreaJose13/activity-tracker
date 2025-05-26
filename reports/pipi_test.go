package reports

import (
	"activity-tracker/database"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPipiReport(t *testing.T) {
	c := require.New(t)

	database.InitMongoMock()

	report, err := NewPipiReport("BrayanEscobar")
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
	c.Equal(rp, "Pillá pues cómo son las vueltas precios@ BrayanEscobar 🍆\n\n\tEsta semana has miado así bb:\n\n\tLunes: 1 miada 🍆\n\tMartes: no miaste bb 🍆\n\tMiércoles: no miaste bb 🍆\n\tJueves: no miaste bb 🍆\n\tViernes: no miaste bb 🍆\n\tSábado: no miaste bb 🍆\n\tDomingo: no miaste bb 🍆\n\n\tSi querés miar más ponete a tomar awa en vez de pensar en tu ex 😘\n\t")
}
