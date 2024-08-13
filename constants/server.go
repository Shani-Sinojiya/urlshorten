package constants

import "urlshorten.kwikbill.in/config"

var (
	// server
	SERVER_PORT = config.GetConfig().GetString("server.port")
	SERVER_MODE = config.GetConfig().GetString("server.mode")
)
