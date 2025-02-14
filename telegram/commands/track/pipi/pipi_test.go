package pipi

import (
	"activity-tracker/database"
	"activity-tracker/shared"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSendTrackPipi(t *testing.T) {
	c := require.New(t)

	client, err := shared.NewMockBot("dummy")
	c.Nil(err)

	database.InitMongoMock()

	err = SendTrackPipi(client, "test", "", 1)
	c.NoError(err)
}
