package user

import (
	"main-backend/bussiness/user"
	"main-backend/controller"
	"main-backend/controller/user/request"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUseCase user.Usecase
}

func NewUserController(e *echo.Echo, uc user.Usecase) *UserController {
	return &UserController{
		userUseCase: uc,
	}
}

func (ctrl *UserController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.User{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	_, err := ctrl.userUseCase.Store(ctx, req.ToDomain(), 0)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, "Successfully inserted")
}
