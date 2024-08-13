package constants

import "urlshorten.kwikbill.in/config"

var (
	CACHE_TIME = config.GetConfig().GetInt64("cacheTime")
)
