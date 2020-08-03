package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/kramanathan01/ship/ship-service-consignment/proto/consignment"
	vp "github.com/kramanathan01/ship/ship-service-vessel/proto/vessel"
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

	service := micro.NewService(micro.Name("ship.cli.consignment"))
	service.Init()

	client := pb.NewShippingService("ship.service.consignment", service.Client())
	vesselClient := vp.NewVesselService("ship.service.vessel", service.Client())
	vessels := &vp.Vessel{Id: "vessel001", Name: "Boaty McBoatface", MaxWeight: 200000, Capacity: 500}
	vv, err := vesselClient.Create(context.Background(), vessels)
	if err != nil {
		log.Printf("Vessel not created: %v", err)
	}
	log.Printf("Vessel, %v", vv.Created)

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
