package users

import (
	"context"
	"time"
)

type Domain struct {
	ID        int
	Username  string
	Password  string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UseCase interface {
	Register(ctx context.Context, data *Domain) error
}

type Repository interface {
	Insert(ctx context.Context, data *Domain) error
	GetByUsername(ctx context.Context, username string) (Domain, error)
}
