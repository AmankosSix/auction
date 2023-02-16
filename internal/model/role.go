package model

type Role struct {
	ID   int    `json:"id" binding:"required"`
	Role string `json:"role" binding:"required"`
}
