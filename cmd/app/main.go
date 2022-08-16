package main

import (
	"log"

	"github.com/cut4cut/toimi-test-work/config"
	"github.com/cut4cut/toimi-test-work/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
