package keratine

import (
	"activity-tracker/database"
	"activity-tracker/shared"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSendTrackKeratine(t *testing.T) {
	c := require.New(t)

	client, err := shared.NewMockBot("dummy")
	c.Nil(err)

	ctx := context.Background()

	database.InitMongoMock()

	err = SendTrackKeratine(ctx, client, "test1", "", 1)
	c.NoError(err)
}
