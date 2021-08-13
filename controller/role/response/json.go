package response

import (
	"main-backend/bussiness/role"
)

type Role struct {
	ID        int    `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	Status    bool   `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func FromDomain(domain role.Domain) Role {
	return Role{
		ID:        domain.ID,
		Code:      domain.Code,
		Name:      domain.Name,
		Status:    domain.Status,
		CreatedAt: domain.CreatedAt.Format("2006-01-01 15:04:05"),
		UpdatedAt: domain.UpdatedAt.Format("2006-01-01 15:04:05"),
	}
}
