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
