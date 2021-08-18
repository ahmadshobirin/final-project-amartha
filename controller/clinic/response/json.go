package response

import (
	"main-backend/bussiness/clinic"
	cityResp "main-backend/controller/city/response"
	userResp "main-backend/controller/user/response"
	"time"
)

type Clinic struct {
	ID          int            `json:"id"`
	UserID      int            `json:"user_id"`
	User        *userResp.User `json:"user"`
	CityID      int            `json:"city_id"`
	City        *cityResp.City `json:"city"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	OpenTime    string         `json:"open_time"`
	CloseTime   string         `json:"close_time"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

func FromDomain(domain *clinic.Domain) (res *Clinic) {
	if domain != nil {
		res = &Clinic{
			ID:          domain.ID,
			UserID:      domain.UserID,
			User:        userResp.FromDomain(domain.User),
			CityID:      domain.CityID,
			City:        cityResp.FromDomain(domain.City),
			Name:        domain.Name,
			Description: domain.Description,
			OpenTime:    domain.OpenTime,
			CloseTime:   domain.CloseTime,
			CreatedAt:   domain.CreatedAt,
			UpdatedAt:   domain.UpdatedAt,
		}
	}
	return res
}
