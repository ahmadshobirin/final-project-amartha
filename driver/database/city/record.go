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

func (rec *City) ToDomain() (res *city.Domain) {
	if rec != nil {
		res = &city.Domain{
			ID:        rec.ID,
			Code:      rec.Code,
			Name:      rec.Name,
			CreatedAt: rec.CreatedAt,
			UpdatedAt: rec.UpdatedAt,
		}
	}
	return res
}
