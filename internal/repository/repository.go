package repository

import (
	"auction/internal/model"
	"context"
	"github.com/jmoiron/sqlx"
)

type Users interface {
	Create(ctx context.Context, user model.User) error
}

type Repositories struct {
	Users Users
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		Users: NewUsersRepo(db),
	}
}
