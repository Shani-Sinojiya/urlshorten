package constants

import "urlshorten.kwikbill.in/config"

var (
	// mongodb
	DATABASE_MONGODB_URI     = config.GetConfig().GetString("databases.mongodb.uri")
	DATABASE_MONGODB_DB_NAME = config.GetConfig().GetString("databases.mongodb.db")

	// redis
	DATABASE_REDIS_HOST     = config.GetConfig().GetString("databases.redis.host")
	DATABASE_REDIS_PORT     = config.GetConfig().GetString("databases.redis.port")
	DATABASE_REDIS_PASSWORD = config.GetConfig().GetString("databases.redis.password")
	DATABASE_REDIS_DB       = config.GetConfig().GetInt("databases.redis.db")
	DATABASE_REDIS_QDB      = config.GetConfig().GetInt("databases.redis.qdb")
)
