package repository

import (
	"auction/internal/model"
	"auction/pkg/database"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

type UsersRepo struct {
	db *sqlx.DB
}

func NewUsersRepo(db *sqlx.DB) *UsersRepo {
	return &UsersRepo{
		db: db,
	}
}

func (r *UsersRepo) Create(user model.User) error {
	var uuid string
	query := fmt.Sprintf("INSERT INTO %s (name, password_hash, email, phone, registered_at, last_visit_at) values ($1, $2, $3, $4, $5, $6) RETURNING uuid", database.UsersTable)
	row := r.db.QueryRow(query, user.Name, user.Password, user.Email, user.Phone, user.RegisteredAt, user.LastVisitAt)
	if err := row.Scan(&uuid); err != nil {
		return err
	}

	return nil
}

func (r *UsersRepo) GetByCredentials(email, password string) (string, error) {
	var uuid string
	query := fmt.Sprintf("SELECT uuid FROM %s WHERE email=$1 AND password_hash=$2", database.UsersTable)
	if err := r.db.Get(&uuid, query, email, password); err != nil {
		return "", model.ErrUserNotFound
	}

	return uuid, nil
}

func (r *UsersRepo) SetSession(uuid string, session model.Session) error {
	query := fmt.Sprintf("UPDATE %s SET last_visit_at=$1 WHERE uuid = $2", database.UsersTable)
	_, err := r.db.Exec(query, time.Now(), uuid)

	return err
}
