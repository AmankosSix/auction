package service

import (
	"auction/internal/model"
	"auction/internal/repository"
	"auction/pkg/auth"
	"auction/pkg/hash"
	"context"
	"time"
)

type Users interface {
	SignUp(ctx context.Context, input UserSignUpInput) error
	SignIn(ctx context.Context, input UserSignInInput) (Tokens, error)
	UserInfo(uuid string) (model.UserInfo, error)
	UserUpdateInfo(uuid string, input model.UpdateUserInfoInput) error
}

type Staff interface {
	SignIn(ctx context.Context, input StaffSignInInput) (Tokens, error)
	StaffInfo(uuid string) (model.StaffInfo, error)
	StaffUpdateInfo(uuid string, input model.UpdateStaffInfoInput) error
}

type Owner interface {
	SignUp(ctx context.Context, input OwnerSignUpInput) error
	StaffList() ([]model.StaffInfo, error)
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type Services struct {
	Users Users
	Staff Staff
	Owner Owner
}

type Deps struct {
	Repos           *repository.Repositories
	Hasher          hash.PasswordHasher
	TokenManager    auth.TokenManager
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
}

func NewService(deps Deps) *Services {
	usersService := NewUsersService(deps.Repos.Users, deps.Hasher, deps.TokenManager, deps.AccessTokenTTL, deps.RefreshTokenTTL)
	staffService := NewStaffService(deps.Repos.Staff, deps.Hasher, deps.TokenManager, deps.AccessTokenTTL, deps.RefreshTokenTTL)
	ownerService := NewOwnerService(deps.Repos.Owner, deps.Hasher, deps.TokenManager, deps.AccessTokenTTL, deps.RefreshTokenTTL)

	return &Services{
		Users: usersService,
		Staff: staffService,
		Owner: ownerService,
	}
}
