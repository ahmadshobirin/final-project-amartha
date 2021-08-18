package response

import (
	"main-backend/bussiness/queue"
	clinicResp "main-backend/controller/clinic/response"
	userResp "main-backend/controller/user/response"
	"time"
)

type Queue struct {
	ID        int                `json:"id"`
	UserID    int                `json:"user_id"`
	User      *userResp.User     `json:"user"`
	ClinicID  int                `json:"clinic_id"`
	Clinic    *clinicResp.Clinic `json:"clinic"`
	Date      string             `json:"date"`
	Price     float64            `json:"price"`
	Status    string             `json:"status"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

func FromDomain(domain *queue.Domain) (res *Queue) {
	if domain != nil {
		res = &Queue{
			ID:        domain.ID,
			UserID:    domain.UserID,
			User:      userResp.FromDomain(domain.User),
			ClinicID:  domain.ClinicID,
			Clinic:    clinicResp.FromDomain(domain.Clinic),
			Date:      domain.Date,
			Price:     domain.Price,
			Status:    domain.Status,
			CreatedAt: domain.CreatedAt,
			UpdatedAt: domain.UpdatedAt,
		}
	}
	return res
}
