package main

import (
	"flag"
	"time"

	"github.com/GQDeltex/go-http-proxy/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Parse CLI Arguments
	loglevel := flag.String("loglevel", "WARNING", "Set the loglevel [DEBUG,INFO,WARNING,ERROR,FATAL]")
	caching_time := flag.Duration("cachetime", 30*time.Minute, "Set the time to cache stuff")
	flag.Parse()
	lvl, err := log.ParseLevel(*loglevel)
	if err != nil {
		log.Fatal("Could not Parse Loglevel")
	}
	log.Info("LogLevel: ", lvl.String())
	log.Info("CacheTime: ", caching_time.String())

	// Setup logger and fiber
	log.SetLevel(lvl)
	app := fiber.New()

	// Setup Cache Middleware
	app.Use(cache.New(cache.Config{
		Expiration:   *caching_time,
		CacheControl: true,
	}))

	// The main route
	app.Get("/*", func(c *fiber.Ctx) error {
		// Parse the URL
		urlstr := c.Params("*")
		log.Debug("Got request for ", urlstr)
		// Check the url and parse
		parsedUrl, err := utils.ParseURL(urlstr)
		if err != nil {
			log.Error(err)
			return err
		}
		log.Debug(parsedUrl)
		// Request the resources on the remote server
		code, contenttype, body, err := utils.DoHttpRequest(parsedUrl)
		log.Debug("Content-Type: ", contenttype)
		// Return the data
		return c.Status(code).Type(contenttype).Send(body)
	})

	// Start the Webserver
	log.Fatal(app.Listen(":3000"))
}
