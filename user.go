package main

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id          uuid.UUID `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"-"`
	Password    string    `json:"-"`
	DateCreated time.Time `json:"dateCreated"`
	LastUpdated time.Time `json:"-"`
}

type UserRepository interface {
	FindUserById(ctx context.Context, id uuid.UUID) (*User, error)
	FindUserByUsername(ctx context.Context, username string) (*User, error)
	CreateUser(ctx context.Context, user *User) error
}
