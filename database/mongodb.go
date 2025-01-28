package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client            MongoClientInterface     = &mongo.Client{}
	Database          MongoDatabaseInterface   = &mongo.Database{}
	Collection        MongoCollectionInterface = &mongo.Collection{}
	funcClientConfig                           = setClientConfig
	funcGetDatabase                            = getDatabase
	funcGetCollection                          = getCollection
)

const (
	user         = "activitytracker"
	databaseName = "activitytracker"
)

func setClientConfig() (context.Context, MongoClientInterface) {
	pswd := os.Getenv("MONGO_TOKEN")
	if pswd == "" {
		panic("failed to get mongo token env var value")
	}

	clientOpts := options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@%s.2ykonih.mongodb.net/", user, pswd, databaseName))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongo, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		panic(err)
	}

	return ctx, mongo
}

func InitMongo() {
	ctx, cli := funcClientConfig()
	Client = cli

	err := Client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}
}

func getDatabase() MongoDatabaseInterface {
	return Client.Database(databaseName)
}

func getCollection(database MongoDatabaseInterface, tableName string) MongoCollectionInterface {
	return Database.Collection(tableName)
}

// GetCollection function to get the mogodb collection
func GetCollection(tableName string) MongoCollectionInterface {
	Database = funcGetDatabase()
	return funcGetCollection(Database, tableName)
}
