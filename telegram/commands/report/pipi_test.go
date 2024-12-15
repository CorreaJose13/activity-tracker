package report

import (
	"activity-tracker/shared"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPipiReport(t *testing.T) {
	c := require.New(t)

	pr, err := GeneratePipiReport(&shared.Bot{}, "test", 1)
	c.NoError(err)

	// It is set to fail to check the trace code in pipi.go
	c.Equal("testToFail", pr)
}
