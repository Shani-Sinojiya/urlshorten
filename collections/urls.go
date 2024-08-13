package collections

import (
	"go.mongodb.org/mongo-driver/mongo"
	"urlshorten.kwikbill.in/databases"
)

func Urls() *mongo.Collection {
	return databases.MongoDBDatabase.Collection("urls")
}
