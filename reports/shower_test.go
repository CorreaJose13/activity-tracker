package reports

import (
	"activity-tracker/database"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestShowerReport(t *testing.T) {
	c := require.New(t)

	database.InitMongoMock()

	report, err := NewShowerReport("BrayanEscobar")
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
	c.Equal(rp, "Miremos cuántas veces te has bañado esta semana BrayanEscobar 🚿\n\n\tSegún la base de datos, tus duchas fueron así:\n\n\tLunes: 1 ducha 🚿\n\tMartes: ni una ducha, cochino 🤢\n\tMiércoles: ni una ducha, cochino 🤢\n\tJueves: ni una ducha, cochino 🤢\n\tViernes: ni una ducha, cochino 🤢\n\tSábado: ni una ducha, cochino 🤢\n\tDomingo: ni una ducha, cochino 🤢\n\n\tNo seas cochino, báñate más seguido 🧼\n\t")
}
