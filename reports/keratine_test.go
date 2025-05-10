package reports

import (
	"activity-tracker/database"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestKeratineReport(t *testing.T) {
	c := require.New(t)

	database.InitMongoMock()

	report, err := NewKeratineReport("BrayanEscobar")
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
	c.Equal(rp, "Pillá pues cómo son las vueltas precios@ BrayanEscobar 🍆\n\n\tEsta semana has tomado keratina así bb:\n\n\tLunes: sisas\n\tMartes: \n\tMiércoles: \n\tJueves: \n\tViernes: \n\tSábado: \n\tDomingo: \n\n\tSi querés mejorar estos números ponete las pilas con la keratina 😘\n\t")
}
