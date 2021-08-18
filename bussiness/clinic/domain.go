package clinic

import (
	"context"
	"main-backend/bussiness/city"
	"main-backend/bussiness/user"
	"time"
)

type Domain struct {
	ID          int
	UserID      int
	User        *user.Domain
	CityID      int
	City        *city.Domain
	Name        string
	Description string
	OpenTime    string
	CloseTime   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Usecase interface {
	Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error)
	GetByID(ctx context.Context, ID int) (Domain, error)
	GetByUserID(ctx context.Context, userID int) (Domain, error)
	Store(ctx context.Context, data *Domain) error
	Update(ctx context.Context, data *Domain) error
	Delete(ctx context.Context, data *Domain) error
}

type Repository interface {
	Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error)
	GetByID(ctx context.Context, ID int) (Domain, error)
	GetByUserID(ctx context.Context, userID int) (Domain, error)
	Store(ctx context.Context, data *Domain) error
	Update(ctx context.Context, data *Domain) error
	Delete(ctx context.Context, data *Domain) error
}
