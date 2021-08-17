package response

import (
	"main-backend/bussiness/user"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	RoleID    int       `json:"role_id"`
	Role      string    `json:"role"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain user.Domain) User {
	return User{
		ID:        domain.ID,
		RoleID:    domain.RoleID,
		Role:      domain.Role,
		Name:      domain.Name,
		Email:     domain.Email,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
