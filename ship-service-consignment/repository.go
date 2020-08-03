package main

import (
	"context"
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
