package service

import (
	"auction/internal/repository"
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
	Repos  *repository.Repositories
	Hasher hash.PasswordHasher
}

func NewService(deps Deps) *Services {
	usersService := NewUsersService(deps.Repos.Users, deps.Hasher)

	return &Services{
		Users: usersService,
	}
}
