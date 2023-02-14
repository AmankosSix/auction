package model

import "errors"

var (
	ErrUserAlreadyExists = errors.New("user with such email already exists")
	ErrUserNotFound      = errors.New("user doesn't exists")
)
