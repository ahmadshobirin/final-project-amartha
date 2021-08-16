package transaction

import (
	"main-backend/driver/database/clinic"
	"main-backend/driver/database/user"
	"time"
)

type Transaction struct {
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
