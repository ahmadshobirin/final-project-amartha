package queue

import (
	"context"
	"main-backend/bussiness/clinic"
	"main-backend/bussiness/user"
	"time"
)

type Domain struct {
	ID        int
	ClinicID  int
	Clinic    *clinic.Domain
	UserID    int
	User      *user.Domain
	Date      string
	Price     float64
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

const (
	StatusPending = "pending"
	StatusPaid    = "paid"
	StatusFailed  = "failed"
)

type Usecase interface {
	Fetch(ctx context.Context, userID, page, perpage int) ([]Domain, int, error)
	FindByID(ctx context.Context, ID int) (Domain, error)
	FindOne(ctx context.Context, userID, clinicID int, status string) (Domain, error)
	Store(ctx context.Context, userID int, data *Domain) (err error)
	Update(ctx context.Context, adminID int, data *Domain) (err error)
}

type Repository interface {
	FindByID(ctx context.Context, ID int) (Domain, error)
	Fetch(ctx context.Context, userID, page, perpage int) ([]Domain, int, error)
	FindOne(ctx context.Context, userID, clinicID int, status string) (Domain, error)
	Store(ctx context.Context, data *Domain) (err error)
	Update(ctx context.Context, data *Domain) (err error)
}
