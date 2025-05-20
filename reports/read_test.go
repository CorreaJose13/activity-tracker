package reports

import (
	"activity-tracker/database"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestReadReport(t *testing.T) {
	c := require.New(t)

	database.InitMongoMock()

	report, err := NewReadReport("BrayanEscobar")
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
	c.Equal(rp, "Miremos cómo te ha ido con la lectura esta semana BrayanEscobar 🥸\n\n\tsegún la esquizofrenia de la base de datos, por día leíste la siguiente cantidad de páginas:\n\n\tLunes: 1 perra página a lo bien??? 🤨\n\tMartes: sos un analfabeta\n\tMiércoles: sos un analfabeta\n\tJueves: sos un analfabeta\n\tViernes: sos un analfabeta\n\tSábado: sos un analfabeta\n\tDomingo: sos un analfabeta\n\n\tPa' que dejés de decir 'haiga' toca leer un toque más 🙄\n\t")
}
