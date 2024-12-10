package report

import (
	"activity-tracker/shared"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShowerReport(t *testing.T) {
	c := require.New(t)

	a, b, e, d, err := generateShowerReport(&shared.Bot{}, "test", 1)
	c.NoError(err)
	c.Equal(0, 1)
	c.Equal(0, a)
	c.Equal(0, b)
	c.Equal(0, e)
	c.Equal(0, d)

}
