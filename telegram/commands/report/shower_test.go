package report

import (
	"activity-tracker/shared"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShowerReport(t *testing.T) {
	c := require.New(t)

	rp, err := GenerateShowerReport(&shared.Client{}, "test", 1)
	c.NoError(err)
	c.NotEmpty(rp)
}
