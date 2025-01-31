package shared

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGenerateActivityItemID(t *testing.T) {
	c := require.New(t)

	now := time.Now()
	id := GenerateActivityItemID(now, "test", Water)

	c.Equal(id, fmt.Sprintf("%s-test-water", now.Format(time.RFC3339)))
}

func TestGetNow(t *testing.T) {
	c := require.New(t)

	now, err := GetNow()
	c.NoError(err)
	c.NotNil(now)
}

func TestIsValidContent(t *testing.T) {
	c := require.New(t)

	content := "a"
	b := IsValidNumber(content)
	c.False(b)

	content = "0"
	b = IsValidNumber(content)
	c.False(b)

	content = "10"
	b = IsValidNumber(content)
	c.True(b)
}
