package user

import (
	"main-backend/app/middleware"
	"main-backend/bussiness/user"
	"main-backend/controller"
	"main-backend/controller/user/request"
	"main-backend/controller/user/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUseCase user.Usecase
	jwtAuth     *middleware.ConfigJWT
}

func NewUserController(e *echo.Echo, uc user.Usecase, jwt *middleware.ConfigJWT) *UserController {
	return &UserController{
		userUseCase: uc,
		jwtAuth:     jwt,
	}
}

func (ctrl *UserController) Fetch(c echo.Context) error {
	ctx := c.Request().Context()

	page, _ := strconv.Atoi(c.QueryParam("page"))
	perpage, _ := strconv.Atoi(c.QueryParam("per_page"))

	resp, count, err := ctrl.userUseCase.Fetch(ctx, page, perpage)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	res := []response.User{}
	for _, value := range resp {
		res = append(res, response.FromDomain(value))
	}

	return controller.NewSuccessResponseWithTotal(c, res, count)
}

func (ctrl *UserController) Profile(c echo.Context) error {
	ctx := c.Request().Context()

	user := ctrl.jwtAuth.GetUser(c)

	result, err := ctrl.userUseCase.FindByID(ctx, user.ID)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(result))
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
