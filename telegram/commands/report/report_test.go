package report

import (
	"activity-tracker/database"
	"activity-tracker/shared"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSendReportGomita(t *testing.T) {
	c := require.New(t)

	client, err := shared.NewMockBot("dummy")
	c.Nil(err)

	database.InitMongoMock()

	ctx := context.Background()

	err = SendReportActicvity(ctx, client, shared.Gomita, "test", "", 1)
	c.NoError(err)
}

func TestSendReportKeratine(t *testing.T) {
	c := require.New(t)

	client, err := shared.NewMockBot("dummy")
	c.Nil(err)

	database.InitMongoMock()

	ctx := context.Background()

	err = SendReportActicvity(ctx, client, shared.Keratine, "test", "", 1)
	c.NoError(err)
}

func TestSendReportPipi(t *testing.T) {
	c := require.New(t)

	client, err := shared.NewMockBot("dummy")
	c.Nil(err)

	database.InitMongoMock()

	ctx := context.Background()

	err = SendReportActicvity(ctx, client, shared.Pipi, "test", "", 1)
	c.NoError(err)
}

func TestSendReportRun(t *testing.T) {
	c := require.New(t)

	client, err := shared.NewMockBot("dummy")
	c.Nil(err)

	database.InitMongoMock()

	ctx := context.Background()

	err = SendReportActicvity(ctx, client, shared.Run, "test", "", 1)
	c.NoError(err)
}

func TestSendReportShower(t *testing.T) {
	c := require.New(t)

	client, err := shared.NewMockBot("dummy")
	c.Nil(err)

	database.InitMongoMock()

	ctx := context.Background()

	err = SendReportActicvity(ctx, client, shared.Shower, "test", "", 1)
	c.NoError(err)
}

func TestSendReportSleep(t *testing.T) {
	c := require.New(t)

	client, err := shared.NewMockBot("dummy")
	c.Nil(err)

	database.InitMongoMock()

	ctx := context.Background()

	err = SendReportActicvity(ctx, client, shared.Sleep, "test", "", 1)
	c.NoError(err)
}

func TestSendReportTooth(t *testing.T) {
	c := require.New(t)

	client, err := shared.NewMockBot("dummy")
	c.Nil(err)

	database.InitMongoMock()

	ctx := context.Background()

	err = SendReportActicvity(ctx, client, shared.ToothBrush, "test", "", 1)
	c.NoError(err)
}

func TestSendReportWater(t *testing.T) {
	c := require.New(t)

	client, err := shared.NewMockBot("dummy")
	c.Nil(err)

	database.InitMongoMock()

	ctx := context.Background()

	err = SendReportActicvity(ctx, client, shared.Water, "test", "", 1)
	c.NoError(err)
}
