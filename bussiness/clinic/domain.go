package clinic

import (
	"context"
	"time"
)

type Domain struct {
	ID          int
	UserID      int
	User        string
	CityID      int
	City        string
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
