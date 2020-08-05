package main

import (
	ub "github.com/kramanathan01/ship/ship-service-user/proto/user"
)

// MarshalUser -
func MarshalUser(user *ub.User) *User {
	return &User{
		ID:       user.Id,
		Name:     user.Name,
		Company:  user.Company,
		Email:    user.Email,
		Password: user.Password,
	}
}

// MarshalUserCollection -
func MarshalUserCollection(users []*ub.User) []*User {
	collection := make([]*User, len(users))
	for _, user := range users {
		collection = append(collection, MarshalUser(user))
	}
	return collection
}

// UnMarshalUser -
func UnMarshalUser(user *User) *ub.User {
	return &ub.User{
		Id:       user.ID,
		Name:     user.Name,
		Company:  user.Company,
		Email:    user.Email,
		Password: user.Password,
	}
}

// UnMarshalUserCollection -
func UnMarshalUserCollection(users []*User) []*ub.User {
	collection := make([]*ub.User, len(users))
	for _, user := range users {
		collection = append(collection, UnMarshalUser(user))
	}
	return collection
}
