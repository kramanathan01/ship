package main

import "context"

// Vessel -
type Vessel struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	OwnerID   string `json:"ownerid"`
	Capacity  int32  `json:"capacity"`
	MaxWeight int32  `json:"maxweight"`
	Available bool   `json:"available"`
}

// Specification -
type Specification struct {
	Capacity  int32 `json:"capacity"`
	MaxWeight int32 `json:"maxweight"`
}

type repository interface {
	FindVessel(context.Context, *Specification) (*Vessel, error)
	Create(context.Context, *Vessel) error
}
