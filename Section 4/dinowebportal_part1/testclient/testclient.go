package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type Animal struct {
	ID         int
	AnimalType string
	Nickname   string
	Zone       int
	Age        int
}

func main() {
	//Carnotaurus, Carno, 3, 22
	data := &Animal{
		AnimalType: "Velociraptor",
		Nickname:   "patro",
		Zone:       3,
		Age:        13,
	}
	var b bytes.Buffer
	json.NewEncoder(&b).Encode(data)
	/*
		resp, err := http.Post("http://localhost:8080/api/dinos/add", "application/json", &b)
		if err != nil || resp.StatusCode != 200 {
			log.Fatal(err)
		}
	*/

	resp, err := http.Post("http://localhost:8080/api/dinos/edit/patro", "application/json", &b)
	if err != nil || resp.StatusCode != 200 {
		log.Fatal(resp.Status, err)
	}

}
