package repository

import (
	"auction/internal/model"
	"auction/pkg/database"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type UsersRepo struct {
	db *sqlx.DB
}

func NewUsersRepo(db *sqlx.DB) *UsersRepo {
	return &UsersRepo{
		db: db,
	}
}

func (r *UsersRepo) Create(ctx context.Context, user model.User) error {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, password_hash, email, phone, registered_at, last_visit_at) values ($1, $2, $3, $4, $5, $6) RETURNING id", database.UsersTable)
	row := r.db.QueryRow(query, user.Name, user.Password, user.Email, user.Phone, user.RegisteredAt, user.LastVisitAt)
	if err := row.Scan(&id); err != nil {
		return err
	}
	logrus.Info("Aman's id: ", id)
	return nil
}
