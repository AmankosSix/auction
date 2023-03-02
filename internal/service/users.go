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

type UsersService struct {
	repo         repository.Users
	hasher       hash.PasswordHasher
	tokenManager auth.TokenManager

	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

func NewUsersService(repo repository.Users, hasher hash.PasswordHasher, tokenManager auth.TokenManager, accessTokenTTL, refreshTokenTTL time.Duration) *UsersService {
	return &UsersService{
		repo:            repo,
		hasher:          hasher,
		tokenManager:    tokenManager,
		accessTokenTTL:  accessTokenTTL,
		refreshTokenTTL: refreshTokenTTL,
	}
}

type UserSignUpInput struct {
	Name     string
	Email    string
	Phone    string
	Password string
}

type UserSignInInput struct {
	Email    string
	Password string
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

	if err = s.repo.Create(user); err != nil {
		if errors.Is(err, model.ErrUserAlreadyExists) {
			return err
		}

		return err
	}

	return nil
}

func (s *UsersService) SignIn(ctx context.Context, input UserSignInInput) (Tokens, error) {
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

func (s *UsersService) createSession(body model.TokenBody) (Tokens, error) {
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

func (s *UsersService) UserInfo(uuid string) (model.UserInfo, error) {
	user, err := s.repo.GetByUUID(uuid)
	if err != nil {
		return model.UserInfo{}, err
	}

	return user, nil
}

func (s *UsersService) UserUpdateInfo(uuid string, input model.UpdateUserInfoInput) error {
	return s.repo.UpdateUserInfo(uuid, input)
}
