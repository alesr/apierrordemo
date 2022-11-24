package repository

import (
	"context"
)

type FakePostgres struct{}

func NewFakePostgres() *FakePostgres {
	return &FakePostgres{}
}

func (p *FakePostgres) GetUserByID(ctx context.Context, id string) (*User, error) {
	if id != "123" {
		return nil, ErrUserNotFound
	}

	return &User{
		Name:  "Joe Doe",
		Email: "joedoe@foobar.quz",
	}, nil
}
