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

func setClientConfig() MongoClientInterface {
	pswd := os.Getenv("MONGO_TOKEN")
	if pswd == "" {
		panic("failed to get mongo token env var value")
	}

	clientOpts := options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@%s.2ykonih.mongodb.net/", user, pswd, databaseName))

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	mongo, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		panic(err)
	}

	err = mongo.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	return mongo
}

func InitMongo() {
	Client = funcClientConfig()
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
