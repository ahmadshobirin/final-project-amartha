package response

import (
	"main-backend/bussiness/queue"
	"main-backend/driver/database/clinic"
	"main-backend/driver/database/user"
	"time"
)

type Queue struct {
	ID        int           `json:"id"`
	UserID    int           `json:"user_id"`
	User      user.User     `json:"user"`
	ClinicID  int           `json:"clinic_id"`
	Clinic    clinic.Clinic `json:"clinic"`
	Date      string        `json:"date"`
	Price     float64       `json:"price"`
	Status    string        `json:"status"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

func FromDomain(domain queue.Domain) Queue {
	return Queue{
		ID:        domain.ID,
		UserID:    domain.UserID,
		User:      domain.User,
		ClinicID:  domain.ClinicID,
		Clinic:    domain.Clinic,
		Date:      domain.Date,
		Price:     domain.Price,
		Status:    domain.Status,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
