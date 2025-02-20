package poop

import (
	"activity-tracker/database"
	"activity-tracker/shared"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPoopReport(t *testing.T) {
	c := require.New(t)

	client, err := shared.NewMockBot("dummy")
	c.Nil(err)

	database.InitMongoMock()

	ctx := context.Background()

	rp, err := GeneratePoopReport(ctx, client, "test", 1)
	c.NoError(err)
	c.NotEmpty(rp)
}
