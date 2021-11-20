package main

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
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
	app.Get("/*", func(c *fiber.Ctx) error {
		// Parse the URL
		urlstr := c.Params("*")
		if urlstr == "" {
			log.Error("No url Parameter")
			return errors.New("No URl parameter")
		}
		// Check the validity of the url
		url, err := url.Parse(urlstr)
		if err != nil {
			log.Error(err)
			return err
		}
		log.Debug("URL:", url.String())
		if url.Hostname() == "" {
			log.Error("No Host was given")
			return errors.New("No hostname was given")
		}
		// Do a http request to that URL
		resp, err := http.Get(url.String())
		if err != nil {
			log.Error(err)
			return err
		}
		// Read the response Body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		// Extract the content-type
		content_type := resp.Header["Content-Type"][0]
		content_type = strings.Split(content_type, "/")[1]
		log.Debug("Content-Type:", content_type)
		return c.Status(resp.StatusCode).Type(content_type).Send(body)
	})

	// Start the Webserver
	log.Fatal(app.Listen(":3000"))
}
