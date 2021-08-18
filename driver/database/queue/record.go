package queue

import (
	"main-backend/bussiness/queue"
	"main-backend/driver/database/clinic"
	"main-backend/driver/database/user"
	"time"
)

type Queue struct {
	ID        int
	ClinicID  int `gorm:"foreignKey:ClinicID;references:ID"`
	Clinic    *clinic.Clinic
	UserID    int
	User      *user.User `gorm:"foreignKey:UserID;references:ID"`
	Date      string
	Price     float64
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func fromDomain(domain *queue.Domain) *Queue {
	return &Queue{
		ID:       domain.ID,
		UserID:   domain.UserID,
		ClinicID: domain.ClinicID,
		Date:     domain.Date,
		Price:    domain.Price,
		Status:   domain.Status,
	}
}

func (rec *Queue) ToDomain() (res *queue.Domain) {
	if rec != nil {
		res = &queue.Domain{
			ID:        rec.ID,
			UserID:    rec.UserID,
			User:      rec.User.ToDomain(),
			ClinicID:  rec.ClinicID,
			Clinic:    rec.Clinic.ToDomain(),
			Date:      rec.Date,
			Price:     rec.Price,
			Status:    rec.Status,
			CreatedAt: rec.CreatedAt,
			UpdatedAt: rec.UpdatedAt,
		}
	}
	return res
}
