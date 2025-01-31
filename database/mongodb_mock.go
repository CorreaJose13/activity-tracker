package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// To mock cursor items, cursor are the items used in reports, its like the bulk result of pagination
var findItems = []bson.M{}

type MongoCollectionInterface interface {
	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (cur *mongo.Cursor, err error)
}

type MongoDatabaseInterface interface {
	Collection(name string, opts ...*options.CollectionOptions) *mongo.Collection
}

type MongoClientInterface interface {
	Ping(ctx context.Context, rp *readpref.ReadPref) error
	Database(name string, opts ...*options.DatabaseOptions) *mongo.Database
}

type MongoCollection struct {
	Collection MongoCollectionInterface
}

type MongoDatase struct {
	Database MongoDatabaseInterface
}

type MongoClient struct {
	Client MongoClientInterface
}

type clientMock struct{}
type databaseMock struct{}
type collectionMock struct{}

func AddFindItem(item []bson.M) {
	findItems = append(findItems, item...)
}

func ClearFindItems() {
	findItems = []bson.M{}
}

func AddDummyFindItem() {
	var dummyData = []bson.M{
		{
			"_id":        "64eafc1234567890abcd1234",
			"username":   "johanFlorez",
			"created_at": "2025-01-01T18:13:25-05:00",
			"activity":   "gym",
		},
	}

	findItems = append(findItems, dummyData...)
}

func InitMongoMock() {
	Client = &clientMock{}
	Database = &databaseMock{}
	Collection = &collectionMock{}

	funcClientConfig = func() MongoClientInterface {
		return Client
	}

	funcGetDatabase = func() MongoDatabaseInterface {
		return Database
	}

	funcGetCollection = func(_ MongoDatabaseInterface, _ string) MongoCollectionInterface {
		return Collection
	}
}

func (c *clientMock) Database(name string, opts ...*options.DatabaseOptions) *mongo.Database {
	return nil
}

func (c *clientMock) Ping(ctx context.Context, rp *readpref.ReadPref) error {
	return nil
}

func (d *databaseMock) Collection(name string, opts ...*options.CollectionOptions) *mongo.Collection {
	return nil
}

func (c *collectionMock) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return &mongo.InsertOneResult{}, nil
}

func (c *collectionMock) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return &mongo.UpdateResult{}, nil
}

func (c *collectionMock) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return &mongo.DeleteResult{}, nil
}

func (c *collectionMock) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	documents := make([]interface{}, len(findItems))
	for i, doc := range findItems {
		documents[i] = doc
	}

	return mongo.NewCursorFromDocuments(documents, nil, nil)
}
