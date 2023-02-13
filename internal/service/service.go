package service

import (
	"auction/pkg/hash"
	"context"
)

type Users interface {
	SignUp(ctx context.Context, input UserSignUpInput) error
}

type UserSignUpInput struct {
	Name     string
	Email    string
	Phone    string
	Password string
}

type Services struct {
	Users Users
}

type Deps struct {
	Hasher hash.PasswordHasher
}

func NewService(deps Deps) *Services {
	usersService := NewUsersService(deps.Hasher)

	return &Services{
		Users: usersService,
	}
}
