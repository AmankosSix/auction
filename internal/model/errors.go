package model

import "errors"

var (
	ErrUserAlreadyExists  = errors.New("user with such email already exists")
	ErrUserNotFound       = errors.New("user doesn't exists")
	ErrStaffAlreadyExists = errors.New("user with such email already exists")
	ErrStaffNotFound      = errors.New("staff doesn't exists")
	ErrRoleNotFound       = errors.New("role doesn't exists")
)
