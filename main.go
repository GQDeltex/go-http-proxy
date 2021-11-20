package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Setup logger and fiber
	log.SetLevel(log.DebugLevel)
	app := fiber.New()

	// Setup Cache Middleware
	app.Use(cache.New(cache.Config{
		Expiration:   30 * time.Minute,
		CacheControl: true,
	}))

	// The main route
	app.Get("/:url", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	// Start the Webserver
	log.Fatal(app.Listen(":3000"))
}
