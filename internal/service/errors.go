package service

import "errors"

// Enumerate service errors

var (
	ErrUserNotFound     error = errors.New("user not found")
	ErrUserIDNotAllowed error = errors.New("we should not accept requests from this user due to suspicious activity")
)
