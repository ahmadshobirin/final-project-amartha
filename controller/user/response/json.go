package response

import (
	"main-backend/bussiness/user"
	roleResp "main-backend/controller/role/response"
	"time"
)

type User struct {
	ID        int            `json:"id"`
	RoleID    int            `json:"role_id"`
	Role      *roleResp.Role `json:"role"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

func FromDomain(domain *user.Domain) (res *User) {
	if domain != nil {
		res = &User{
			ID:        domain.ID,
			RoleID:    domain.RoleID,
			Role:      roleResp.FromDomain(domain.Role),
			Name:      domain.Name,
			Email:     domain.Email,
			CreatedAt: domain.CreatedAt,
			UpdatedAt: domain.UpdatedAt,
		}
	}
	return res
}
