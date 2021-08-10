package response

import (
	"main-backend/bussiness/city"
)

type City struct {
	ID        int    `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func FromDomain(domain city.Domain) City {
	return City{
		ID:        domain.ID,
		Code:      domain.Code,
		Name:      domain.Name,
		CreatedAt: domain.CreatedAt.Format("2006-01-01 15:04:05"),
		UpdatedAt: domain.UpdatedAt.Format("2006-01-01 15:04:05"),
	}
}
