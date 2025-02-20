package wishlist

import (
	"activity-tracker/database"
	"activity-tracker/shared"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandleWishlist(t *testing.T) {
	c := require.New(t)

	client, err := shared.NewMockBot("dummy")
	c.Nil(err)

	database.InitMongoMock()

	ctx := context.Background()

	err = HandleWishlist(ctx, client, 1, "test1", "test1")
	c.Equal(errInvalidWishlistCommand, err)

	err = HandleWishlist(ctx, client, 1, "test1", "test1 pepe")
	c.Equal(errInvalidURL, err)

	err = HandleWishlist(ctx, client, 1, "test1", "item https://www.google.com")
	c.NoError(err)
}

func TestGetWishlist(t *testing.T) {
	c := require.New(t)

	database.InitMongoMock()

	ctx := context.Background()

	client, err := shared.NewMockBot("dummy")
	c.Nil(err)

	err = GetWishlist(ctx, client, "test1", 1)
	c.NoError(err)
}
