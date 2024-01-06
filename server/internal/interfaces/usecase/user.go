package usecase

import (
	"context"
)

type UserInterface interface{
	CreateUser(c context.Context, req *domain.User) (*domain.User, error)
	Login(c context.Context, req *domain.User) (*domain.User, error)
}

type User struct {
	ID       int64
	Username string
	Email    string
	Password string
}
