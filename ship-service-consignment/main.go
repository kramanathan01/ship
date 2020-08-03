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

// type repository interface {
// 	Create(*pb.Consignment) (*pb.Consignment, error)
// 	GetAll() []*pb.Consignment
// }

// // Repository - Dummy Repo for now
// type Repository struct {
// 	mu           sync.RWMutex
// 	consignments []*pb.Consignment
// }

// // Create a new Consignment
// func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
// 	repo.mu.Lock()
// 	u := append(repo.consignments, consignment)
// 	repo.consignments = u
// 	repo.mu.Unlock()

// 	return consignment, nil
// }

// // GetAll - Returns all existing Consignments
// func (repo *Repository) GetAll() []*pb.Consignment {
// 	return repo.consignments
// }

// Service implements the gRPC interface
type consignmentService struct {
	repo         repository
	vesselClient vp.VesselService
}

// CreateConsignment - Create method that takes in a protobuf Consignment and
// returns protobuf Response message

// func (s *consignmentService) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {

// 	vesselResponse, err := s.vesselClient.FindAvailable(context.Background(), &vp.Specification{
// 		MaxWeight: req.Weight,
// 		Capacity:  int32(len(req.Containers)),
// 	})
// 	if err != nil {
// 		return err
// 	}
// 	req.VesselId = vesselResponse.Vessel.Id

// 	consignment, err := s.repo.Create(req)
// 	if err != nil {
// 		return err
// 	}
// 	res.Created = true
// 	res.Consignment = consignment
// 	return nil
// }

// func (s *consignmentService) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
// 	consignments := s.repo.GetAll()
// 	res.Consignments = consignments
// 	return nil
// }

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
