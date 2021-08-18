package user

import (
	user "main-backend/bussiness/user"
	"main-backend/driver/database/role"

	"time"
)

type User struct {
	ID        int `gorm:"primaryKey"`
	RoleID    int
	Role      *role.Role
	Name      string
	Email     string
	Password  string
	Bio       string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func fromDomain(domain *user.Domain) *User {
	return &User{
		ID:       domain.ID,
		RoleID:   domain.RoleID,
		Name:     domain.Name,
		Email:    domain.Email,
		Password: domain.Password,
	}
}

func (rec *User) ToDomain() (res *user.Domain) {
	if rec != nil {
		res = &user.Domain{
			ID:        rec.ID,
			RoleID:    rec.RoleID,
			Role:      rec.Role.ToDomain(),
			Name:      rec.Name,
			Email:     rec.Email,
			Password:  rec.Password,
			CreatedAt: rec.CreatedAt,
			UpdatedAt: rec.UpdatedAt,
		}
	}
	return res
}
