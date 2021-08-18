package role

import (
	"main-backend/bussiness/role"
	"main-backend/controller"
	"main-backend/controller/role/response"
	"main-backend/helper/str"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RoleController struct {
	roleUsecase role.Usecase
}

func NewRoleController(e *echo.Echo, cu role.Usecase) *RoleController {
	return &RoleController{
		roleUsecase: cu,
	}
}

func (ctrl *RoleController) Find(c echo.Context) error {
	ctx := c.Request().Context()
	resp, err := ctrl.roleUsecase.Find(ctx)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseController := []response.Role{}
	for _, value := range resp {
		responseController = append(responseController, *response.FromDomain(&value))
	}

	return controller.NewSuccessResponse(c, responseController)
}

func (ctrl *RoleController) FindByID(c echo.Context) error {
	ctx := c.Request().Context()
	ID := str.StringToInt(c.Param("id"))

	resp, err := ctrl.roleUsecase.FindByID(ctx, ID)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(&resp))
}
