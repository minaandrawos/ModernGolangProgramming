package main

import (
	"dino/databaselayer"
	"dino/dinowebportal/dinoapi"
	"log"
)

func main() {
	db, err := databaselayer.GetDatabaseHandler(databaselayer.MONGODB, "mongodb://127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}
	dinoapi.RunApi("localhost:8080", db)
}
