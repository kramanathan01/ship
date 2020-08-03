package main

import (
	vp "github.com/kramanathan01/ship/ship-service-vessel/proto/vessel"
)

// MarshalVessel -
func MarshalVessel(vessel *vp.Vessel) *Vessel {
	return &Vessel{
		ID:        vessel.Id,
		Capacity:  vessel.Capacity,
		MaxWeight: vessel.MaxWeight,
		Name:      vessel.Name,
		OwnerID:   vessel.OwnerId,
		Available: vessel.Available,
	}
}

// MarshalSpecification -
func MarshalSpecification(spec *vp.Specification) *Specification {
	return &Specification{
		Capacity: spec.Capacity,
		MaxWeight: spec.MaxWeight,
	}
}

// UnMarshalVessel -
func UnMarshalVessel(vessel *Vessel) *vp.Vessel {
	return &vp.Vessel{
		Id: vessel.ID,
		Capacity: vessel.Capacity,
		MaxWeight: vessel.MaxWeight,
		Name: vessel.Name,
		OwnerId: vessel.OwnerID,
		Available: vessel.Available,
	}
}
