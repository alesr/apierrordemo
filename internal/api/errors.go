package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/alesr/apierrordemo/internal/service"
)

type TransportError struct {
	Code       string // Error the frontend will use to map to an user friendly message (can be translated)
	Message    string // Help the frontend dev to debug the application
	StatusCode int    // Nice to have...
}

// Implement error interface
func (e TransportError) Error() string {
	return fmt.Sprintf("code: %s, message: %s, status_code: %d",
		e.Code, e.Message, e.StatusCode)
}

// Enumerate transport errors

var (
	ErrInvalidUserID = TransportError{
		Code:       "ErrInvalidUserID",
		Message:    "User ID must be a 3 digit string",
		StatusCode: http.StatusBadRequest,
	}

	ErrUserNotFound = TransportError{
		Code:       "ErrUserNotFound",
		Message:    "Could not find user",
		StatusCode: http.StatusNotFound,
	}

	ErrUserIDNotAllowed = TransportError{
		Code:       "ErrUserIDNotAllowed",
		Message:    "user id not allowed due to suspiscious activity",
		StatusCode: http.StatusUnauthorized,
	}

	ErrInternal = TransportError{
		Code:       "ErrInternal",
		Message:    "Internal Server Error",
		StatusCode: http.StatusInternalServerError,
	}
)

// Translate service to transport errors
func transportError(err error) TransportError {
	switch {
	case errors.Is(service.ErrUserNotFound, err):
		return ErrUserNotFound
	case errors.Is(service.ErrUserIDNotAllowed, err):
		return ErrUserIDNotAllowed
	default:
		return ErrInternal
	}
}
