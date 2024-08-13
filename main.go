package main

import (
	"log"
	"os"
	"time"

	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"urlshorten.kwikbill.in/constants"
	"urlshorten.kwikbill.in/databases"
	"urlshorten.kwikbill.in/handlers"
)

func main() {
	defer databases.CloseDB()

	app := fiber.New(fiber.Config{
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
		Prefork:       isProduction(),
		CaseSensitive: true,
		ServerHeader:  "KwikBill v2 Url Shorten",
		AppName:       "KwikBill v2 Url Shorten",
	})

	app.Use(requestid.New())
	app.Use(recover.New())

	app.Use(etag.New(etag.Config{
		Weak: true,
	}))

	if !isProduction() {
		app.Use(logger.New(logger.Config{
			Format:     "[${ip}]:${port} ${time} ${status} - ${method} ${latency} ${path}\n",
			TimeFormat: "02-Jan-2006 15:04:05",
			TimeZone:   "Asia/Kolkata",
		}))
	} else {
		file, err := os.OpenFile("./urlshortenbackend"+time.Now().Format("dd/mm/yyyy hh:mm:ss a")+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}

		app.Use(logger.New(logger.Config{
			Format:     "[${ip}]:${port} ${time} ${status} - ${method} ${latency} ${path}\n",
			TimeFormat: "02-Jan-2006 15:04:05",
			TimeZone:   "Asia/Kolkata",
			Output:     file,
		}))
	}

	app.Post("/urlshort", handlers.UrlShort)
	app.Get("/r/:shorturl", handlers.GetUrl)

	app.Listen(constants.SERVER_PORT)
}

func isProduction() bool {
	if constants.SERVER_MODE == "production" {
		return true
	} else {
		return false
	}
}
