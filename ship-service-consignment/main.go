package main

import (
	"context"
	"log"
	"sync"

	pb "github.com/kramanathan01/ship/ship-service-consignment/proto/consignment"
	vp "github.com/kramanathan01/ship/ship-service-vessel/proto/vessel"
	"github.com/micro/go-micro/v2"
)

const (
	port = ":50501"
)

type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

// Repository - Dummy Repo for now
type Repository struct {
	mu           sync.RWMutex
	consignments []*pb.Consignment
}

// Create a new Consignment
func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.mu.Lock()
	u := append(repo.consignments, consignment)
	repo.consignments = u
	repo.mu.Unlock()

	return consignment, nil
}

// GetAll - Returns all existing Consignments
func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

// Service implements the gRPC interface
type consignmentService struct {
	repo         repository
	vesselClient vp.VesselService
}

// CreateConsignment - Create method that takes in a protobuf Consignment and
// returns protobuf Response message

func (s *consignmentService) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {

	vesselResponse, err := s.vesselClient.FindAvailable(context.Background(), &vp.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	if err != nil {
		return err
	}
	req.VesselId = vesselResponse.Vessel.Id

	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}
	res.Created = true
	res.Consignment = consignment
	return nil
}

func (s *consignmentService) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments := s.repo.GetAll()
	res.Consignments = consignments
	return nil
}

func main() {

	repo := &Repository{}

	// Create a new service
	service := micro.NewService(
		micro.Name("ship.service.consignment"),
	)

	vesselClient := vp.NewVesselService("ship.service.vessel", service.Client())

	// Init parses command line flags to pass in options
	service.Init()

	// Register our service as handler for protobuf interface
	if err := pb.RegisterShippingServiceHandler(service.Server(), &consignmentService{repo, vesselClient}); err != nil {
		log.Panic(err)
	}

	// log.Println("Running on port:", port)
	if err := service.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
