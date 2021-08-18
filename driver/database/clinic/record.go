package clinic

import (
	"main-backend/bussiness/clinic"
	"main-backend/driver/database/city"
	"main-backend/driver/database/user"
	"time"
)

type Clinic struct {
	ID          int `gorm:"primaryKey"`
	UserID      int
	User        *user.User `gorm:"foreignKey:UserID;references:ID"`
	CityID      int
	City        *city.City `gorm:"foreignKey:CityID;references:ID"`
	Name        string
	Description string
	OpenTime    string
	CloseTime   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func fromDomain(domain *clinic.Domain) *Clinic {
	return &Clinic{
		ID:          domain.ID,
		UserID:      domain.UserID,
		CityID:      domain.CityID,
		Name:        domain.Name,
		Description: domain.Description,
		OpenTime:    domain.OpenTime,
		CloseTime:   domain.CloseTime,
	}
}

func (rec *Clinic) ToDomain() (res *clinic.Domain) {
	if rec != nil {
		res = &clinic.Domain{
			ID:          rec.ID,
			UserID:      rec.UserID,
			User:        rec.User.ToDomain(),
			CityID:      rec.CityID,
			City:        rec.City.ToDomain(),
			Name:        rec.Name,
			Description: rec.Description,
			OpenTime:    rec.OpenTime,
			CloseTime:   rec.CloseTime,
			CreatedAt:   rec.CreatedAt,
			UpdatedAt:   rec.UpdatedAt,
		}
	}
	return res
}
