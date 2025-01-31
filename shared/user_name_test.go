package shared

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetRandomUserName(t *testing.T) {
	c := require.New(t)

	c.Contains([]string{Valery, Brayan, Johan, Jose, Juan, Daniela}, GetRandomUserName())
}
