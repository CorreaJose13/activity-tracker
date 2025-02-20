package storage

import (
	"activity-tracker/database"
	"activity-tracker/shared"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

const usersTableName = "users"

// Create an user in database
func CreateUser(ctx context.Context, user shared.User) error {
	database.InitMongo()

	defer func() {
		if err := database.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	collection := database.GetCollection(usersTableName)

	_, err := collection.InsertOne(ctx, user)

	return err
}

// GetUser returns a user from the database
func GetUser(ctx context.Context, userName string) (shared.User, error) {
	database.InitMongo()

	defer func() {
		if err := database.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	collection := database.GetCollection(usersTableName)

	var user shared.User
	err := collection.FindOne(ctx, bson.M{"name": userName}).Decode(&user)
	return user, err
}

// UpdateUser updates a user in the database
func UpdateUser(ctx context.Context, user shared.User) error {
	database.InitMongo()

	defer func() {
		if err := database.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	collection := database.GetCollection(usersTableName)

	_, err := collection.UpdateOne(ctx, bson.M{"name": user.Name}, bson.M{"$set": user})
	return err
}
