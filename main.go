package main

import (
	"errors"
	"flag"
	"strconv"
	"strings"
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
	secret := flag.String("secret", "secret", "Set the secret used for signing the requests")
	allowed_headers := flag.String("allowedheaders", "Content-Type", "Set the allowed headers to be copied over")
	flag.Parse()
	lvl, err := log.ParseLevel(*loglevel)
	if err != nil {
		log.Fatal("Could not Parse Loglevel")
	}
	allowedHeaders := strings.Split(*allowed_headers, ",")
	log.Info("LogLevel: ", lvl.String())
	log.Info("CacheTime: ", caching_time.String())
	log.Info("AllowedHeaders: ", allowedHeaders)

	// Setup logger and fiber
	log.SetLevel(lvl)
	app := fiber.New()

	// Setup Cache Middleware
	app.Use(cache.New(cache.Config{
		Expiration:   *caching_time,
		CacheControl: true,
	}))

	// The main route
	app.Get("/", func(c *fiber.Ctx) error {
		// Parse the Parameters
		urlstr := c.Query("url")
		token := c.Query("token")
		expiry := c.Query("expires")
		if token == "" || expiry == "" || urlstr == "" {
			log.Error(errors.New("Not all necessary query parameters given"))
			return fiber.ErrUnauthorized
		}
		log.Debug("Got request for ", urlstr)
		// Check the url and the query parameters
		parsedUrl, err := utils.ParseURL(urlstr)
		if err != nil {
			log.Error(err)
			return err
		}
		expires, err := strconv.ParseInt(expiry, 10, 64)
		if err != nil {
			log.Error(err.Error())
			return err
		}
		err = utils.ValidateToken(token, parsedUrl.String(), *secret, expires)
		if err != nil {
			log.Error(err.Error())
			return err
		}
		log.Debug(parsedUrl)
		// Request the resources on the remote server
		code, headers, body, err := utils.DoHttpRequest(parsedUrl, allowedHeaders)
		log.Debug("Sending Headers: ", headers)
		// Return the data
		for header, value := range headers {
			c.Response().Header.Add(header, value)
		}
		c.Response().SetStatusCode(code)
		return c.Send(body)
	})

	// Start the Webserver
	log.Fatal(app.Listen(":3000"))
}
