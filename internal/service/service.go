package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/alesr/apierrordemo/internal/repository"
)

const suspiciousUser string = "456"

type repo interface {
	GetUserByID(ctx context.Context, id string) (*repository.User, error)
}

type DefaultService struct {
	repo repo
}

func NewDefaultService(repo repo) *DefaultService {
	return &DefaultService{
		repo: repo,
	}
}

func (s *DefaultService) FetchUser(ctx context.Context, id string) (*User, error) {
	// Semantic check
	if id == suspiciousUser {
		return nil, ErrUserIDNotAllowed
	}

	storedUser, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("could not get user by id: %w", err)
	}

	return &User{
		Name:  storedUser.Name,
		Email: storedUser.Email,
	}, nil
}
