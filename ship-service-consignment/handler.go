package main

import (
	"context"
	"log"

	pb "github.com/kramanathan01/ship/ship-service-consignment/proto/consignment"
	vp "github.com/kramanathan01/ship/ship-service-vessel/proto/vessel"
	"github.com/pkg/errors"
)

type handler struct {
	repository
	vesselClient vp.VesselService
}

// CreateConsignment - Takes a context and a request as an
// argument, these are handled by the gRPC server.
func (s *handler) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {

	// Here we call a client instance of our vessel service with our consignment weight,
	// and the amount of containers as the capacity value
	vesselResponse, err := s.vesselClient.FindAvailable(ctx, &vp.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	if vesselResponse == nil {
		return errors.New("error fetching vessel, returned nil")
	}

	if err != nil {
		return err
	}

	// We set the VesselId as the vessel we got back from our
	// vessel service
	req.VesselId = vesselResponse.Vessel.Id

	// Save our consignment
	if err = s.repository.Create(ctx, MarshalConsignment(req)); err != nil {
		return err
	}

	res.Created = true
	res.Consignment = req
	return nil
}

// GetConsignments -
func (s *handler) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments, err := s.repository.GetAll(ctx)
	if err != nil {
		log.Printf("Handler Error: %v", err)
		return err
	}
	res.Consignments = UnmarshalConsignmentCollection(consignments)
	return nil
}
