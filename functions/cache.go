package functions

import (
	"context"
	"time"

	"urlshorten.kwikbill.in/constants"
	"urlshorten.kwikbill.in/databases"
)

func GetCacheUrl(short string) (string, error) {
	info := databases.RedisClient.Get(context.Background(), short)

	if info.Err() != nil {
		return "", info.Err()
	}

	return info.Val(), nil
}

func SetCacheUrl(short string, long string) error {
	info := databases.RedisClient.Set(context.Background(), short, long, time.Duration(constants.CACHE_TIME)*time.Minute)

	if info.Err() != nil {
		return info.Err()
	}

	return nil
}
