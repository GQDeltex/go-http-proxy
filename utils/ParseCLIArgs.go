package utils

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

type CliArgs struct {
	LogLevel       log.Level
	CachingTime    time.Duration
	AllowedHeaders []string
	Secret         string
	Port           int
	UserAgent      string
}

func ParseCLIArgs() (*CliArgs, error) {
	// Setup arguments
	loglevel := flag.String("loglevel", "WARNING", "Set the loglevel [DEBUG,INFO,WARNING,ERROR,FATAL]")
	caching_time := flag.Duration("cachetime", 30*time.Minute, "Set the time to cache stuff")
	secret := flag.String("secret", "secret", "Set the secret used for signing the requests")
	allowed_headers := flag.String("allowedheaders", "Content-Type", "Set the allowed headers to be copied over")
	port := flag.Int("port", 3000, "Port this software should run on")
	user_agent := flag.String("useragent", "go-http-proxy", "Custom Useragent to be sent with every remote request")
	// Check the command line
	flag.Parse()
	// Process variables
	var cliArgs = CliArgs{}
	lvl, err := log.ParseLevel(*loglevel)
	if err != nil {
		return nil, err
	}
	cliArgs.LogLevel = lvl
	log.Info("LogLevel: ", lvl)
	cliArgs.Secret = *secret
	log.Info("Secret: *****") // Do not print real secret
	cliArgs.CachingTime = *caching_time
	log.Info("CachingTime: ", *caching_time)
	cliArgs.AllowedHeaders = strings.Split(*allowed_headers, ",")
	log.Info("AllowedHeaders: ", *allowed_headers)
	cliArgs.Port = *port
	log.Info("Port: ", *port)
	cliArgs.UserAgent = *user_agent
	log.Info("UserAgent: ", *user_agent)
	return &cliArgs, nil
}
