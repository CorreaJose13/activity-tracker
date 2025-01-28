package storage

import (
	"activity-tracker/database"
	"activity-tracker/shared"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMoc(t *testing.T) {
	c := require.New(t)

	database.InitMongoMock()

	user := shared.PersonalGoal{
		Username: "br",
		Activity: shared.Cycling,
	}

	err := CreatePersonalGoal(user)
	c.Nil(err)
}
