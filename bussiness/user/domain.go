package user

import (
	"context"
	"time"
)

type Domain struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	RoleID    int    `json:"role_id"`
	Role      string `json:"role"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Bio       string `json:"bio"`
	Status    bool   `json:"status"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error)
	FindByID(ctx context.Context, ID int) (Domain, error)
	FindByEmail(ctx context.Context, email string) (Domain, error)
	Store(ctx context.Context, data *Domain, roleID int) (Domain, error)
}

type Repository interface {
	Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error)
	FindByID(ctx context.Context, ID int) (Domain, error)
	FindByEmail(ctx context.Context, email string) (Domain, error)
	Store(ctx context.Context, data *Domain) (Domain, error)
}
