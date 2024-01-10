package domain

import (
	"context"
)

type User struct {
	ID       int64
	Username string
	Email    string
	Password string
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}

type UserUsecase interface {
	CreateUser(c context.Context, req *User) (*User, error)
	Login(c context.Context, req *User) (*User, error)
}
