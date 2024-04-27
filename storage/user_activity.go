package storage

import (
	"activity-tracker/database"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

const tableName = "user-activity"

var (
	Leg    Exercise = "leg"
	Bicep  Exercise = "bicep"
	Back   Exercise = "back"
	Tricep Exercise = "tricep"
	Abs    Exercise = "abs"
	Cardio Exercise = "cardio"
	Chest  Exercise = "chest"

	collection = database.GetCollection(tableName)
)

type Exercise string

// UserActivity contains an user activity info
type UserActivity struct {
	ID           string    `bson:"id"`
	Name         string    `bson:"name"`
	Activity     string    `bson:"activity"`
	CreatedAt    time.Time `bson:"created_at"`
	ExerciseType Exercise  `bson:"excercise_type,omitempty"`
}

// Create an user activity in database
func Create(userActivity UserActivity) error {
	_, err := collection.InsertOne(context.Background(), userActivity)

	return err
}

// GetCurrentDayActivities returns the current day activities from inputs
func GetCurrentDayActivities(name, activity string) ([]*UserActivity, error) {
	now := time.Now()

	startDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endDay := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())

	filter := bson.M{}

	filter["name"] = name
	filter["activity"] = activity
	filter["created_at"] = bson.M{
		"$gte": startDay,
		"$lt":  endDay,
	}

	ctx := context.Background()

	items, err := collection.Find(ctx, filter)
	if err != nil {
		panic(err)
	}

	defer items.Close(ctx)

	var activities []*UserActivity

	for items.Next(ctx) {
		var bs bson.M

		err := items.Decode(&bs)
		if err != nil {
			return nil, fmt.Errorf("decode bson failed")
		}

		var activity UserActivity

		bsBytes, _ := bson.Marshal(bs)

		err = bson.Unmarshal(bsBytes, &activity)
		if err != nil {
			return nil, fmt.Errorf("decode activity failed")
		}

		activities = append(activities, &activity)
	}

	return activities, nil
}

// GetLastWeekUserHistoryPerActivity returns the last week activities by username and activity
func GetLastWeekUserHistoryPerActivity(name, activity string) ([]*UserActivity, error) {
	// We assume that only runs on sundays when the water scheduler is implemented
	now := time.Now()
	daysToMonday := 1 - int(now.Weekday())
	monday := now.AddDate(0, 0, daysToMonday)

	filter := bson.M{}

	filter["name"] = name
	filter["activity"] = activity
	filter["created_at"] = bson.M{
		"$gte": monday,
		"$lt":  now,
	}

	ctx := context.Background()

	items, err := collection.Find(ctx, filter)
	if err != nil {
		panic(err)
	}

	defer items.Close(ctx)

	var activities []*UserActivity

	for items.Next(ctx) {
		var bs bson.M

		err := items.Decode(&bs)
		if err != nil {
			return nil, fmt.Errorf("decode bson failed")
		}

		var activity UserActivity

		bsBytes, _ := bson.Marshal(bs)

		err = bson.Unmarshal(bsBytes, &activity)
		if err != nil {
			return nil, fmt.Errorf("decode activity failed")
		}

		activities = append(activities, &activity)
	}

	return activities, nil
}
