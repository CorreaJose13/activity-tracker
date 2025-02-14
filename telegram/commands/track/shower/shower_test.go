package shower

import (
	"activity-tracker/database"
	"activity-tracker/shared"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSendTrackShower(t *testing.T) {
	c := require.New(t)

	client, err := shared.NewMockBot("dummy")
	c.Nil(err)

	database.InitMongoMock()

	ctx := context.Background()

	err = SendTrackShower(ctx, client, "test1", "", 1)
	c.NoError(err)
}
