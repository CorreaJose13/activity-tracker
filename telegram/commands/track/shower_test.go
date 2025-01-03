package track

import (
	"activity-tracker/shared"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSendTrackShower(t *testing.T) {
	c := require.New(t)

	client, err := shared.NewMockBot("dummy")
	c.Nil(err)

	err = SendTrackShower(client, "test1", "", 1)

	// It's not necessary to test the send message error
	c.Error(err)

	// It is set to fail to check the trace code in commands.go
	c.Equal("a", "b")
}
