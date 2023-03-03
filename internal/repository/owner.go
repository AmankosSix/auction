package repository

import (
	"auction/internal/model"
	"auction/pkg/database"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type OwnerRepo struct {
	db *sqlx.DB
}

func NewOwnerRepo(db *sqlx.DB) *OwnerRepo {
	return &OwnerRepo{
		db: db,
	}
}

func (r *OwnerRepo) Create(owner model.Staff) error {
	var uuid string
	roleUUID, err := r.getRole("owner")
	if err != nil {
		return err
	}
	query := fmt.Sprintf("INSERT INTO %s (name, password_hash, email, phone, registered_at, last_visit_at, role_uuid) values ($1, $2, $3, $4, $5, $6, $7) RETURNING uuid", database.StaffTable)
	row := r.db.QueryRow(query, owner.Name, owner.Password, owner.Email, owner.Phone, owner.RegisteredAt, owner.LastVisitAt, roleUUID)
	if err := row.Scan(&uuid); err != nil {
		return err
	}

	return nil
}

func (r *OwnerRepo) getRole(role string) (string, error) {
	var uuid string
	query := fmt.Sprintf("SELECT uuid FROM %s WHERE role=$1", database.RolesTable)

	if err := r.db.Get(&uuid, query, role); err != nil {
		return "", model.ErrRoleNotFound
	}
	return uuid, nil
}

func (r *OwnerRepo) GetAllStaff() ([]model.StaffInfo, error) {
	var staff []model.StaffInfo
	query := fmt.Sprintf("SELECT s.uuid, name, email, phone, role FROM %s s INNER JOIN %s r ON r.uuid = s.role_uuid WHERE r.role = 'staff'", database.StaffTable, database.RolesTable)
	if err := r.db.Select(&staff, query); err != nil {
		return []model.StaffInfo{}, model.ErrStaffNotFound
	}

	return staff, nil
}

func (r *OwnerRepo) RemoveStaff(uuid string) error {
	logrus.Info(uuid)
	query := fmt.Sprintf("DELETE FROM %s WHERE uuid = $1", database.StaffTable)
	if _, err := r.db.Exec(query, uuid); err != nil {
		return err
	}

	return nil
}
