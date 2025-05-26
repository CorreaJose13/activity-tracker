package reports

import (
	"activity-tracker/database"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestWaterReport(t *testing.T) {
	c := require.New(t)

	database.InitMongoMock()

	report, err := NewWaterReport("BrayanEscobar")
	c.NoError(err)

	now := time.Now()
	monday := now.AddDate(0, 0, -int(now.Weekday())+1)

	report.SetDailyReports([]DailyReport{
		{
			Date:   monday,
			Result: float64(3),
		},
	})

	rp := report.String()
	c.Equal(rp, "Veamos cuánta agua has tomado esta semana BrayanEscobar 💧\n\n\tSegún tus registros, tus vasos de agua fueron:\n\n\tLunes: 3 vasos 💧\n\tMartes: ni un vaso, deshidratado 💀\n\tMiércoles: ni un vaso, deshidratado 💀\n\tJueves: ni un vaso, deshidratado 💀\n\tViernes: ni un vaso, deshidratado 💀\n\tSábado: ni un vaso, deshidratado 💀\n\tDomingo: ni un vaso, deshidratado 💀\n\n\tNo te deshidrates, toma más agua 💦\n\t")
}
