package entity

import "errors"

var (
	ErrInvalidQuantity = errors.New("purchase quantity cannot be less than 1")
)
