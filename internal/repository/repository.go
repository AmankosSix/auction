package repository

import (
	"auction/internal/model"
	"github.com/jmoiron/sqlx"
)

type Users interface {
	Create(user model.User) error
	GetByCredentials(email, password string) (model.TokenBody, error)
	GetByUUID(uuid string) (model.UserInfo, error)
	UpdateUserInfo(uuid string, input model.UpdateUserInfoInput) error
	SetSession(uuid string, session model.Session) error
}

type Staff interface {
	GetByCredentials(email, password string) (model.TokenBody, error)
	GetByUUID(uuid string) (model.StaffInfo, error)
	UpdateStaffInfo(uuid string, input model.UpdateStaffInfoInput) error
	SetSession(uuid string, session model.Session) error
}

type Owner interface {
	Create(staff model.Staff) error
	GetAllStaff() ([]model.StaffInfo, error)
	RemoveStaff(uuid string) error
}

type Dictionaries interface {
	GetAllRoles() ([]model.Role, error)
}

type Repositories struct {
	Users        Users
	Staff        Staff
	Owner        Owner
	Dictionaries Dictionaries
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		Users:        NewUsersRepo(db),
		Staff:        NewStaffRepo(db),
		Owner:        NewOwnerRepo(db),
		Dictionaries: NewDictionariesRepo(db),
	}
}
