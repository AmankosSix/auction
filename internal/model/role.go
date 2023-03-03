package model

type Role struct {
	UUID string `json:"uuid" binding:"required"`
	Role string `json:"role" binding:"required"`
}
