package main

import (
	"bytes"
	"encoding/json"
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
		Nickname:   "Ocri",
		Zone:       3,
		Age:        19,
	}
	var b bytes.Buffer
	json.NewEncoder(&b).Encode(data)
	/*
		resp, err := http.Post("http://localhost:8181/api/dinos/add", "application/json", &b)
		if err != nil || resp.StatusCode != 200 {
			log.Fatal(err)
		}


			resp, err := http.Post("http://localhost:8181/api/dinos/edit/patro", "application/json", &b)
			if err != nil || resp.StatusCode != 200 {
				log.Fatal(resp.Status, err)
			}
	*/
}
