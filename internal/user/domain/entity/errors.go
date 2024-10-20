package entity

import "errors"

var (
	ErrEmailIsUsed       = errors.New("email is already being used")
	ErrNotRegistered     = errors.New("email is not registered")
	ErrIncorrectPassword = errors.New("incorrect password")
	ErrUserNotFound      = errors.New("user is not found")
)
