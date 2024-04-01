package user_activity

import (
	"activity-tracker/database"
	"context"
	"time"
)

const tableName = "users"

var (
	collection = database.GetCollection(tableName)
)

// UserActivity contains an user activity info
type UserActivity struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Activity  string    `json:"activity"`
	CreatedAt time.Time `json:"created_at"`
	Message   string    `json:"message,omitempty"`
}

// Create an user activity in database
func Create(userActivity UserActivity) error {
	_, err := collection.InsertOne(context.Background(), userActivity)

	return err
}
