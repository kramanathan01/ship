package main

import (
	"context"
	"log"
	"os"

	pb "github.com/kramanathan01/ship/ship-service-consignment/proto/consignment"
	vp "github.com/kramanathan01/ship/ship-service-vessel/proto/vessel"
	"github.com/micro/go-micro/v2"
)

const (
	defaultHost = "datastore:27017"
)

func main() {

	// Create a new service
	service := micro.NewService(
		micro.Name("ship.service.consignment"),
	)

	// Init parses command line flags to pass in options
	service.Init()

	// DB Connection
	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}

	client, err := CreateClient(context.Background(), uri, 0)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.Background())
	consignmentCollection := client.Database("ship").Collection("consignments")
	repository := &MongoRepository{consignmentCollection}

	vesselClient := vp.NewVesselService("ship.service.vessel", service.Client())

	// Register our service as handler for protobuf interface
	h := &handler{repository, vesselClient}
	if err := pb.RegisterShippingServiceHandler(service.Server(), h); err != nil {
		log.Panic(err)
	}

	if err := service.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
