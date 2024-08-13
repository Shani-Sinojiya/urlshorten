package databases

import (
	"context"

	"github.com/go-redis/redis/v8"
	"urlshorten.kwikbill.in/constants"
)

// RedisClient is the Redis client
var RedisClient *redis.Client

// ConnectRedis connects to Redis
func ConnectRedis() error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     constants.DATABASE_REDIS_HOST + ":" + constants.DATABASE_REDIS_PORT,
		Password: constants.DATABASE_REDIS_PASSWORD,
		DB:       constants.DATABASE_REDIS_DB,
	})

	err := RedisClient.Ping(context.Background()).Err()

	return err
}

// DisconnectRedis disconnects from Redis
func DisconnectRedis() error {
	if err := RedisClient.Close(); err != nil {
		return err
	}

	return nil
}
