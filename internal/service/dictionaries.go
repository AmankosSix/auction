package service

import (
	"auction/internal/model"
	"auction/internal/repository"
	"auction/pkg/auth"
	"auction/pkg/hash"
	"time"
)

type DictionariesService struct {
	repo         repository.Dictionaries
	hasher       hash.PasswordHasher
	tokenManager auth.TokenManager

	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

func NewDictionariesService(repo repository.Dictionaries, hasher hash.PasswordHasher, tokenManager auth.TokenManager, accessTokenTTL, refreshTokenTTL time.Duration) *DictionariesService {
	return &DictionariesService{
		repo:            repo,
		hasher:          hasher,
		tokenManager:    tokenManager,
		accessTokenTTL:  accessTokenTTL,
		refreshTokenTTL: refreshTokenTTL,
	}
}

func (s *DictionariesService) RolesList() ([]model.Role, error) {
	roles, err := s.repo.GetAllRoles()
	if err != nil {
		return []model.Role{}, err
	}

	return roles, nil
}
