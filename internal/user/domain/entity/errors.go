package entity

import "errors"

var (
	ErrEmailIsUsed = errors.New("email is already being used")
)
