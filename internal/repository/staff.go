package repository

import (
	"auction/internal/model"
	"auction/pkg/database"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

type StaffRepo struct {
	db *sqlx.DB
}

func NewStaffRepo(db *sqlx.DB) *StaffRepo {
	return &StaffRepo{
		db: db,
	}
}

func (r *StaffRepo) getRole(role string) (string, error) {
	var uuid string
	query := fmt.Sprintf("SELECT uuid FROM %s WHERE role=$1", database.RolesTable)

	if err := r.db.Get(&uuid, query, role); err != nil {
		return "", model.ErrRoleNotFound
	}
	return uuid, nil
}

func (r *StaffRepo) GetByCredentials(email, password string) (model.TokenBody, error) {
	var res model.TokenBody
	query := fmt.Sprintf("SELECT s.uuid, role FROM %s s INNER JOIN %s r ON s.role_uuid = r.uuid WHERE s.email=$1 AND s.password_hash=$2", database.StaffTable, database.RolesTable)
	if err := r.db.Get(&res, query, email, password); err != nil {
		return model.TokenBody{}, model.ErrStaffNotFound
	}

	return res, nil
}

func (r *StaffRepo) SetSession(uuid string, session model.Session) error {
	query := fmt.Sprintf("UPDATE %s SET last_visit_at=$1 WHERE uuid = $2", database.StaffTable)
	_, err := r.db.Exec(query, time.Now(), uuid)

	return err
}

func (r *StaffRepo) GetByUUID(uuid string) (model.StaffInfo, error) {
	var staff model.StaffInfo
	query := fmt.Sprintf("SELECT s.uuid, name, email, phone, role FROM %s s INNER JOIN %s r ON r.uuid = s.role_uuid WHERE s.uuid = $1", database.StaffTable, database.RolesTable)
	if err := r.db.Get(&staff, query, uuid); err != nil {
		return model.StaffInfo{}, model.ErrStaffNotFound
	}

	return staff, nil
}

func (r *StaffRepo) UpdateStaffInfo(uuid string, input model.UpdateStaffInfoInput) error {
	query := fmt.Sprintf("UPDATE %s SET name=$1, phone=$2 WHERE uuid = $3", database.StaffTable)
	_, err := r.db.Exec(query, input.Name, &input.Phone, uuid)

	return err
}
