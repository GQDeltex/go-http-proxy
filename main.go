package main

import (
	"errors"
	"flag"
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
			log.Error(err)
			return err
		}
		// Extract the content-type
		content_type := resp.Header["Content-Type"][0]
		content_type = strings.Split(content_type, "/")[1]
		log.Debug("Content-Type: ", content_type)
		return c.Status(resp.StatusCode).Type(content_type).Send(body)
	})

	// Start the Webserver
	log.Fatal(app.Listen(":3000"))
}
