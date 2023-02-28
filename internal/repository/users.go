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
	roleUUID, err := r.getRole("user")
	if err != nil {
		return err
	}
	query := fmt.Sprintf("INSERT INTO %s (name, password_hash, email, phone, registered_at, last_visit_at, role_uuid) values ($1, $2, $3, $4, $5, $6, $7) RETURNING uuid", database.UsersTable)
	row := r.db.QueryRow(query, user.Name, user.Password, user.Email, user.Phone, user.RegisteredAt, user.LastVisitAt, roleUUID)
	if err := row.Scan(&uuid); err != nil {
		return err
	}

	return nil
}

func (r *UsersRepo) getRole(role string) (string, error) {
	var uuid string
	query := fmt.Sprintf("SELECT uuid FROM %s WHERE role=$1", database.RolesTable)

	if err := r.db.Get(&uuid, query, role); err != nil {
		return "", model.ErrRoleNotFound
	}
	return uuid, nil
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

func (r *UsersRepo) GetByUUID(uuid string) (model.UserInfo, error) {
	var user model.UserInfo
	query := fmt.Sprintf("SELECT u.uuid, name, email, phone, role FROM %s u INNER JOIN %s r ON r.uuid = u.role_uuid WHERE u.uuid = $1", database.UsersTable, database.RolesTable)
	if err := r.db.Get(&user, query, uuid); err != nil {
		return model.UserInfo{}, model.ErrUserNotFound
	}

	return user, nil
}

func (r *UsersRepo) UpdateUserInfo(uuid string, input model.UpdateUserInfoInput) error {
	query := fmt.Sprintf("UPDATE %s SET name=$1, phone=$2 WHERE uuid = $3", database.UsersTable)
	_, err := r.db.Exec(query, input.Name, &input.Phone, uuid)

	return err
}
