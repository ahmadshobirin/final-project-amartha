package response

import (
	"main-backend/bussiness/clinic"
	"time"
)

type Clinic struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	User        string    `json:"user"`
	CityID      int       `json:"city_id"`
	City        string    `json:"city"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	OpenTime    string    `json:"open_time"`
	CloseTime   string    `json:"close_time"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FromDomain(domain clinic.Domain) Clinic {
	return Clinic{
		ID:          domain.ID,
		UserID:      domain.UserID,
		User:        domain.User,
		CityID:      domain.CityID,
		City:        domain.City,
		Name:        domain.Name,
		Description: domain.Description,
		OpenTime:    domain.OpenTime,
		CloseTime:   domain.CloseTime,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}
