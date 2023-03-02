package model

import "time"

type Staff struct {
	UUID         string    `json:"uuid" db:"uuid"`
	Name         string    `json:"name" binding: "required"`
	Password     string    `json:"password_hash" binding: "required"`
	Email        string    `json:"email" binding: "required"`
	Phone        string    `json:"phone" binding: "required"`
	RegisteredAt time.Time `json:"registered_at" binding: "required"`
	LastVisitAt  time.Time `json:"last_visit_at" binding: "required"`
	RoleUUID     string    `json:"role_uuid" binding: "required"`
}

type StaffInfo struct {
	UUID  string `json:"uuid" db:"uuid"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Role  string `json:"role"`
}

type UpdateStaffInfoInput struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}
