package city

import (
	"context"
	"time"
)

type Domain struct {
	ID        int
	Code      string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	Find(ctx context.Context) ([]Domain, error)
	FindByID(ctx context.Context, id int) (Domain, error)
}

type Repository interface {
	Find(ctx context.Context) ([]Domain, error)
	FindByID(ctx context.Context, id int) (Domain, error)
}
