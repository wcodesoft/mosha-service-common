package database

import (
	"context"
	"github.com/charmbracelet/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDatabase struct {
	client     *mongo.Client
	Collection *mongo.Collection
}

func NewMongoClient(mongoURI string) (*mongo.Client, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil, err
	}
	log.Infof("Connected to MongoDB: %s", mongoURI)
	return client, nil
}

func NewMongoDatabase(client *mongo.Client, database string, collection string) *MongoDatabase {
	coll := client.Database(database).Collection(collection)
	return &MongoDatabase{
		client:     client,
		Collection: coll,
	}
}
