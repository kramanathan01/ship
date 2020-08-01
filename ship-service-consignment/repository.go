package main

import (
	"context"
	"log"

	pb "github.com/kramanathan01/ship/ship-service-consignment/proto/consignment"
	"go.mongodb.org/mongo-driver/mongo"
)

// Consignment - Struct for ser/deser
type Consignment struct {
	ID          string     `json:"id"`
	Weight      int32      `json:"weight"`
	Description string     `json:"description"`
	Containers  Containers `json:"containers"`
	VesselID    string     `json:"vessel_id"`
}

// Container - Struct for ser/deser
type Container struct {
	ID         string `json:"id"`
	CustomerID string `json:"customer_id"`
	UserID     string `json:"user_id"`
}

// Containers - Slice for ser/deser
type Containers []*Container

func MarshalContainerCollection(containers []*pb.Container) []*Container {
	collection := make([]*Container, 0)
	for _, container := range containers {
		collection = append(collection, MarshalContainer(container))
	}
	return collection
}

func UnmarshalContainerCollection(containers []*Container) []*pb.Container {
	collection := make([]*pb.Container, 0)
	for _, container := range containers {
		collection = append(collection, UnmarshalContainer(container))
	}
	return collection
}

func UnmarshalConsignmentCollection(consignments []*Consignment) []*pb.Consignment {
	collection := make([]*pb.Consignment, 0)
	for _, consignment := range consignments {
		collection = append(collection, UnmarshalConsignment(consignment))
	}
	return collection
}

func UnmarshalContainer(container *Container) *pb.Container {
	return &pb.Container{
		Id:         container.ID,
		CustomerId: container.CustomerID,
		UserId:     container.UserID,
	}
}

func MarshalContainer(container *pb.Container) *Container {
	return &Container{
		ID:         container.Id,
		CustomerID: container.CustomerId,
		UserID:     container.UserId,
	}
}

// Marshal an input consignment type to a consignment model
func MarshalConsignment(consignment *pb.Consignment) *Consignment {
	containers := MarshalContainerCollection(consignment.Containers)
	return &Consignment{
		ID:          consignment.Id,
		Weight:      consignment.Weight,
		Description: consignment.Description,
		Containers:  containers,
		VesselID:    consignment.VesselId,
	}
}

func UnmarshalConsignment(consignment *Consignment) *pb.Consignment {
	return &pb.Consignment{
		Id:          consignment.ID,
		Weight:      consignment.Weight,
		Description: consignment.Description,
		Containers:  UnmarshalContainerCollection(consignment.Containers),
		VesselId:    consignment.VesselID,
	}
}

type repository interface {
	Create(context.Context, *Consignment) error
	GetAll(context.Context) ([]*Consignment, error)
}

// // Repository - Dummy Repo for now
// type Repository struct {
// 	mu           sync.RWMutex
// 	consignments []*pb.Consignment
// }

// Create a new Consignment
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

// MongoRepository implementation
type MongoRepository struct {
	collection *mongo.Collection
}

// Create -
func (repository *MongoRepository) Create(ctx context.Context, consignment *Consignment) error {
	_, err := repository.collection.InsertOne(ctx, consignment)
	return err
}

// GetAll -
func (repository *MongoRepository) GetAll(ctx context.Context) ([]*Consignment, error) {
	cur, err := repository.collection.Find(ctx, nil, nil)
	var consignments []*Consignment
	for cur.Next(ctx) {
		var consignment *Consignment
		if err := cur.Decode(&consignment); err != nil {
			log.Printf("Repo -- Decode Error: %v", err)
			return nil, err
		}
		consignments = append(consignments, consignment)
	}
	log.Println("REPO: Normal Return")
	return consignments, err
}
