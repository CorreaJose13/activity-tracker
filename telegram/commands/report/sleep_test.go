package report

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSleepReport(t *testing.T) {
	c := require.New(t)

	_, err := generateSleepReport("test")
	c.NoError(err)
}
