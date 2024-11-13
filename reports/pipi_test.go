package reports

import (
	"activity-tracker/api/telegram"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPipiReport(t *testing.T) {
	c := require.New(t)

	pr, err := GeneratePipiReport(&telegram.Bot{}, "test", 1)
	c.NoError(err)

	// It is set to fail to check the trace code in pipi.go
	c.Equal("testToFail", pr)
}
