package role

import (
	"main-backend/bussiness/role"
	"time"
)

type Role struct {
	ID        int
	Code      string
	Name      string
	Status    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (rec *Role) ToDomain() role.Domain {
	return role.Domain{
		ID:        rec.ID,
		Code:      rec.Code,
		Name:      rec.Name,
		Status:    rec.Status,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}
