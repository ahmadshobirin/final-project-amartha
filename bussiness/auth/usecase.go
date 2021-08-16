package auth

import (
	"context"
	"main-backend/app/middleware"
	"main-backend/bussiness/role"
	"main-backend/bussiness/user"
	"main-backend/controller/auth/response"
	"main-backend/helper/encrypt"
	"main-backend/helper/messages"
	"strings"
	"time"
)

type authUsecase struct {
	roleUsecase    role.Usecase
	userUsecase    user.Usecase
	contextTimeout time.Duration
	jwtAuth        *middleware.ConfigJWT
}

func NewAuthUsecase(timeout time.Duration, userUC user.Usecase, roleUC role.Usecase, jwt *middleware.ConfigJWT) Usecase {
	return &authUsecase{
		roleUsecase:    roleUC,
		userUsecase:    userUC,
		jwtAuth:        jwt,
		contextTimeout: timeout,
	}
}

func (uc authUsecase) Register(ctx context.Context, data *user.Domain) (res response.AuthResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	role, err := uc.roleUsecase.FindByCode(ctx, "US")
	if err != nil {
		return res, err
	}

	newUser, err := uc.userUsecase.Store(ctx, data, role.ID)
	if err != nil {
		return res, err
	}

	res = response.AuthResponse{
		Token: uc.jwtAuth.GenerateToken(newUser.ID),
	}

	return res, err
}

func (uc authUsecase) Login(ctx context.Context, data *user.Domain) (res response.AuthResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if strings.TrimSpace(data.Email) == "" && strings.TrimSpace(data.Password) == "" {
		return res, messages.ErrUsernamePasswordNotFound
	}

	user, err := uc.userUsecase.FindByEmail(ctx, data.Email)
	if err != nil {
		return res, err
	}

	if !encrypt.ValidateHash(data.Password, user.Password) {
		return res, messages.ErrInternalServer
	}

	res = response.AuthResponse{
		Token: uc.jwtAuth.GenerateToken(user.ID),
	}

	return res, err
}
