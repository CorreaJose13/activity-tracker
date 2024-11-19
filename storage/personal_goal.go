package storage

import (
	"activity-tracker/database"
	"activity-tracker/shared"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

const personalGoaltableName = "personal-goal"

var (
	personalGoalCollection = database.GetCollection(personalGoaltableName)
)

// Create an user activity in database
func CreatePersonalGoal(personalGoal shared.PersonalGoal) error {
	_, err := personalGoalCollection.InsertOne(context.Background(), personalGoal)

	return err
}

// UpdatePersonalGoal updates the personal goal of an user
func UpdatePersonalGoal(personalGoal shared.PersonalGoal) error {
	filter := bson.M{}

	filter["username"] = personalGoal.Username
	filter["activity"] = personalGoal.Activity

	_, err := personalGoalCollection.UpdateOne(context.Background(), filter, bson.M{"$set": personalGoal})

	return err
}

// DeletePersonalGoal deletes the personal goal of an user
func DeletePersonalGoal(username string, activity shared.Activity) error {
	filter := bson.M{}

	filter["username"] = username
	filter["activity"] = activity

	_, err := personalGoalCollection.DeleteOne(context.Background(), filter)

	return err
}

// GetAllPersonalGoals returns all the personal goals of an user
func GetAllPersonalGoals(username string) ([]*shared.PersonalGoal, error) {
	filter := bson.M{}

	filter["username"] = username

	ctx := context.Background()

	items, err := personalGoalCollection.Find(ctx, filter)
	if err != nil {
		panic(err)
	}

	defer items.Close(ctx)

	var personalGoals []*shared.PersonalGoal

	for items.Next(ctx) {
		var bs bson.M

		err := items.Decode(&bs)
		if err != nil {
			return nil, fmt.Errorf("decode bson failed")
		}

		var personalGoal shared.PersonalGoal

		bsBytes, _ := bson.Marshal(bs)

		err = bson.Unmarshal(bsBytes, &personalGoal)
		if err != nil {
			return nil, fmt.Errorf("decode personal goal failed")
		}

		personalGoals = append(personalGoals, &personalGoal)
	}

	return personalGoals, nil
}
