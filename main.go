package main

import (
	"github.com/gofiber/fiber/v2"
	"urlshorten.kwikbill.in/database"
	"urlshorten.kwikbill.in/handlers"
)

func main() {
	defer database.DisconnectMongo()
	app := fiber.New()

	app.Post("/urlshort", handlers.UrlShort)
	app.Get("/r/:shorturl", handlers.GetUrl)

	app.Listen(":3000")
}
