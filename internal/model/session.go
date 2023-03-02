package model

import "time"

type Session struct {
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
}

type TokenBody struct {
	Uuid string `json:"uuid"`
	Role string `json:"role"`
}
