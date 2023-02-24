package model

import "time"

type User struct {
	UUID         string    `json:"uuid" db:"uuid"`
	Name         string    `json:"name" binding: "required"`
	Password     string    `json:"password_hash" binding: "required"`
	Email        string    `json:"email" binding: "required"`
	Phone        string    `json:"phone" binding: "required"`
	RegisteredAt time.Time `json:"registered_at" binding: "required"`
	LastVisitAt  time.Time `json:"last_visit_at" binding: "required"`
	RoleUUID     string    `json:"role_uuid" binding: "required"`
}

type UserInfo struct {
	UUID  string `json:"uuid" db:"uuid"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	//RegisteredAt time.Time `json:"registered_at"`
	//LastVisitAt  time.Time `json:"last_visit_at"`
	Role string `json:"role"`
}
