package repository

import (
	"auction/internal/model"
	"github.com/jmoiron/sqlx"
)

type Users interface {
	Create(user model.User) error
	GetByCredentials(email, password string) (string, error)
	GetByUUID(uuid string) (model.UserInfo, error)
	SetSession(uuid string, session model.Session) error
}

type Repositories struct {
	Users Users
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		Users: NewUsersRepo(db),
	}
}
