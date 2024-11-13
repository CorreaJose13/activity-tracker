package reports

import (
	"activity-tracker/api/telegram"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWaterReport(t *testing.T) {
	c := require.New(t)

	pr, err := GenerateWaterReport(&telegram.Bot{}, "test", 1)
	c.NoError(err)

	// It is set to fail to check the trace code in pipi.go
	c.Equal("testToFail", pr)
}
