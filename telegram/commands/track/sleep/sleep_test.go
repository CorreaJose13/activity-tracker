package sleep

import (
	"activity-tracker/database"
	"activity-tracker/shared"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestSendTrackSleep(t *testing.T) {
	c := require.New(t)

	client, err := shared.NewMockBot("dummy")
	c.NoError(err)

	database.InitMongoMock()

	ctx := context.Background()

	err = SendTrackSleep(ctx, client, "test", "", 1)
	c.NoError(err)
}

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
