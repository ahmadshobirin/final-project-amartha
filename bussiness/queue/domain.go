package queue

import (
	"context"
	"main-backend/driver/database/clinic"
	"main-backend/driver/database/user"
	"time"
)

type Domain struct {
	ID        int
	ClinicID  int
	Clinic    clinic.Clinic
	UserID    int
	User      user.User
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
	FindByID(ctx context.Context, ID int) (Domain, error)
	FindOne(ctx context.Context, userID, clinicID int, status string) (Domain, error)
	Store(ctx context.Context, userID int, data *Domain) (err error)
	Update(ctx context.Context, data *Domain) (err error)
}

type Repository interface {
	FindByID(ctx context.Context, ID int) (Domain, error)
	FindOne(ctx context.Context, userID, clinicID int, status string) (Domain, error)
	Store(ctx context.Context, data *Domain) (err error)
	Update(ctx context.Context, data *Domain) (err error)
}
