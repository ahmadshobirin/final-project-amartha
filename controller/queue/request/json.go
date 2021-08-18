package request

import (
	"main-backend/bussiness/queue"
)

type Queue struct {
	UserID   int     `json:"user_id"`
	ClinicID int     `json:"clinic_Id"`
	Date     string  `json:"date"`
	Price    float64 `json:"price"`
	Status   string  `json:"status"`
}

func (req *Queue) ToDomain() *queue.Domain {
	return &queue.Domain{
		UserID:   req.UserID,
		ClinicID: req.ClinicID,
		Date:     req.Date,
		Price:    req.Price,
		Status:   req.Status,
	}
}
