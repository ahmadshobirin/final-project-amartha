package request

import "main-backend/bussiness/clinic"

type Clinic struct {
	UserID      int    `json:"user_id"`
	CityID      int    `json:"city_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	OpenTime    string `json:"open_time"`
	CloseTime   string `json:"close_time"`
}

func (req *Clinic) ToDomain() *clinic.Domain {
	return &clinic.Domain{
		UserID:      req.UserID,
		CityID:      req.CityID,
		Name:        req.Name,
		Description: req.Description,
		OpenTime:    req.OpenTime,
		CloseTime:   req.CloseTime,
	}
}
