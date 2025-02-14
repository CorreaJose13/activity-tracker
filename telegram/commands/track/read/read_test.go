package read

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
