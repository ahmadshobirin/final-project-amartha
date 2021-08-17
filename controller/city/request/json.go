package request

import "main-backend/bussiness/city"

type City struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func (req *City) ToDomain() *city.Domain {
	return &city.Domain{
		Code: req.Code,
		Name: req.Name,
	}
}
