package main

import (
	"context"
	"log"
	"os"

	pb "github.com/kramanathan01/ship/ship-service-vessel/proto/vessel"
	"github.com/micro/go-micro/v2"
)

const (
	defaultHost = "datastore:27017"
)

// // Repository - Interface for the repp
// type Repository interface {
// 	FindAvailable(*pb.Specification) (*pb.Vessel, error)
// 	// Create(*pb.Vessel) error
// }

// // VesselRepository - Struct for repo
// type VesselRepository struct {
// 	vessels []*pb.Vessel
// }

// var vessels = make([]*pb.Vessel, 2, 2)

// // FindAvailable - checks a specification against a map of vessels,
// // if capacity and max weight are below a vessels capacity and max weight,
// // then return that vessel.
// func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
// 	for _, vessel := range repo.vessels {
// 		if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight {
// 			return vessel, nil
// 		}
// 	}
// 	return nil, errors.New("No vessel found by that spec")
// }

// // Our grpc service handler
// type vesselService struct {
// 	repo Repository
// }

// func (s *vesselService) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {

// 	// Find the next available vessel
// 	vessel, err := s.repo.FindAvailable(req)
// 	if err != nil {
// 		return err
// 	}

// 	// Set the vessel as part of the response message type
// 	res.Vessel = vessel
// 	return nil
// }

func main() {
	// vessels := []*pb.Vessel{
	// 	&pb.Vessel{Id: "vessel001", Name: "Boaty McBoatface", MaxWeight: 200000, Capacity: 500},
	// }

	// vessels := &pb.Vessel{Id: "vessel001", Name: "Boaty McBoatface", MaxWeight: 200000, Capacity: 500}
	// repo := &VesselRepository{vessels}

	service := micro.NewService(
		micro.Name("ship.service.vessel"),
	)

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
	vesselCollection := client.Database("ship").Collection("vessels")
	repository := &MongoRepository{vesselCollection}

	// Register our implementation
	if err := pb.RegisterVesselServiceHandler(service.Server(), &handler{repository}); err != nil {
		log.Panic(err)
	}

	if err := service.Run(); err != nil {
		log.Panic(err)
	}
}
