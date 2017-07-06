package main

import (
	"dino/communicationlayer/dinoproto2"
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
	1- We will serialize some data via proto2
	2- We will send this data via TCP to a different service
	3- We will deserialize the data via proto2, and print out the extracted values

	A- A TCP client needs to be written to send the data
	B- A TCP Server to receive the data
*/
func main() {
	op := flag.String("op", "s", "s for server, and c for client ") //proto2test -op s => will run as a server, proto2test -op c => will run as a client
	flag.Parse()
	switch strings.ToLower(*op) {
	case "s":
		RunProto2Server()
	case "c":
		RunProto2Client()
	}
}

func RunProto2Client() {
	a := &dinoproto2.Animal{
		Id:         proto.Int(1),
		AnimalType: proto.String("Raptor"),
		Nickname:   proto.String("rapto"),
		Zone:       proto.Int(3),
		Age:        proto.Int(21),
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

func RunProto2Server() {
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
		go func(cn net.Conn) {
			defer cn.Close()
			data, err := ioutil.ReadAll(cn)
			if err != nil {
				return
			}
			a := &dinoproto2.Animal{}
			proto.Unmarshal(data, a)
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
		a := &dinoproto2.Animal{
			Id:         proto.Int(animal.ID),
			AnimalType: proto.String(animal.AnimalType),
			Nickname:   proto.String(animal.Nickname),
			Zone:       proto.Int(animal.Zone),
			Age:        proto.Int(animal.Age),
		}
		data, err := proto.Marshal(a)
		if err != nil {
			log.Fatal(err)
		}
		SendData(data)
	}
}
