package shared

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetRandomUserName(t *testing.T) {
	c := require.New(t)

	c.Contains([]string{Valery, Brayan, Johan, Jose, Juan, Daniela}, GetRandomUserName())
}

func TestNewUser(t *testing.T) {
	c := require.New(t)

	user := NewUser("test", 123, []Activity{Shower})

	c.Equal("test", user.Name)
	c.Equal(int64(123), user.ChatID)
	c.Equal([]Activity{Shower}, user.EnabledActivites)
}
