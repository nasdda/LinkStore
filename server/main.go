package main

import (
	"fmt"
	"log"

	"github.com/nasdda/linkstore/config"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	fmt.Println("Config loaded:", config)

}
