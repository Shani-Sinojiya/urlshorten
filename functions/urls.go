package functions

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"urlshorten.kwikbill.in/collections"
)

func CreateUrl(longurl, shorturl string) (bool, primitive.ObjectID, error) {
	info, err := collections.Urls().InsertOne(context.Background(), bson.M{
		"longurl":  longurl,
		"shorturl": shorturl,

		"createdAt": primitive.NewDateTimeFromTime(time.Now()),
	})

	if err != nil {
		return false, primitive.NilObjectID, err
	}

	return true, info.InsertedID.(primitive.ObjectID), nil
}

func GetUrl(shorturl string) (bool, string, error) {
	var result bson.M
	err := collections.Urls().FindOne(context.Background(), bson.M{"shorturl": shorturl}).Decode(&result)

	if err != nil {
		return false, "", err
	}

	return true, result["longurl"].(string), nil
}

func IsExistLongurl(longurl string) (bool, string, error) {
	var result bson.M
	err := collections.Urls().FindOne(context.Background(), bson.M{"longurl": longurl}).Decode(&result)

	if err != nil {
		return false, "", err
	}

	return true, result["shorturl"].(string), nil
}

func IsExistShorturl(shorturl string) (bool, string, error) {
	var result bson.M
	err := collections.Urls().FindOne(context.Background(), bson.M{"shorturl": shorturl}).Decode(&result)

	if err != nil {
		return false, "", err
	}

	return true, result["longurl"].(string), nil
}
