package main

import (
	"flag"
	"log"

	"github.com/AHTI6IOTIK/anti-bruteforce/config"
	"github.com/AHTI6IOTIK/anti-bruteforce/internal/app"
	_ "github.com/lib/pq"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "config.yml", "Path to configuration file")
}

func main() {
	flag.Parse()

	// Configuration
	cfg, err := config.GetConfig(configFile)
	if err != nil {
		log.Fatalf("Loading config: %v", err)
	}

	// Run
	app.Run(cfg)
}
