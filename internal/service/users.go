package service

import (
	"auction/internal/model"
	"auction/internal/repository"
	"auction/pkg/hash"
	"context"
	"errors"
	"time"
)

type UsersService struct {
	repo   repository.Users
	hasher hash.PasswordHasher
}

func NewUsersService(repo repository.Users, hasher hash.PasswordHasher) *UsersService {
	return &UsersService{
		repo:   repo,
		hasher: hasher,
	}
}

func (s *UsersService) SignUp(ctx context.Context, input UserSignUpInput) error {
	passwordHash, err := s.hasher.Hash(input.Password)
	if err != nil {
		return err
	}

	user := model.User{
		Name:         input.Name,
		Password:     passwordHash,
		Phone:        input.Phone,
		Email:        input.Email,
		RegisteredAt: time.Now(),
		LastVisitAt:  time.Now(),
	}

	if err = s.repo.Create(ctx, user); err != nil {
		if errors.Is(err, model.ErrUserAlreadyExists) {
			return err
		}

		return err
	}

	return nil
}
