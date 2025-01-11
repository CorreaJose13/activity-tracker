package wishlist

import (
	"activity-tracker/shared"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandleWishlist(t *testing.T) {
	c := require.New(t)

	client, err := shared.NewMockBot("dummy")
	c.Nil(err)

	err = HandleWishlist(client, 1, "test1", "test1")
	c.Equal(errInvalidWishlistCommand, err)

	err = HandleWishlist(client, 1, "test1", "test1 pepe")
	c.Equal(errInvalidURL, err)

	err = HandleWishlist(client, 1, "test1", "item https://www.google.com")

	// It's not necessary to test the send message error, at this point the item is added to the database
	c.Error(err)
}

func TestGetWishlist(t *testing.T) {
	c := require.New(t)

	client, err := shared.NewMockBot("dummy")
	c.Nil(err)

	err = GetWishlist(client, "test1", 1)

	// It's not necessary to test the send message error, at this point the wishlist is obtained from the database
	c.Error(err)
}
