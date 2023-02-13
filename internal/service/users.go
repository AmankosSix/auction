package service

import (
	"auction/internal/model"
	"auction/pkg/hash"
	"context"
	"github.com/sirupsen/logrus"
	"time"
)

type UsersService struct {
	hasher hash.PasswordHasher
}

func NewUsersService(hasher hash.PasswordHasher) *UsersService {
	return &UsersService{
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

	logrus.Info(user)

	return nil
}
