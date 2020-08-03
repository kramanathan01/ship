package main

import (
	"context"
	"sync"
)

// Repository - Dummy Repo for now
type Repository struct {
	mu           sync.RWMutex
	consignments []*Consignment
}

// Create - a new Consignment
func (repo *Repository) Create(ctx context.Context, consignment *Consignment) (*Consignment, error) {
	repo.mu.Lock()
	u := append(repo.consignments, consignment)
	repo.consignments = u
	repo.mu.Unlock()

	return consignment, nil
}

// GetAll - Returns all existing Consignments
func (repo *Repository) GetAll() []*Consignment {
	return repo.consignments
}
