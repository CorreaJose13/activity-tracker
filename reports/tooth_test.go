package reports

import (
	"activity-tracker/database"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestToothReport(t *testing.T) {
	c := require.New(t)

	database.InitMongoMock()

	report, err := NewToothReport("BrayanEscobar")
	c.NoError(err)

	now := time.Now()
	monday := now.AddDate(0, 0, -int(now.Weekday())+1)

	report.SetDailyReports([]DailyReport{
		{
			Date:   monday,
			Result: float64(2),
		},
	})

	rp := report.String()
	c.Equal(rp, "Miremos cuántas veces te has cepillado los dientes esta semana BrayanEscobar 🪥\n\n\tSegún la base de datos, tus cepilladas fueron:\n\n\tLunes: 2 cepilladas 🪥\n\tMartes: ni una cepillada, cochino 🦷\n\tMiércoles: ni una cepillada, cochino 🦷\n\tJueves: ni una cepillada, cochino 🦷\n\tViernes: ni una cepillada, cochino 🦷\n\tSábado: ni una cepillada, cochino 🦷\n\tDomingo: ni una cepillada, cochino 🦷\n\n\tNo querés que se te caigan los dientes, ¿o sí? 🦷\n\t")
}
