package auth

import (
	"context"
	"main-backend/bussiness/user"
)

type Usecase interface {
	Register(ctx context.Context, data *user.Domain) (res string, err error)
	Login(ctx context.Context, data *user.Domain) (res string, err error)
}
