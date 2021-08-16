package auth

import (
	"context"
	"main-backend/bussiness/user"
	"main-backend/controller/auth/response"
)

type Usecase interface {
	Register(ctx context.Context, data *user.Domain) (res response.AuthResponse, err error)
	Login(ctx context.Context, data *user.Domain) (res response.AuthResponse, err error)
}
