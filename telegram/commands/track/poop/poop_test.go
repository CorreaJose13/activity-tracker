package poop

import (
	"activity-tracker/shared"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSendTrackPoop(t *testing.T) {
	c := require.New(t)

	client, err := shared.NewMockBot("dummy")
	c.NoError(err)

	err = SendTrackPoop(client, "test", "", 1)
	c.NoError(err)
}
