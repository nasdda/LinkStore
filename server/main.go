package main

import (
	"log"

	"github.com/nasdda/linkstore/config"
	"github.com/nasdda/linkstore/internal/app"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}
	app.Run(cfg)
}
