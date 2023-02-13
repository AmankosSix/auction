package model

import "time"

type User struct {
	ID           int       `json:"id" db:"id"`
	Name         string    `json:"name" binding: "required"`
	Password     string    `json:"password" binding: "required"`
	Email        string    `json:"username" binding: "required"`
	Phone        string    `json:"phone" binding: "required"`
	RegisteredAt time.Time `json:"registeredAt" binding: "required"`
	LastVisitAt  time.Time `json:"lastVisitAt" binding: "required"`
}
