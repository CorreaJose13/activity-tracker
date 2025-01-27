package report

import (
	"activity-tracker/database"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSleepReport(t *testing.T) {
	c := require.New(t)

	database.InitMongoMock()

	_, err := generateSleepReport("test")
	c.NoError(err)
}
