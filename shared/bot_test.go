package shared

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBotSend(t *testing.T) {
	c := require.New(t)

	client, err := NewMockBot("dummy-token")
	c.Nil(err)

	err = client.SendMessage(123, "dummy-text")
	c.Nil(err)

	err = client.SendPhoto(123, "https://www.dummyurl.com")
	c.Nil(err)
}

func TestBotForcedFailure(t *testing.T) {
	c := require.New(t)

	client, err := NewMockBot("dummy-token")
	c.Nil(err)

	ForceMockFailure = true
	defer func() { ForceMockFailure = false }()

	err = client.SendPhoto(123, "https://www.dummyurl.com")
	c.ErrorIs(err, ErrForcedFailure)

	err = client.SendMessage(123, "dummy-text")
	c.ErrorIs(err, ErrForcedFailure)
}
