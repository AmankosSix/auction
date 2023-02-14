package repository

import (
	"auction/internal/model"
	"github.com/jmoiron/sqlx"
)

type Users interface {
	Create(user model.User) error
	GetByCredentials(email, password string) (int, error)
	SetSession(userID int, session model.Session) error
}

type Repositories struct {
	Users Users
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		Users: NewUsersRepo(db),
	}
}
