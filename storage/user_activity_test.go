package storage

import (
	"activity-tracker/shared"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	c := require.New(t)

	now := shared.GetNow().Format(time.RFC3339)

	user := UserActivity{
		ID:        "123",
		Name:      "br",
		Activity:  "shit",
		CreatedAt: now,
	}

	err := Create(user)

	if err != nil {
		fmt.Println("Falla en crear user activity")
	} else {
		fmt.Println("epa")
	}

	c.Equal("a", user.Name)
}
