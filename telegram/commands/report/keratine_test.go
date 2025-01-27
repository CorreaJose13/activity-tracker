package report

import (
	"activity-tracker/database"
	"activity-tracker/shared"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKeratineReport(t *testing.T) {
	c := require.New(t)

	client, err := shared.NewMockBot("dummy")
	c.Nil(err)

	database.InitMongoMock()

	pr, err := GenerateKeratineReport(client, "test", 1)
	c.NoError(err)
	c.NotEmpty(pr)
}
