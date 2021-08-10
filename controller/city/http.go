package city

import (
	"main-backend/bussiness/city"
	"main-backend/controller"
	"main-backend/controller/city/response"

	"net/http"

	"github.com/labstack/echo/v4"
)

type CityController struct {
	cityUsecase city.Usecase
}

func NewCityController(e *echo.Echo, cu city.Usecase) {
	controller := &CityController{
		cityUsecase: cu,
	}

	category := e.Group("/api/v1/city")
	category.GET("", controller.Find)
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
