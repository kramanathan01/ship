package main

import (
	"context"

	ub "github.com/kramanathan01/ship/ship-service-user/proto/user"
	"golang.org/x/crypto/bcrypt"
)

type handler struct {
	repository
}

// Create -
func (h *handler) Create(ctx context.Context, req *ub.User, res *ub.Response) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	req.Password = string(hashedPass)
	if err := h.repository.Create(ctx, MarshalUser(req)); err != nil {
		return err
	}

	// Strip the password back out, so's we're not returning it
	req.Password = ""
	res.User = req
	return nil
}

// Get -
func (h *handler) Get(ctx context.Context, req *ub.User, res *ub.Response) error {
	u, err := h.repository.GetByID(ctx, req.Id)
	if err != nil {
		return err
	}

	res.User = UnMarshalUser(u)
	return nil
}

// GetAll -
func (h *handler) GetAll(ctx context.Context, req *ub.Request, res *ub.Response) error {
	users := make([]*User, 0)
	users, err := h.repository.GetAll(ctx)
	if err != nil {
		return err
	}
	res.Users = UnMarshalUserCollection(users)
	return nil
}

// Auth - Placeholder
func (h *handler) Auth(ctx context.Context, req *ub.User, res *ub.Token) error {
	res.Token = ""
	return nil
}

// ValidateToken -- Placeholder
func (h *handler) ValidateToken(ctx context.Context, req *ub.Token, res *ub.Token) error {
	res.Token = req.Token
	return nil
}
