package register

import (
	"activity-tracker/database"
	"activity-tracker/shared"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRegisterUser(t *testing.T) {
	c := require.New(t)

	client, err := shared.NewMockBot("dummy")
	c.NoError(err)

	database.InitMongoMock()

	ctx := context.Background()

	user := shared.User{
		Name:   "testuser",
		ChatID: 12345,
	}

	err = RegisterUser(ctx, client, user.Name, user.ChatID)
	c.NoError(err)
}
