package service

import (
	"auction/internal/model"
	"auction/internal/repository"
	"auction/pkg/auth"
	"auction/pkg/hash"
	"context"
	"errors"
	"time"
)

type OwnerService struct {
	repo         repository.Owner
	hasher       hash.PasswordHasher
	tokenManager auth.TokenManager

	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

func NewOwnerService(repo repository.Owner, hasher hash.PasswordHasher, tokenManager auth.TokenManager, accessTokenTTL, refreshTokenTTL time.Duration) *OwnerService {
	return &OwnerService{
		repo:            repo,
		hasher:          hasher,
		tokenManager:    tokenManager,
		accessTokenTTL:  accessTokenTTL,
		refreshTokenTTL: refreshTokenTTL,
	}
}

type OwnerSignUpInput struct {
	Name     string
	Email    string
	Phone    string
	Password string
}

func (s *OwnerService) SignUp(ctx context.Context, input OwnerSignUpInput) error {
	passwordHash, err := s.hasher.Hash(input.Password)
	if err != nil {
		return err
	}

	owner := model.Staff{
		Name:         input.Name,
		Password:     passwordHash,
		Phone:        input.Phone,
		Email:        input.Email,
		RegisteredAt: time.Now(),
		LastVisitAt:  time.Now(),
	}

	if err = s.repo.Create(owner); err != nil {
		if errors.Is(err, model.ErrStaffAlreadyExists) {
			return err
		}

		return err
	}

	return nil
}

func (s *OwnerService) StaffList() ([]model.StaffInfo, error) {
	staff, err := s.repo.GetAllStaff()
	if err != nil {
		return []model.StaffInfo{}, err
	}

	return staff, nil
}
