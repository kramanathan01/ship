package main

import (
	"context"
	"log"
	"net"
	"sync"

	pb "github.com/kramanathan01/ship/ship-service-consignment/proto/consignment"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50501"
)

type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
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

func main() {

	repo := &Repository{}

	// Set up gRPC server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}
	s := grpc.NewServer()

	// Register our service as handler for protobuf interface
	pb.RegisterShippingServiceServer(s, &service{repo})
	reflection.Register(s)

	log.Println("Running on port:", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
