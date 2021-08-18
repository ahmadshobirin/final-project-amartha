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

func (rec *Role) ToDomain() (res *role.Domain) {
	if rec != nil {
		res = &role.Domain{
			ID:        rec.ID,
			Code:      rec.Code,
			Name:      rec.Name,
			Status:    rec.Status,
			CreatedAt: rec.CreatedAt,
			UpdatedAt: rec.UpdatedAt,
		}
	}
	return res
}
