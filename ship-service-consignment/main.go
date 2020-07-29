package main

import (
	"context"
	"log"
	"sync"

	pb "github.com/kramanathan01/ship/ship-service-consignment/proto/consignment"
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
type service struct {
	repo repository
}

// CreateConsignment - Create method that takes in a protobuf Consignment and
// returns protobuf Response message

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {

	consignment, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}

	return &pb.Response{Created: true, Consignment: consignment}, nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	consignments := s.repo.GetAll()
	return &pb.Response{Consignments: consignments}, nil
}

func main() {

	repo := &Repository{}

	// Create a new service
	service := micro.NewService(
		micro.Name("ship.service.consignment"),
	)

	// Init parses command line flags to pass in options
	service.Init()

	// Register our service as handler for protobuf interface
	if err := pb.RegisterShippingServiceHandler(service.Server(), &consignmentService{repo}); err != nil {
		log.Panic(err)
	}

	// log.Println("Running on port:", port)
	if err := service.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
