package storage

import (
	"activity-tracker/database"
	"activity-tracker/shared"
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

const tableName = "user-activity"

var (
	collection = database.GetCollection(tableName)

	// ErrNoActivitiesFound is returned when no activities are found
	ErrNoActivitiesFound = errors.New("no activities found")
)

// Create an user activity in database
func Create(userActivity shared.UserActivity) error {
	_, err := collection.InsertOne(context.Background(), userActivity)

	return err
}

// Update an user activity in database
func UpdateContent(userActivity shared.UserActivity) error {
	filter := bson.M{
		"id": userActivity.ID,
	}

	update := bson.M{
		"$set": bson.M{
			"content":    userActivity.Content,
			"updated_at": userActivity.UpdatedAt,
		},
	}

	_, err := collection.UpdateOne(
		context.Background(),
		filter,
		update,
	)

	return err
}

// GetCurrentDayActivities returns the current day activities from inputs
func GetCurrentDayActivities(name string, activity shared.Activity) ([]*shared.UserActivity, error) {
	now, err := shared.GetNow()
	if err != nil {
		return nil, err
	}

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

	var activities []*shared.UserActivity

	for items.Next(ctx) {
		var bs bson.M

		err := items.Decode(&bs)
		if err != nil {
			return nil, fmt.Errorf("decode bson failed")
		}

		var activity shared.UserActivity

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
func GetLastWeekUserHistoryPerActivity(name string, activity shared.Activity) ([]*shared.UserActivity, error) {
	now, err := shared.GetNow()
	if err != nil {
		return nil, err
	}

	nowStr := now.Format(time.RFC3339)

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

	var activities []*shared.UserActivity

	for items.Next(ctx) {
		var bs bson.M

		err := items.Decode(&bs)
		if err != nil {
			return nil, fmt.Errorf("decode bson failed")
		}

		var activity shared.UserActivity

		bsBytes, _ := bson.Marshal(bs)

		err = bson.Unmarshal(bsBytes, &activity)
		if err != nil {
			return nil, fmt.Errorf("decode activity failed")
		}

		activities = append(activities, &activity)
	}

	return activities, nil
}

// GetActivityHistory returns the activities by username and activity
func GetActivityHistory(name string, activity shared.Activity) ([]*shared.UserActivity, error) {
	filter := bson.M{}

	filter["name"] = name
	filter["activity"] = activity

	ctx := context.Background()

	items, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	defer items.Close(ctx)

	var activities []*shared.UserActivity

	for items.Next(ctx) {
		var bs bson.M

		err := items.Decode(&bs)
		if err != nil {
			return nil, fmt.Errorf("decode bson failed")
		}

		var activity shared.UserActivity

		bsBytes, _ := bson.Marshal(bs)

		err = bson.Unmarshal(bsBytes, &activity)
		if err != nil {
			return nil, fmt.Errorf("decode activity failed")
		}

		activities = append(activities, &activity)
	}

	if len(activities) == 0 {
		return nil, ErrNoActivitiesFound
	}

	return activities, nil
}
