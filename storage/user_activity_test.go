package storage

import (
	"activity-tracker/database"
	"activity-tracker/shared"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	c := require.New(t)

	now, err := shared.GetNow()
	c.NoError(err)

	nowStr := now.Format(time.RFC3339)

	database.InitMongoMock()

	user := shared.UserActivity{
		ID:        "123",
		Name:      "br",
		Activity:  "shit",
		CreatedAt: nowStr,
	}

	err = Create(user)
	c.NoError(err)
}
