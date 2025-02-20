package read

import (
	"activity-tracker/shared"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSendTrackRead(t *testing.T) {
	c := require.New(t)

	ctx := context.Background()

	client, err := shared.NewMockBot("dummy")
	c.NoError(err)

	err = SendTrackRead(ctx, client, "test", "", 1)
	c.NoError(err)
}
