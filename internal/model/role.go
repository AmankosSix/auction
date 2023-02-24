package model

type Role struct {
	UUID int    `json:"uuid" binding:"required"`
	Role string `json:"role" binding:"required"`
}
