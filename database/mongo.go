package database

import (
	"context"
	"github.com/charmbracelet/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoConnection is a MongoDB connection.
type MongoConnection struct {
	client     *mongo.Client
	Collection *mongo.Collection
}

// NewMongoClient creates a new MongoDB client.
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

// NewMongoConnection creates a new MongoDB connection.
func NewMongoConnection(client *mongo.Client, database string, collection string) *MongoConnection {
	coll := client.Database(database).Collection(collection)
	return &MongoConnection{
		client:     client,
		Collection: coll,
	}
}
