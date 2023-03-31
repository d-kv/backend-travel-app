package main

import (
	"log"

	"github.com/d-kv/backend-travel-app/cmd/place-service/config"
)

func main() {
	cfg, _ := config.New("cmd/place-service", "config.example")
	log.Print(cfg.DB.Mongo.URI)
}
