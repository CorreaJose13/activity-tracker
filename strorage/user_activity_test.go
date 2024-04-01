package user

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	c := require.New(t)

	user := UserActivity{
		ID:        "123",
		Name:      "br",
		Activity:  "shit",
		CreatedAt: time.Now(),
	}

	err := Create(user)

	if err != nil {
		fmt.Println("Falla en crear user activity")
	} else {
		fmt.Println("epa")
	}

	c.Equal("a", user.Name)
}
