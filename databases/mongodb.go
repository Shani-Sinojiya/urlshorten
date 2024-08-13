package databases

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"urlshorten.kwikbill.in/constants"
)

// MongoDBClient is the MongoDB client
var MongoDBClient *mongo.Client

// MongoDBDatabase is the MongoDB database
var MongoDBDatabase *mongo.Database

// ConnectMongoDB connects to MongoDB
func ConnectMongoDB() error {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(constants.DATABASE_MONGODB_URI))
	if err != nil {
		return err
	}

	MongoDBClient = client
	MongoDBDatabase = client.Database(constants.DATABASE_MONGODB_DB_NAME)

	return nil
}

// DisconnectMongoDB disconnects from MongoDB
func DisconnectMongoDB() error {
	if err := MongoDBClient.Disconnect(context.Background()); err != nil {
		return err
	}

	return nil
}
