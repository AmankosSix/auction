package model

import "time"

type User struct {
	ID           int       `json:"id" db:"id"`
	Name         string    `json:"name" binding: "required"`
	Password     string    `json:"password_hash" binding: "required"`
	Email        string    `json:"email" binding: "required"`
	Phone        string    `json:"phone" binding: "required"`
	RegisteredAt time.Time `json:"registered_at" binding: "required"`
	LastVisitAt  time.Time `json:"last_visit_at" binding: "required"`
}
