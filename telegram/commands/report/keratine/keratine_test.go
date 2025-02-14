package keratine

import (
	"activity-tracker/database"
	"activity-tracker/shared"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKeratineReport(t *testing.T) {
	c := require.New(t)

	client, err := shared.NewMockBot("dummy")
	c.Nil(err)

	ctx := context.Background()
	database.InitMongoMock()

	rp, err := GenerateKeratineReport(ctx, client, "BrayanEscobar", 1)
	c.NoError(err)
	c.NotEmpty(rp)
}
