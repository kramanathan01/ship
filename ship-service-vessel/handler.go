package main

import (
	"log"

	vp "github.com/kramanathan01/ship/ship-service-vessel/proto/vessel"
	"golang.org/x/net/context"
)

type handler struct {
	repository
}

// Create - Handler for grpc endpoint
func (r *handler) Create(ctx context.Context, req *vp.Vessel, res *vp.Response) error {
	if err := r.repository.Create(ctx, MarshalVessel(req)); err != nil {
		return err
	}
	res.Created = true
	res.Vessel = req
	return nil
}

// FindAvailable - Handler for grpc endpoint
func (r *handler) FindAvailable(ctx context.Context, req *vp.Specification, res *vp.Response) error {
	vessel, err := r.repository.FindVessel(ctx, MarshalSpecification(req))
	if err != nil {
		log.Printf("Handler Error: %v", err)
		return err
	}
	res.Vessel = UnMarshalVessel(vessel)
	return nil
}
