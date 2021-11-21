package main

import (
	"errors"
	"strconv"

	"github.com/GQDeltex/go-http-proxy/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Parse CLI Arguments
	cliArgs, err := utils.ParseCLIArgs()
	if err != nil {
		log.Fatal(err.Error())
	}
	// Setup logger and fiber
	log.SetLevel(cliArgs.LogLevel)
	app := fiber.New()

	// Setup Cache Middleware
	app.Use(cache.New(cache.Config{
		Expiration:   cliArgs.CachingTime,
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
		err = utils.ValidateToken(token, parsedUrl.String(), cliArgs.Secret, expires)
		if err != nil {
			log.Error(err.Error())
			return err
		}
		log.Debug(parsedUrl)
		// Request the resources on the remote server
		code, headers, body, err := utils.DoHttpRequest(parsedUrl, cliArgs.AllowedHeaders, cliArgs.UserAgent)
		log.Debug("Sending Headers: ", headers)
		// Return the data
		for header, value := range headers {
			c.Response().Header.Add(header, value)
		}
		c.Response().SetStatusCode(code)
		return c.Send(body)
	})

	// Start the Webserver
	log.Fatal(app.Listen(":" + strconv.Itoa(cliArgs.Port)))
}
