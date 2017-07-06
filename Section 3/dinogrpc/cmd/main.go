package main

import (
	"flag"
	"fmt"
	"net"
	"strings"

	"golang.org/x/net/context"

	"log"

	"dino/communicationlayer/dinogrpc"
	"dino/databaselayer"

	"io"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	op := flag.String("op", "s", "s for server, and c for client ") //proto2test -op s => will run as a server, proto2test -op c => will run as a client
	flag.Parse()
	switch strings.ToLower(*op) {
	case "s":
		runGRPCServer()
	case "c":
		runGRPCClient()
	}
}

func runGRPCServer() {
	grpclog.Println("Starting GRPC Server")
	lis, err := net.Listen("tcp", ":8282")
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	grpclog.Println("Listening on 127.0.0.1:8282")

	grpcServer := grpc.NewServer()
	dinoServer, err := dinogrpc.NewDinoGrpcServer(databaselayer.MONGODB, "mongodb://127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}
	dinogrpc.RegisterDinoServiceServer(grpcServer, dinoServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}

func runGRPCClient() {
	conn, err := grpc.Dial("127.0.0.1:8282", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := dinogrpc.NewDinoServiceClient(conn)
	input := ""
	fmt.Println("All animals?(y/n)")
	fmt.Scanln(&input)
	if strings.EqualFold(input, "y") {
		animals, err := client.GetAllAnimals(context.Background(), &dinogrpc.Request{})
		if err != nil {
			log.Fatal(err)
		}

		for {
			animal, err := animals.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				grpclog.Fatal(err)
			}
			grpclog.Println(animal)
		}
		return
	}
	fmt.Println("Nickname?")
	fmt.Scanln(&input)
	a, err := client.GetAnimal(context.Background(), &dinogrpc.Request{Nickname: input})
	if err != nil {
		log.Fatal(err)
	}
	grpclog.Println(*a)
}
