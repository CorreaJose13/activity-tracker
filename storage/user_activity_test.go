package storage

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	c := require.New(t)

	location, err := time.LoadLocation("America/Bogota")
	if err != nil {
		t.Fatal(err)
	}

	now := time.Now().In(location)

	user := UserActivity{
		ID:        "123",
		Name:      "br",
		Activity:  "shit",
		CreatedAt: now,
	}

	err = Create(user)

	if err != nil {
		fmt.Println("Falla en crear user activity")
	} else {
		fmt.Println("epa")
	}

	c.Equal("a", user.Name)
}
