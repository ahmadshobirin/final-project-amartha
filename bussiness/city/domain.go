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
	GetAll(ctx context.Context) ([]Domain, error)
	FindByID(ctx context.Context, id int) (Domain, error)
	Store(ctx context.Context, data *Domain) error
	Update(ctx context.Context, data *Domain) error
	Delete(ctx context.Context, data *Domain) error
}

type Repository interface {
	GetAll(ctx context.Context) ([]Domain, error)
	FindByID(ctx context.Context, id int) (Domain, error)
	GetByName(ctx context.Context, name string) (Domain, error)
	Store(ctx context.Context, data *Domain) error
	Update(ctx context.Context, data *Domain) error
	Delete(ctx context.Context, data *Domain) error
}
