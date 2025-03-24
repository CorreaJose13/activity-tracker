package track

import (
	"activity-tracker/database"
	"activity-tracker/shared"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSendTrackGomita(t *testing.T) {
	c := require.New(t)

	client, err := shared.NewMockBot("dummy")
	c.Nil(err)

	database.InitMongoMock()

	ctx := context.Background()

	err = SendTrackActivity(ctx, shared.Gomita, client, "test", "", 1)
	c.NoError(err)
}

func TestSendTrackKeratine(t *testing.T) {
	c := require.New(t)

	client, err := shared.NewMockBot("dummy")
	c.Nil(err)

	ctx := context.Background()

	database.InitMongoMock()

	err = SendTrackActivity(ctx, shared.Keratine, client, "test1", "", 1)
	c.NoError(err)
}

func TestSendTrackPipi(t *testing.T) {
	c := require.New(t)

	client, err := shared.NewMockBot("dummy")
	c.Nil(err)

	ctx := context.Background()

	database.InitMongoMock()

	err = SendTrackActivity(ctx, shared.Pipi, client, "test", "", 1)
	c.NoError(err)
}

func TestSendTrackPoop(t *testing.T) {
	c := require.New(t)

	ctx := context.Background()

	client, err := shared.NewMockBot("dummy")
	c.NoError(err)

	err = SendTrackActivity(ctx, shared.Poop, client, "test", "", 1)
	c.NoError(err)
}

func TestSendTrackRead(t *testing.T) {
	c := require.New(t)

	ctx := context.Background()

	client, err := shared.NewMockBot("dummy")
	c.NoError(err)

	err = SendTrackActivity(ctx, shared.Read, client, "test", "", 1)
	c.NoError(err)
}

func TestSendTrackShower(t *testing.T) {
	c := require.New(t)

	client, err := shared.NewMockBot("dummy")
	c.Nil(err)

	database.InitMongoMock()

	ctx := context.Background()

	err = SendTrackActivity(ctx, shared.Shower, client, "test1", "", 1)
	c.NoError(err)
}

func TestSendTrackSleep(t *testing.T) {
	c := require.New(t)

	client, err := shared.NewMockBot("dummy")
	c.NoError(err)

	database.InitMongoMock()

	ctx := context.Background()

	err = SendTrackActivity(ctx, shared.Sleep, client, "test", "", 1)
	c.NoError(err)
}
