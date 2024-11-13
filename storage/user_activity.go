package storage

import (
	"activity-tracker/database"
	"activity-tracker/shared"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

const tableName = "user-activity"

var (
	collection = database.GetCollection(tableName)
)

// UserActivity contains an user activity info
type UserActivity struct {
	ID           string          `bson:"id"`
	Name         string          `bson:"name"`
	Activity     shared.Activity `bson:"activity"`
	ExerciseType shared.Exercise `bson:"excercise_type,omitempty"`
	CreatedAt    string          `bson:"created_at"`
	Content      string          `bson:"content,omitempty"`
}

// Create an user activity in database
func Create(userActivity UserActivity) error {
	_, err := collection.InsertOne(context.Background(), userActivity)

	return err
}

// GenerateActivityItemID generate the unique id of the activity item that will be saved in the activity database
func GenerateActivityItemID(now time.Time, username string, activity shared.Activity) string {
	formattedNow := now.Format(time.RFC3339)

	return fmt.Sprintf("%s-%s-%s", formattedNow, username, activity)
}

// GetCurrentDayActivities returns the current day activities from inputs
func GetCurrentDayActivities(name string, activity shared.Activity) ([]*UserActivity, error) {
	now := shared.GetNow()

	startDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endDay := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())

	startDayStr := startDay.Format(time.RFC3339)
	endDayStr := endDay.Format(time.RFC3339)

	filter := bson.M{}

	filter["name"] = name
	filter["activity"] = activity
	filter["created_at"] = bson.M{
		"$gte": startDayStr,
		"$lt":  endDayStr,
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
func GetLastWeekUserHistoryPerActivity(name string, activity shared.Activity) ([]*UserActivity, error) {
	now := shared.GetNow()
	nowStr := shared.GetNow().Format(time.RFC3339)

	// Calculate the days until the last monday
	daysUntilLastMonday := int(now.Weekday())
	if daysUntilLastMonday == 0 { // it's sunday
		daysUntilLastMonday = 7
	}

	// Go back to the last monday at 00:00:00
	startDate := now.AddDate(0, 0, -daysUntilLastMonday)
	startDate = time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, startDate.Location())

	startDateStr := startDate.Format(time.RFC3339)

	filter := bson.M{}

	filter["name"] = name
	filter["activity"] = activity
	filter["created_at"] = bson.M{
		"$gte": startDateStr,
		"$lt":  nowStr,
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
