package track

import (
	"activity-tracker/shared"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSendTrackRead(t *testing.T) {
	c := require.New(t)

	client, err := shared.NewMockBot("dummy")
	c.NoError(err)

	err = SendTrackRead(client, "test", "", 1)
	c.NoError(err)
}

func TestIsValidContent(t *testing.T) {
	c := require.New(t)

	content := "a"

	b := isValidContent(content)
	c.False(b)

	content = "10"

	b = isValidContent(content)
	c.True(b)
}
