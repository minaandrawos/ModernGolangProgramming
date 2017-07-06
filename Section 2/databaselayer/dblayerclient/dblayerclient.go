package main

import "dino/databaselayer"
import "log"

func main() {
	//handler, err := databaselayer.GetDatabaseHandler(databaselayer.MYSQL, "gouser:gouser@/Dino")
	//handler, err := databaselayer.GetDatabaseHandler(databaselayer.SQLITE, "dino.db")
	//handler, err := databaselayer.GetDatabaseHandler(databaselayer.POSTGRESQL, "user=postgres dbname=dino sslmode=disable")
	handler, err := databaselayer.GetDatabaseHandler(databaselayer.MONGODB, "mongodb://127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}
	err = handler.UpdateAnimal(databaselayer.Animal{
		AnimalType: "Carnotaurus",
		Nickname:   "Carno",
		Zone:       3,
		Age:        23,
	}, "Carno")
	log.Println(err)

	log.Println(handler.GetAvailableDynos())
	log.Println(handler.GetDynoByNickname("rex"))
	log.Println(handler.GetDynosByType("Velociraptor"))
}
