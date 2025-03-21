package pipi

import (
	"activity-tracker/database"
	"activity-tracker/shared"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSendTrackPipi(t *testing.T) {
	c := require.New(t)

	client, err := shared.NewMockBot("dummy")
	c.Nil(err)

	ctx := context.Background()

	database.InitMongoMock()

	err = SendTrackPipi(ctx, client, "test", "", 1)
	c.NoError(err)
}
