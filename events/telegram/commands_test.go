package telegram

import (
	"activity-tracker/api/telegram"
	"testing"

	"github.com/stretchr/testify/require"
)

// TODO: Add more unit tests for other commands

func TestSendTrackPipi(t *testing.T) {
	c := require.New(t)

	err := sendTrackPipi(&telegram.Bot{}, "test", "", 1)

	// It's not necessary to test the send message error
	c.Error(err)

	// It is set to fail to check the trace code in commands.go
	c.Equal("a", "b")
}

func TestSendTrackKeratine(t *testing.T) {
	c := require.New(t)

	err := sendTrackKeratine(&telegram.Bot{}, "test1", "", 1)

	// It's not necessary to test the send message error
	c.Error(err)

	// It is set to fail to check the trace code in commands.go
	c.Equal("a", "b")
}
