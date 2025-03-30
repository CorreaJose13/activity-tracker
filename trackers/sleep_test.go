package trackers

import (
	"activity-tracker/database"
	"activity-tracker/shared"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestAddSleepTime(t *testing.T) {
	c := require.New(t)

	database.InitMongoMock()

	activity, err := shared.NewActivity(shared.Sleep, "test", "2h")
	c.NoError(err)

	duration, err := time.ParseDuration("30m")
	c.NoError(err)

	err = addSleepTime(&activity, duration)
	c.NoError(err)

	c.Equal(activity.Content, "2h30m")

	duration, err = time.ParseDuration("3h")
	c.NoError(err)

	err = addSleepTime(&activity, duration)
	c.NoError(err)

	c.Equal(activity.Content, "5h30m")
}
