package main

import (
	"fmt"

	"github.com/d-kv/backend-travel-app/cmd/afterwork/config"
)

func main() {
	c, _ := config.New("cmd/afterwork", "config.example")
	fmt.Print(c.DB.Mongo.CollectionName)
}
