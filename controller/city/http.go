package city

import (
	"main-backend/bussiness/city"
	"main-backend/controller"
	"main-backend/controller/city/response"
	"main-backend/helper/str"

	"net/http"

	"github.com/labstack/echo/v4"
)

type CityController struct {
	cityUsecase city.Usecase
}

func NewCityController(e *echo.Echo, cu city.Usecase) *CityController {
	return &CityController{
		cityUsecase: cu,
	}
}

func (ctrl *CityController) Find(c echo.Context) error {
	ctx := c.Request().Context()

	resp, err := ctrl.cityUsecase.Find(ctx)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseController := []response.City{}
	for _, value := range resp {
		responseController = append(responseController, response.FromDomain(value))
	}

	return controller.NewSuccessResponse(c, responseController)
}

func (ctrl *CityController) FindByID(c echo.Context) error {
	ctx := c.Request().Context()
	ID := str.StringToInt(c.Param("id"))

	resp, err := ctrl.cityUsecase.FindByID(ctx, ID)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp))
}
