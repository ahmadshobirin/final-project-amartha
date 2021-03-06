package auth

import (
	"main-backend/bussiness/auth"
	"main-backend/controller"
	"main-backend/controller/auth/response"
	"main-backend/controller/user/request"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	authUsecase auth.Usecase
}

func NewAuthController(e *echo.Echo, cu auth.Usecase) *AuthController {
	return &AuthController{
		authUsecase: cu,
	}
}

func (ctrl *AuthController) Register(c echo.Context) error {
	ctx := c.Request().Context()
	req := request.User{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	data, err := ctrl.authUsecase.Register(ctx, req.ToDomain())
	if err != nil {
		return err
	}

	resp := response.AuthResponse{
		Token: data,
	}

	return controller.NewSuccessResponse(c, resp)

}

func (ctrl *AuthController) Login(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.User{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	data, err := ctrl.authUsecase.Login(ctx, req.ToDomain())
	if err != nil {
		return err
	}

	resp := response.AuthResponse{
		Token: data,
	}

	return controller.NewSuccessResponse(c, resp)

}
