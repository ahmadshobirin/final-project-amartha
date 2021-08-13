package request

import "main-backend/bussiness/user"

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleID   int    `json:"role_id"`
}

func (req *User) ToDomain() *user.Domain {
	return &user.Domain{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		RoleID:   req.RoleID,
	}
}
