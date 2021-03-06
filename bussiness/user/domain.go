package user

import (
	"context"
	"main-backend/bussiness/role"
	"time"
)

type Domain struct {
	ID        int
	RoleID    int
	Role      *role.Domain
	Name      string
	Email     string
	Password  string
	Bio       string
	Status    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	Fetch(ctx context.Context, roleCode string, page, perpage int) ([]Domain, int, error)
	FindByID(ctx context.Context, ID int) (Domain, error)
	FindByEmail(ctx context.Context, email string) (Domain, error)
	Store(ctx context.Context, data *Domain, roleID int) (Domain, error)
	Update(ctx context.Context, data *Domain) error
}

type Repository interface {
	Fetch(ctx context.Context, roleCode string, page, perpage int) ([]Domain, int, error)
	FindByID(ctx context.Context, ID int) (Domain, error)
	FindByEmail(ctx context.Context, email string) (Domain, error)
	Store(ctx context.Context, data *Domain) (Domain, error)
	Update(ctx context.Context, data *Domain) error
}
