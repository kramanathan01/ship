package main

import "context"

// User - Struct to hold users
type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Company  string `json:"company"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Token -- Struct to hold Token
type Token struct {
	Token string `json:"token"`
	Valid bool   `json:"valid"`
}

type repository interface {
	Create(ctx context.Context, user *User) error
	GetByID(ctx context.Context, id string) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetAll(ctx context.Context) ([]*User, error)
}
