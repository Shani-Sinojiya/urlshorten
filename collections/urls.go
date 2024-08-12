package collections

import (
	"go.mongodb.org/mongo-driver/mongo"
	"urlshorten.kwikbill.in/database"
)

func Urls() *mongo.Collection {
	return database.MongoDBDatabase.Collection("urls")
}
