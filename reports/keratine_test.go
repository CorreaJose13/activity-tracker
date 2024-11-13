package reports

import (
	"activity-tracker/api/telegram"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKeratineReport(t *testing.T) {
	c := require.New(t)

	pr, err := GenerateKeratineReport(&telegram.Bot{}, "test", 1)
	c.NoError(err)

	// It is set to fail to check the trace code in pipi.go
	c.Equal("testToFail", pr)
}
