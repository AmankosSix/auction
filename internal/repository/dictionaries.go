package repository

import (
	"auction/internal/model"
	"auction/pkg/database"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type DictionariesRepo struct {
	db *sqlx.DB
}

func NewDictionariesRepo(db *sqlx.DB) *OwnerRepo {
	return &OwnerRepo{
		db: db,
	}
}

func (r *OwnerRepo) GetAllRoles() ([]model.Role, error) {
	var roles []model.Role
	query := fmt.Sprintf("SELECT * FROM %s", database.RolesTable)

	if err := r.db.Select(&roles, query); err != nil {
		return []model.Role{}, err
	}

	return roles, nil
}
