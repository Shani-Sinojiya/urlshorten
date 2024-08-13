package handlers

import (
	"fmt"
	"math/rand"

	"github.com/gofiber/fiber/v2"
	"urlshorten.kwikbill.in/functions"
)

func UrlShort(c *fiber.Ctx) error {
	type Request struct {
		Url string `json:"url"`
	}

	var body Request

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invelid body",
			"success": false,
		})
	}

	// Check if the url is already shortened
	_, shorturl, err := functions.IsExistLongurl(body.Url)

	if err == nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message":  "url already shortened",
			"success":  true,
			"shortUrl": fmt.Sprintf("http://localhost:3000/r/%s", shorturl),
		})
	}

	shortUrl := randomString(6)

	_, _, err = functions.IsExistShorturl(shortUrl)

	for err == nil {
		shortUrl = randomString(6)
		_, _, err = functions.IsExistShorturl(shortUrl)
	}

	_, _, err = functions.CreateUrl(body.Url, shortUrl)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
			"success": false,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":  "url shorten successfully",
		"success":  true,
		"shortUrl": fmt.Sprintf("http://localhost:3000/r/%s", shortUrl),
	})
}

func randomString(length uint8) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func GetUrl(c *fiber.Ctx) error {
	shortUrl := c.Params("shorturl")

	url, err := functions.GetCacheUrl(shortUrl)

	if err == nil {
		return c.Redirect(url, fiber.StatusMovedPermanently)
	}

	_, longUrl, err := functions.IsExistShorturl(shortUrl)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "url not found",
			"success": false,
		})
	}

	err = functions.SetCacheUrl(shortUrl, longUrl)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
			"success": false,
		})
	}

	return c.Redirect(longUrl, fiber.StatusMovedPermanently)
}
