package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

const (
	user         = "activitytracker"
	pswd         = "activitytracker"
	databaseName = "activitytracker"
)

func init() {
	clientOpts := options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@%s.2ykonih.mongodb.net/", user, pswd, databaseName))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongo, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatal(err)
	}

	mongoClient = mongo

	err = mongoClient.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// GetCollection function to get the mogodb collection
func GetCollection(tableName string) *mongo.Collection {
	return mongoClient.Database(databaseName).Collection(tableName)
}
