package service

import (
	"auction/internal/model"
	"auction/internal/repository"
	"auction/pkg/auth"
	"auction/pkg/hash"
	"context"
	"time"
)

type StaffService struct {
	repo         repository.Staff
	hasher       hash.PasswordHasher
	tokenManager auth.TokenManager

	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

func NewStaffService(repo repository.Staff, hasher hash.PasswordHasher, tokenManager auth.TokenManager, accessTokenTTL, refreshTokenTTL time.Duration) *StaffService {
	return &StaffService{
		repo:            repo,
		hasher:          hasher,
		tokenManager:    tokenManager,
		accessTokenTTL:  accessTokenTTL,
		refreshTokenTTL: refreshTokenTTL,
	}
}

type StaffSignUpInput struct {
	Name     string
	Email    string
	Phone    string
	Password string
}

type StaffSignInInput struct {
	Email    string
	Password string
}

func (s *StaffService) SignIn(ctx context.Context, input StaffSignInInput) (Tokens, error) {
	passwordHash, err := s.hasher.Hash(input.Password)
	if err != nil {
		return Tokens{}, err
	}

	credentials, err := s.repo.GetByCredentials(input.Email, passwordHash)
	if err != nil {
		return Tokens{}, err
	}

	return s.createSession(credentials)
}

func (s *StaffService) createSession(body model.TokenBody) (Tokens, error) {
	var (
		res Tokens
		err error
	)

	res.AccessToken, err = s.tokenManager.NewJWT(body, s.accessTokenTTL)
	if err != nil {
		return res, err
	}

	res.RefreshToken, err = s.tokenManager.NewJWT(body, s.refreshTokenTTL)
	if err != nil {
		return res, err
	}

	session := model.Session{
		RefreshToken: res.RefreshToken,
		ExpiresAt:    time.Now().Add(s.refreshTokenTTL),
	}

	err = s.repo.SetSession(body.Uuid, session)

	return res, err
}

func (s *StaffService) StaffInfo(uuid string) (model.StaffInfo, error) {
	staff, err := s.repo.GetByUUID(uuid)
	if err != nil {
		return model.StaffInfo{}, err
	}

	return staff, nil
}

func (s *StaffService) StaffUpdateInfo(uuid string, input model.UpdateStaffInfoInput) error {
	return s.repo.UpdateStaffInfo(uuid, input)
}
