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

const (
	userActivityTableName = "user-activity"
	activitiesTableName   = "activities"
)

var (

	// ErrNoActivitiesFound is returned when no activities are found
	ErrNoActivitiesFound = errors.New("no activities found")
)

// Create an user activity in database
func Create(ctx context.Context, userActivity shared.UserActivity) error {
	database.InitMongo()

	defer func() {
		if err := database.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	collection := database.GetCollection(userActivityTableName)

	_, err := collection.InsertOne(ctx, userActivity)

	return err
}

// Update an user activity in database
func UpdateContent(ctx context.Context, userActivity shared.UserActivity) error {
	database.InitMongo()

	defer func() {
		if err := database.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	collection := database.GetCollection(userActivityTableName)

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
		ctx,
		filter,
		update,
	)

	return err
}

// GetCurrentDayActivities returns the current day activities from inputs
func GetCurrentDayActivities(ctx context.Context, name string, activity shared.Activity) ([]*shared.UserActivity, error) {
	database.InitMongo()

	defer func() {
		if err := database.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	collection := database.GetCollection(userActivityTableName)

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

	items, err := collection.Find(ctx, filter)
	if err != nil {
		panic(err)
	}

	defer items.Close(ctx)

	var activities []*shared.UserActivity

	for items.TryNext(ctx) {
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
func GetLastWeekUserHistoryPerActivity(ctx context.Context, name string, activity shared.Activity) ([]*shared.UserActivity, error) {
	database.InitMongo()

	defer func() {
		if err := database.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	collection := database.GetCollection(userActivityTableName)

	now, err := shared.GetNow()
	if err != nil {
		return nil, err
	}

	nowStr := now.Format(time.RFC3339)

	// Calculate the number of days since Monday.
	daysSinceMonday := (int(now.Weekday()) + 6) % 7

	startDate := now.AddDate(0, 0, -daysSinceMonday)
	startDate = time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, startDate.Location())

	startDateStr := startDate.Format(time.RFC3339)

	filter := bson.M{
		"name":     name,
		"activity": activity,
		"created_at": bson.M{
			"$gte": startDateStr,
			"$lt":  nowStr,
		},
	}

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
			return nil, fmt.Errorf("decode bson failed: %w", err)
		}

		var activity shared.UserActivity

		bsBytes, _ := bson.Marshal(bs)

		err = bson.Unmarshal(bsBytes, &activity)
		if err != nil {
			return nil, fmt.Errorf("decode activity failed: %w", err)
		}

		activities = append(activities, &activity)
	}

	return activities, nil
}

// GetCurrentMonthUserHistoryPerActivity returns the current month activities by username and activity
func GetCurrentMonthUserHistoryPerActivity(ctx context.Context, name string, activity shared.Activity) ([]*shared.UserActivity, error) {
	database.InitMongo()

	defer func() {
		if err := database.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	collection := database.GetCollection(userActivityTableName)

	now, err := shared.GetNow()
	if err != nil {
		return nil, err
	}

	nowStr := now.Format(time.RFC3339)

	currentMonth := now.Month()
	startDate := time.Date(now.Year(), currentMonth, 1, 0, 0, 0, 0, now.Location())
	startDateStr := startDate.Format(time.RFC3339)

	filter := bson.M{
		"name":     name,
		"activity": activity,
		"created_at": bson.M{
			"$gte": startDateStr,
			"$lt":  nowStr,
		},
	}

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
			return nil, fmt.Errorf("decode bson failed: %w", err)
		}

		var activity shared.UserActivity

		bsBytes, _ := bson.Marshal(bs)

		err = bson.Unmarshal(bsBytes, &activity)
		if err != nil {
			return nil, fmt.Errorf("decode activity failed: %w", err)
		}

		activities = append(activities, &activity)
	}

	return activities, nil
}

// GetActivityHistory returns the activities by username and activity
func GetActivityHistory(ctx context.Context, name string, activity shared.Activity) ([]*shared.UserActivity, error) {
	database.InitMongo()

	defer func() {
		if err := database.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	collection := database.GetCollection(userActivityTableName)

	filter := bson.M{}

	filter["name"] = name
	filter["activity"] = activity

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

// CreateActivity creates a new activity
func CreateActivity(ctx context.Context, activity shared.Activity) error {
	database.InitMongo()

	defer func() {
		if err := database.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	collection := database.GetCollection(activitiesTableName)

	_, err := collection.InsertOne(ctx, activity)

	return err
}

// GetAvailableActivities returns the available activities
func GetAvailableActivities(ctx context.Context) ([]*shared.Activity, error) {
	database.InitMongo()

	defer func() {
		if err := database.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	collection := database.GetCollection(activitiesTableName)

	items, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	defer items.Close(ctx)

	var activities []*shared.Activity

	for items.Next(ctx) {
		var bs bson.M

		err := items.Decode(&bs)
		if err != nil {
			return nil, fmt.Errorf("decode bson failed: %w", err)
		}

		activityValue, ok := bs["activity"]
		if ok {
			activity := shared.Activity(activityValue.(string))
			activities = append(activities, &activity)
		}
	}

	return activities, nil
}
