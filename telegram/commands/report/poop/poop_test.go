package poop

import (
	"activity-tracker/database"
	"activity-tracker/shared"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPoopReport(t *testing.T) {
	c := require.New(t)

	client, err := shared.NewMockBot("dummy")
	c.Nil(err)

	database.InitMongoMock()

	rp, err := GeneratePoopReport(client, "test", 1)
	c.NoError(err)
	c.NotEmpty(rp)
}
