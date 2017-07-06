package main

import (
	"dino/communicationlayer/dinoproto3"
	"dino/databaselayer"
	"flag"
	"fmt"
	"net"
	"strings"

	"io/ioutil"

	"log"

	"github.com/golang/protobuf/proto"
)

/*
	1- We will serialize some data via proto3
	2- We will send this data via TCP to a different service
	3- We will deserialize the data via proto3, and print out the extracted values

	(client) ---> (server)
	A- A TCP client needs to be written to send the data
	B- A TCP Server to receive the data
*/

func main() {
	op := flag.String("op", "s", "s for server, c for client") //proto3test -op s => will run as a server, proto3test -op c => will run as a client
	flag.Parse()
	switch strings.ToLower(*op) {
	case "s":
		RunProto3Server()
	case "c":
		RunProto3Client()
	}
}

func RunProto3Client() {
	a := &dinoproto3.Animal{
		Id:         1,
		AnimalType: "Raptor",
		Nickname:   "rapto",
		Zone:       3,
		Age:        21,
	}
	data, err := proto.Marshal(a)
	if err != nil {
		log.Fatal(err)
	}
	SendData(data)
}

func SendData(data []byte) {
	c, err := net.Dial("tcp", "127.0.0.1:8282")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	c.Write(data)
}

func RunProto3Server() {
	l, err := net.Listen("tcp", ":8282")
	if err != nil {
		log.Fatal(err)
	}
	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		defer l.Close()
		go func(c net.Conn) {
			defer c.Close()
			data, err := ioutil.ReadAll(c)
			if err != nil {
				return
			}
			a := &dinoproto3.Animal{}
			err = proto.Unmarshal(data, a)
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Println(a)
		}(c)
	}
}

func SendDBToServer() {
	handler, err := databaselayer.GetDatabaseHandler(databaselayer.MONGODB, "mongodb://127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}
	animals, err := handler.GetAvailableDynos()
	for _, animal := range animals {
		a := &dinoproto3.Animal{
			Id:         int32(animal.ID),
			AnimalType: animal.AnimalType,
			Nickname:   animal.Nickname,
			Zone:       int32(animal.Zone),
			Age:        int32(animal.Age),
		}
		data, err := proto.Marshal(a)
		if err != nil {
			log.Fatal(err)
		}
		SendData(data)
	}
}
