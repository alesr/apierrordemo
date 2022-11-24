package repository

import "errors"

// Enumerate repository errors

var (
	ErrUserNotFound error = errors.New("user not found")
)
