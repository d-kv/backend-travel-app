package main

import (
	"log"

	"github.com/d-kv/backend-travel-app/cmd/afterwork/config"
)

func main() {
	c, _ := config.New("cmd/afterwork", "config.example")
	log.Print(c.DB.Mongo.CollectionName)
}
