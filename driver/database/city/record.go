package city

import (
	"main-backend/bussiness/city"
	"time"
)

type City struct {
	ID        int
	Code      string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func fromDomain(domain *city.Domain) *City {
	return &City{
		ID:   domain.ID,
		Code: domain.Code,
		Name: domain.Name,
	}
}

func (rec *City) ToDomain() city.Domain {
	return city.Domain{
		ID:        rec.ID,
		Code:      rec.Code,
		Name:      rec.Name,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}
