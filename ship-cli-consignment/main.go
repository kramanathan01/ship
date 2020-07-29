package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/kramanathan01/ship/ship-service-consignment/proto/consignment"
	micro "github.com/micro/go-micro/v2"
)

const (
	address         = "localhost:50051"
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err
}

func main() {
	// Setup connection to gRPC server
	// conn, err := grpc.Dial(address, grpc.WithInsecure())
	// if err != nil {
	// 	log.Fatalf("Did not connect: %v", err)
	// }
	// defer conn.Close()
	// client := pb.NewShippingServiceClient(conn)

	service := micro.NewService(micro.Name("ship.consignment.cli"))
	service.Init()

	client := pb.NewShippingService("ship.consignment.service", service.Client())

	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)
	if err != nil {
		log.Fatalf("Unable to parse file: %v", err)
	}

	r, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("Unable to create consignment: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	ga, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Unable to get Consignments: %v", err)
	}

	for _, v := range ga.Consignments {
		log.Println(v)
	}
}
