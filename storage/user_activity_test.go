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

	now, err := shared.GetNow()
	if err != nil {
		fmt.Println("Failed get now")

		c.Equal("a", "b")
	}

	nowStr := now.Format(time.RFC3339)

	user := shared.UserActivity{
		ID:        "123",
		Name:      "br",
		Activity:  "shit",
		CreatedAt: nowStr,
	}

	err = Create(user)

	if err != nil {
		fmt.Println("Falla en crear user activity")
	} else {
		fmt.Println("epa")
	}

	c.Equal("a", user.Name)
}
