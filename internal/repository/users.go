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
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, password_hash, email, phone, registered_at, last_visit_at) values ($1, $2, $3, $4, $5, $6) RETURNING id", database.UsersTable)
	row := r.db.QueryRow(query, user.Name, user.Password, user.Email, user.Phone, user.RegisteredAt, user.LastVisitAt)
	if err := row.Scan(&id); err != nil {
		return err
	}

	return nil
}

func (r *UsersRepo) GetByCredentials(email, password string) (int, error) {
	var id int
	query := fmt.Sprintf("SELECT id FROM %s WHERE email=$1 AND password_hash=$2", database.UsersTable)
	if err := r.db.Get(&id, query, email, password); err != nil {
		return 0, model.ErrUserNotFound
	}

	return id, nil
}

func (r *UsersRepo) SetSession(userID int, session model.Session) error {
	query := fmt.Sprintf("UPDATE %s SET last_visit_at=$1 WHERE id = $2", database.UsersTable)
	_, err := r.db.Exec(query, time.Now(), userID)

	return err
}
