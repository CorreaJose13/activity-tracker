package reports

import (
	"activity-tracker/database"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestRunReport(t *testing.T) {
	c := require.New(t)

	database.InitMongoMock()

	report, err := NewRunReport("BrayanEscobar")
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
	c.Equal(rp, "Llegaron tus métricas de mierda subatleta BrayanEscobar 🏃🏾‍♂️\n\n\tLunes: 5.00 kms 🏃🏾‍♂️\n\tMartes: ni un metro, perezoso 🤢\n\tMiércoles: ni un metro, perezoso 🤢\n\tJueves: ni un metro, perezoso 🤢\n\tViernes: ni un metro, perezoso 🤢\n\tSábado: ni un metro, perezoso 🤢\n\tDomingo: ni un metro, perezoso 🤢\n\n\tSi no querés que te robe alguien de sucre comenzá a correr más 🤢\n\t")
}
